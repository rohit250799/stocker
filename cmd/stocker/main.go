package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
	"go-stocker/pkg/db"
	_ "github.com/lib/pq"
)

// for PG-admin
type User struct {
	Name  string `db:"username"`
	Email string `db:"email"`
}


func main() {

	//to find the .env file
	cwd, _ := os.Getwd()
    log.Println("Current working directory:", cwd)
    err := godotenv.Load(filepath.Join(cwd, ".env"))
    if err != nil {
        log.Printf("Error loading .env file: %v", err)
    }

	//connect to db
	db_connect, err := db.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	//get_market_data
	var company_symbol string
	fmt.Println("Enter the company symbol to use: ")
	fmt.Scan(&company_symbol)

	//Inserting company overview to db
	company, err := db.Insert_Company_Overview_data(company_symbol)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The entered symbol is: %v", company)

	// //Insert weekly series data to db
	company_weekly_data, err := db.Insert_Time_Series_Weekly_data(company_symbol)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The weekly timeseries is entered on: %v", company_weekly_data)

	http.HandleFunc("/view/", makeHandler(ViewHandler))
	http.HandleFunc("/edit/", makeHandler(EditHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	defer db_connect.Close()
}
