package main

import (
	//"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"


	"go-stocker/pkg/db"

	_ "github.com/lib/pq"
)

//for PG-admin
type User struct {
    Name string `db:"username"`
    Email string `db:"email"`
}

func main() {

    //to find the .env file
    err := godotenv.Load("../../.env")
    if err != nil {
        log.Fatalf("Error loading to .env file: %s", err)
    }

    //connect to db and create a new table

    db_connect, err := db.ConnectDB();
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db_connect.Close()

    //get_market_data

    var company_symbol string
    fmt.Println("Enter the company symbol to use: ") 
    fmt.Scan(&company_symbol)

    //Inserting data from api response to db

    company, err := db.Insert_Company_Overview_data(company_symbol)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("The entered symbol is: %v", company)
      
}

