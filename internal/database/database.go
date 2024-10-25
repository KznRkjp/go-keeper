package database

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"

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
func createInitialDB(db *sql.DB) error {

	//что бы не забыть как запускать
	mlogger.Info("DB String" + flags.FlagDBString)

	// ctx := context.Background()

	//START ########## Таблица пользователей
	insertDynStmt := `CREATE TABLE go_k_users (id SERIAL PRIMARY KEY, 
											email text not null unique,
											password TEXT,
											created_at timestamp default current_timestamp);`
	var err error
	_, err = db.Exec(insertDynStmt)
	if err.Error() == "ERROR: relation \"go_k_users\" already exists (SQLSTATE 42P07)" {
		err = nil
		mlogger.Info("Table go_k_users already exists")
	}
	if err != nil {
		log.Fatal(err)
	}

	//STOP ########## Таблица пользователей

	//START ########## Таблица учетных данных - логопас
	insertDynStmt = `CREATE TABLE logopass (id SERIAL PRIMARY KEY,
		 									login TEXT,
											password TEXT,
											go_k_user_id INTEGER,
											created_at timestamp default current_timestamp,
											CONSTRAINT fk_go_k_user_id FOREIGN KEY (go_k_user_id) REFERENCES go_k_users (id));`
	_, err = db.Exec(insertDynStmt)
	if err.Error() == "ERROR: relation \"logopass\" already exists (SQLSTATE 42P07)" {
		err = nil
		mlogger.Info("Table logopass already exists")
	}
	if err != nil {
		log.Fatal(err)
	}
	//STOP ######### Таблица учетных данных - логопас

	//START ######### Таблица учетных данных - логопас
	insertDynStmt = `CREATE TABLE bank_card (id SERIAL PRIMARY KEY,
		 									card_holder_name TEXT,
											card_number TEXT,
											expiration_date TEXT,
											go_k_user_id INTEGER,
											created_at timestamp default current_timestamp,
											CONSTRAINT fk_go_k_user_id FOREIGN KEY (go_k_user_id) REFERENCES go_k_users (id));`
	_, err = db.Exec(insertDynStmt)
	if err.Error() == "ERROR: relation \"bank_card\" already exists (SQLSTATE 42P07)" {
		err = nil
		mlogger.Info("Table bank_card already exists")
	}
	if err != nil {
		mlogger.Logger.Fatal(err.Error())

	}
	//STOP ######### Таблица учетных данных - логопас

	// START ######### Таблица текстовых данных
	insertDynStmt = `CREATE TABLE text_data (id SERIAL PRIMARY KEY,
		 									text TEXT,
											go_k_user_id INTEGER,
											created_at timestamp default current_timestamp,
											CONSTRAINT fk_go_k_user_id FOREIGN KEY (go_k_user_id) REFERENCES go_k_users (id));`
	_, err = db.Exec(insertDynStmt)
	if err.Error() == "ERROR: relation \"text_data\" already exists (SQLSTATE 42P07)" {
		err = nil
		mlogger.Info("Table text_data already exists")
	}
	if err != nil {
		log.Fatal(err)
	}
	// STOP ######### Таблица текстовых данных

	// START ######### Таблица бинарных данных
	insertDynStmt = `CREATE TABLE binary_data (id SERIAL PRIMARY KEY,
		 									file_name TEXT,
											location TEXT,
											go_k_user_id INTEGER,
											created_at timestamp default current_timestamp,
											CONSTRAINT fk_go_k_user_id FOREIGN KEY (go_k_user_id) REFERENCES go_k_users (id));`
	_, err = db.Exec(insertDynStmt)
	if err.Error() == "ERROR: relation \"binary_data\" already exists (SQLSTATE 42P07)" {
		err = nil
		mlogger.Info("Table binary_data already exists")
	}
	if err != nil {
		log.Fatal(err)
	}
	// STOP ######### Таблица бинарных данных

	return err
}

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
	return user, nil
}

func returnUser(email string) (*models.User, error) {
	var user models.User
	err := db.QueryRow("select * from go_k_users where email = $1", email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
