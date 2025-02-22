package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)



func main() {

    //to find the .env file
    err:= godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading to .env file: %s", err)
    }

    my_db_port := os.Getenv("DB_PORT") 

    fmt.Println(my_db_port)
}

