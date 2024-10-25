package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/KznRkjp/go-keeper.git/internal/flags"
	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
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
	defer db.Close()
	createInitialDB(db)
	return db.Ping()
}

// Just in case we going to use db outside the module
func GetDB() *sql.DB {
	return db
}

// Создание необходимых таблиц
func createInitialDB(db *sql.DB) error {
	fmt.Println("DB String", flags.FlagDBString)
	ctx := context.Background()

	//Таблица пользователей
	insertDynStmt := `CREATE TABLE go_k_users (id SERIAL PRIMARY KEY, 
											email text not null unique,
											password TEXT,
											created_at timestamp default current_timestamp);`
	var err error
	_, err = db.ExecContext(ctx, insertDynStmt)
	if err.Error() == "ERROR: relation \"go_k_users\" already exists (SQLSTATE 42P07)" {
		err = nil
		mlogger.Logger.Info("Table go_k_users already exists")
	}
	if err != nil {
		log.Fatal(err)
	}
	//Таблица учетных данных - логопас
	insertDynStmt = `CREATE TABLE logopass (id SERIAL PRIMARY KEY,
		 									login TEXT,
											password TEXT,
											go_k_user_id INTEGER,
											created_at timestamp default current_timestamp,
											CONSTRAINT fk_go_k_user_id FOREIGN KEY (go_k_user_id) REFERENCES go_k_users (id));`
	_, err = db.ExecContext(ctx, insertDynStmt)
	if err.Error() == "ERROR: relation \"logopass\" already exists (SQLSTATE 42P07)" {
		err = nil
		mlogger.Logger.Info("Table logopass already exists")
	}
	if err != nil {
		log.Fatal(err)
	}

	//Таблица учетных данных - логопас
	insertDynStmt = `CREATE TABLE bank_card (id SERIAL PRIMARY KEY,
		 									card_holder_name TEXT,
											card_number TEXT,
											expiration_date TEXT,
											go_k_user_id INTEGER,
											created_at timestamp default current_timestamp,
											CONSTRAINT fk_go_k_user_id FOREIGN KEY (go_k_user_id) REFERENCES go_k_users (id));`
	_, err = db.ExecContext(ctx, insertDynStmt)
	if err.Error() == "ERROR: relation \"bank_card\" already exists (SQLSTATE 42P07)" {
		err = nil
		mlogger.Logger.Info("Table bank_card already exists")
	}
	if err != nil {
		log.Fatal(err)
	}

	return err
}
