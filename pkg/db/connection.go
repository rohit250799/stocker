package db

import (
	"os"
	"fmt"
	"database/sql"
	"log"
)

func ConnectDB() (*sql.DB, error) {
    db_host := os.Getenv("DB_HOST");
    db_port := os.Getenv("DB_PORT");
    db_user := os.Getenv("POSTGRES_USER")
    db_password := os.Getenv("POSTGRES_PASSWORD")
    db_name := os.Getenv("POSTGRES_DB")

    

    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        db_host, db_port, db_user, db_password, db_name);
    db, err := sql.Open("postgres", connStr);
    if err != nil {
        log.Fatalln(err);
    }

    //defer db.Close();

    //Testing connection to db
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    } else {
        log.Println("Successfully connected to db");
    }

    return db, nil;
}