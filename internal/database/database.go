package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/KznRkjp/go-keeper.git/internal/flags"
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

func createInitialDB(db *sql.DB) error {
	fmt.Println("DB String", flags.FlagDBString)
	ctx := context.Background()
	insertDynStmtUser := `CREATE TABLE url_users (id SERIAL PRIMARY KEY, uuid TEXT UNIQUE, token TEXT);`
	var err error
	_, err = db.ExecContext(ctx, insertDynStmtUser)
	if err != nil {
		log.Println("Database 'user' exists", err)
	}

	insertDynStmtURL := `CREATE TABLE url (id SERIAL PRIMARY KEY,
		 									correlationid TEXT,
											url_user_uuid TEXT,
											shorturl TEXT, 
											originalurl TEXT,
											deleted_flag BOOLEAN DEFAULT FALSE,
											CONSTRAINT fk_url_user_uuid FOREIGN KEY (url_user_uuid) REFERENCES url_users (uuid));`
	_, err = db.ExecContext(ctx, insertDynStmtURL)
	if err != nil {
		log.Println("Database 'url' exists", err)
	}
	return err
}
