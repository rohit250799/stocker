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

    //connectDB()

    db, err := db.ConnectDB();
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

    //get_market_data

     if pkg.Check_company_existence()  {
        pkg.Get_company_overview("AMD")
     }  
}

