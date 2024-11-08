package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/KznRkjp/go-keeper.git/internal/encrypt"
	"github.com/KznRkjp/go-keeper.git/internal/flags"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

var goKUsersTable = models.Table{
	Name: "go_k_users",
	InitString: `CREATE TABLE go_k_users (id SERIAL PRIMARY KEY, 
											email text not null unique,
											password TEXT,
											created_at timestamp default current_timestamp);`,
}

var logopassTable = models.Table{
	Name: "logopass",
	InitString: `CREATE TABLE logopass (id SERIAL PRIMARY KEY,
		name bytea,
	   login bytea,
	   password bytea,
	   go_k_user_id INTEGER,
	   created_at timestamp default current_timestamp,
	   CONSTRAINT fk_go_k_user_id FOREIGN KEY (go_k_user_id) REFERENCES go_k_users (id));`,
}

var bCTable = models.Table{
	Name: "bank_card",
	InitString: `CREATE TABLE bank_card (id SERIAL PRIMARY KEY,
											card_name bytea,
		 									card_holder_name bytea,
											card_number bytea,
											expiration_date bytea,
											go_k_user_id INTEGER,
											created_at timestamp default current_timestamp,
											CONSTRAINT fk_go_k_user_id FOREIGN KEY (go_k_user_id) REFERENCES go_k_users (id));`,
}

var txtMsgTable = models.Table{
	Name: "text_data",
	InitString: `CREATE TABLE text_data (id SERIAL PRIMARY KEY,
											name bytea,
		 									text bytea,
											go_k_user_id INTEGER,
											created_at timestamp default current_timestamp,
											CONSTRAINT fk_go_k_user_id FOREIGN KEY (go_k_user_id) REFERENCES go_k_users (id));`,
}

var binTable = models.Table{
	Name: "binary_data",
	InitString: `CREATE TABLE binary_data (id SERIAL PRIMARY KEY,
											name bytea,
		 									file_name bytea,
											location bytea,
											go_k_user_id INTEGER,
											created_at timestamp default current_timestamp,
											CONSTRAINT fk_go_k_user_id FOREIGN KEY (go_k_user_id) REFERENCES go_k_users (id));`,
}

var tableList models.TableList

// InitDB запускает подключение к базе данных
func InitDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("pgx", dataSourceName)
	if err != nil {
		return err
	}
	// defer db.Close()
	createInitialDB(db)
	return db.Ping()
}

// Just in case we going to use db outside the module
func GetDB() *sql.DB {
	return db
}

// Создание необходимых таблиц - можно конечно универасальную функцию... но потом.
func collectTables(tablesList *models.TableList) {
	tablesList.Tables = append(tableList.Tables, goKUsersTable)
	tablesList.Tables = append(tableList.Tables, logopassTable)
	tablesList.Tables = append(tableList.Tables, bCTable)
	tablesList.Tables = append(tableList.Tables, txtMsgTable)
	tablesList.Tables = append(tableList.Tables, binTable)
}

func createTable(db *sql.DB, table models.Table) error {
	_, err := db.Exec(table.InitString)
	if err != nil {
		if err.Error() == "ERROR: relation \""+table.Name+"\" already exists (SQLSTATE 42P07)" {
			mlogger.Info("Table " + table.Name + " already exists")
			return nil
		}

		return err

	}
	return nil
}

// Создание базы данных (таблиц)
func createInitialDB(db *sql.DB) error {

	//что бы не забыть как запускать
	mlogger.Info("DB String" + flags.FlagDBString)
	var err error
	collectTables(&tableList)
	// fmt.Println(tableList)
	for _, table := range tableList.Tables {
		err = createTable(db, table)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil

}

// RegisterUser - регистрация нового пользователя
func RegisterUser(user *models.User, ctx context.Context) (*models.User, error) {
	insertDynStmt := `insert into "go_k_users"("email", "password") values($1, $2)`
	password, err := encrypt.HashPassword(user.Password)
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return nil, err
	}
	_, err = db.ExecContext(ctx, insertDynStmt, user.Email, password)
	if err != nil {
		mlogger.Logger.Error(err.Error())
		return nil, err
	}
	// fmt.Println(val.RowsAffected())
	return returnUser(user.Email)
}

// LoginUser - проверка пароля пользователя, если ок - выдача его ID
func LoginUser(user *models.User, ctx context.Context) (*models.User, error) {
	db_user, err := returnUser(user.Email)
	if err != nil {
		return nil, err
	}
	mlogger.Info("Found user with id " + strconv.FormatInt(db_user.ID, 10))
	if !encrypt.VerifyPassword(user.Password, db_user.Password) {
		mlogger.Info(db_user.Password)
		mlogger.Info(user.Password)
		return nil, errors.New("wrong password")
	}
	return db_user, nil
}

// returnUser - вспомогатеьная функция для получения ID пользователя по email
func returnUser(email string) (*models.User, error) {
	var user models.User
	err := db.QueryRow("select * from go_k_users where email = $1", email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetData -  возвращает ВСЕ данные.
func GetData(userId *int, ctx context.Context) (models.DBSearchAll, error) {
	var data models.DBSearchAll

	var errLP error
	data.LoginPass, errLP = dbSearchLoginPassword(userId, ctx)
	if errLP != nil {
		mlogger.Logger.Error(errLP.Error())
		return data, errLP
	}

	var errBC error
	data.BankCards, errBC = dbSearchBankCard(userId, ctx)
	if errBC != nil {
		mlogger.Logger.Error(errBC.Error())
		return data, errBC
	}

	var errTM error
	data.TextMsgs, errTM = dbSearchTextMessage(userId, ctx)
	if errTM != nil {
		mlogger.Logger.Error(errTM.Error())
		return data, errTM
	}

	var errBM error
	data.BinaryMsgs, errBM = dbSearchBinaryMessages(userId, ctx)
	if errBM != nil {
		mlogger.Logger.Error(errBM.Error())
		return data, errBM
	}

	return data, nil
}
