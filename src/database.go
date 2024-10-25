package main

import (
	"database/sql"
	"fmt"
	"log"
	"mysadapi/logs"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	passwd := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, passwd, dbname)
	var err error
	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the PostgreSQL database")
}

func CreateToDoTable() {
	query := `
	SET TIME ZONE 'Europe/Madrid';

	CREATE TABLE IF NOT EXISTS toDo (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		completed BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	logs.PostLog("INFO", "Table 'toDo' created if not exists!")
}
