package main

import (
	"fmt"
	"log"
	//"os"

	"github.com/joho/godotenv"

	//"database/sql"

    "go-stocker/pkg"

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

    //connect to db

    db, err := db.ConnectDB();
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()

    //get_market_data

    var company_symbol string
    fmt.Println("Enter the company symbol to use: ") 
    fmt.Scan(&company_symbol)

    if pkg.Check_company_existence(company_symbol)  {
        pkg.Get_company_overview(company_symbol)
    }  
}

