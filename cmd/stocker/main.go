package main

import (
	"database/sql"
	"fmt"
	"log"

	//"os"

	"github.com/joho/godotenv"

	//"database/sql"

	"go-stocker/pkg"

	"go-stocker/pkg/db"

	_ "github.com/lib/pq"
)

var db_connect *sql.DB

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

    if pkg.Check_company_existence(company_symbol)  {
        pkg.Get_company_overview(company_symbol)
        //db.Insert_Company_Overview_data()

        compId, err := db.Insert_Company_Overview_data(db.Company_overview {
            Symbol: "ABC",
            AssetType: "Test_One",
            Name: "Test_Name",
        })
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID of added company: %v\n", compId)
    }  
}

