package database

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/KznRkjp/go-keeper.git/internal/encrypt"
	"github.com/KznRkjp/go-keeper.git/internal/flags"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/KznRkjp/go-keeper.git/internal/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

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
// Все так же падает при старте, надо переделать
func createInitialDB(db *sql.DB) error {

	//что бы не забыть как запускать
	mlogger.Info("DB String" + flags.FlagDBString)
	var err error

	//START ########## Таблица пользователей
	insertDynStmtUsers := `CREATE TABLE go_k_users (id SERIAL PRIMARY KEY, 
											email text not null unique,
											password TEXT,
											created_at timestamp default current_timestamp);`
	// var errUsers error
	_, errUsers := db.Exec(insertDynStmtUsers)
	time.Sleep(time.Second * 1)
	if errUsers.Error() == "ERROR: relation \"go_k_users\" already exists (SQLSTATE 42P07)" {
		errUsers = nil
		mlogger.Info("Table go_k_users already exists")
	}
	if errUsers != nil {
		log.Fatal(errUsers)
		err = errUsers
	}

	//STOP ########## Таблица пользователей

	//START ########## Таблица учетных данных - логопас
	insertDynStmtLogopass := `CREATE TABLE logopass (id SERIAL PRIMARY KEY,
		 									name bytea,
											login bytea,
											password bytea,
											go_k_user_id INTEGER,
											created_at timestamp default current_timestamp,
											CONSTRAINT fk_go_k_user_id FOREIGN KEY (go_k_user_id) REFERENCES go_k_users (id));`
	// var errLP error
	_, errLP := db.Exec(insertDynStmtLogopass)
	if errLP.Error() == "ERROR: relation \"logopass\" already exists (SQLSTATE 42P07)" {
		errLP = nil
		mlogger.Info("Table logopass already exists")
	}
	if errLP != nil {
		log.Fatal(errLP)
		err = errLP
	}
	time.Sleep(time.Second * 1)
	//STOP ######### Таблица учетных данных - логопас

	//START ######### Таблица учетных данных - банковские карты
	insertDynStmtBankCard := `CREATE TABLE bank_card (id SERIAL PRIMARY KEY,
		 									card_holder_name bytea,
											card_number bytea,
											expiration_date bytea,
											go_k_user_id INTEGER,
											created_at timestamp default current_timestamp,
											CONSTRAINT fk_go_k_user_id FOREIGN KEY (go_k_user_id) REFERENCES go_k_users (id));`
	_, errBC := db.Exec(insertDynStmtBankCard)
	time.Sleep(time.Second * 1)

	if errBC.Error() == "ERROR: relation \"bank_card\" already exists (SQLSTATE 42P07)" {
		errBC = nil
		mlogger.Info("Table bank_card already exists")
	}
	if errBC != nil {
		mlogger.Logger.Fatal(errBC.Error())
		err = errBC
	}
	//STOP ######### Таблица учетных данных - банковские карты

	// START ######### Таблица текстовых данных

	insertDynStmtTextData := `CREATE TABLE text_data (id SERIAL PRIMARY KEY,
											name bytea,
		 									text bytea,
											go_k_user_id INTEGER,
											created_at timestamp default current_timestamp,
											CONSTRAINT fk_go_k_user_id FOREIGN KEY (go_k_user_id) REFERENCES go_k_users (id));`
	_, errTXT := db.Exec(insertDynStmtTextData)
	// time.Sleep(time.Second * 1)
	if errTXT.Error() == "ERROR: relation \"text_data\" already exists (SQLSTATE 42P07)" {
		errTXT = nil
		mlogger.Info("Table text_data already exists")
	}
	if errTXT != nil {
		log.Fatal(errTXT)
		err = errTXT
	}

	// STOP ######### Таблица текстовых данных

	// START ######### Таблица бинарных данных
	insertDynStmtBinaryData := `CREATE TABLE binary_data (id SERIAL PRIMARY KEY,
											name bytea,
		 									file_name bytea,
											location bytea,
											go_k_user_id INTEGER,
											created_at timestamp default current_timestamp,
											CONSTRAINT fk_go_k_user_id FOREIGN KEY (go_k_user_id) REFERENCES go_k_users (id));`
	_, errBin := db.Exec(insertDynStmtBinaryData)
	// time.Sleep(time.Second * 1)
	if errBin.Error() == "ERROR: relation \"binary_data\" already exists (SQLSTATE 42P07)" {
		errBin = nil
		mlogger.Info("Table binary_data already exists")
	}
	if errBin != nil {
		log.Fatal(errBin)
		err = errBin
	}
	// STOP ######### Таблица бинарных данных

	return err
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
