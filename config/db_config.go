package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eulbyvan/idk/app/go-user-management/pkg"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	host := pkg.GetEnv("DB_HOST")
	port := pkg.GetEnv("DB_PORT")
	user := pkg.GetEnv("DB_USER")
	password := pkg.GetEnv("DB_PASSWORD")
	dbName := pkg.GetEnv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database")
	return db, nil
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Printf("Error closing the database connection: %s", err)
	} else {
		log.Println("Database connection closed")
	}
}

