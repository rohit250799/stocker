package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"database/sql"

	_ "github.com/lib/pq"
)

//for PG-admin
type User struct {
    Name string `db:"username"`
    Email string `db:"email"`
}

func connectDB() (*sql.DB, error) {
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

    //Testing connection to db
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    } else {
        log.Println("Successfully connected to db");
    }

    return db, nil;
}

func main() {

    err := godotenv.Load("../../.env")
    if err != nil {
        log.Fatalf("Error loading to .env file: %s", err)
    }

    //connectDB()

    db, err := connectDB();
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()

    //create table 
    // err = CreateTable(db);
    // if err != nil {
    //     fmt.Println("Error while creating table", err);
    //     return;
    // }

    get_market_data()

    testing_one()
}

