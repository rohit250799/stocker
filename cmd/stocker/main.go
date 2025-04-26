package main

import (
	//"context"
	"fmt"
	"log"
	"net/http"
	//"os"

	//"html/template"
	//"templates"

	//"github.com/gorilla/mux"
	"github.com/a-h/templ"
	"github.com/joho/godotenv"

	"go-stocker/pkg/db"
	//"go-stocker/templates"

	//"go-stocker/templates"

	_ "github.com/lib/pq"
)

//for PG-admin
type User struct {
    Name string `db:"username"`
    Email string `db:"email"`
}

// templates map to store pre-parsed templates
//var templates map[string]*template.Template

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

    //creating time series weekly table
    //db.CreateTimeSeriesWeeklyTable(db_connect)

    defer db_connect.Close()

    // //get_market_data -> relevant statements starts, to be uncommented after this

    // var company_symbol string
    // fmt.Println("Enter the company symbol to use: ") 
    // fmt.Scan(&company_symbol)

    // // //Inserting company overview to db

    // company, err := db.Insert_Company_Overview_data(company_symbol)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // fmt.Printf("The entered symbol is: %v", company)

    


    // //Insert weekly series data to db
    // company_weekly_data, err := db.Insert_Time_Series_Weekly_data(company_symbol)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // fmt.Printf("The weekly timeseries is entered on: %v", company_weekly_data)  -> relevant statements, to be uncommented after this


    //for the frontend
    // LoadTemplates()

    // fs := http.FileServer(http.Dir("/.static"))
    // http.Handle("/static/", http.StripPrefix("/static/", fs))

    // http.HandleFunc("/", HomeHandler)
    // http.HandleFunc("/about", AboutHandler)

    // http.ListenAndServe(":8080", nil)

    //component := templates.Hello("Rohit")
	
	http.Handle("/", templ.Handler(LandingPage()))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)

    //templates.RegisterButton("Register").Render(context.Background(), os.Stdout)
    //templates.LoginButton("Login").Render(context.Background(), os.Stdout)

}

