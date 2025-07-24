package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := "user=postgres password=postgres dbname=todo_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Databasega ulanishda xatolik:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Databasega ulanish (ping) muvaffaqiyatsiz:", err)
	}

	return db
}
