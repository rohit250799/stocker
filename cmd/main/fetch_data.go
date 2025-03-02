package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	//"net/url"
	"os"
	//"github.com/joho/godotenv"
)

func check_market_open_status() bool {
	//api_key := os.Getenv("apikey")
	// url := "https://www.alphavantage.co/query?function=MARKET_STATUS&apikey="
	response, err := http.Get("https://www.alphavantage.co/query?function=MARKET_STATUS&apikey=test_key")

	if err != nil {
		fmt.Print(err.Error())
		return false
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
	return true
}

func get_market_data() {
	if check_market_open_status() {
		//response, err := http.Get("")
		fmt.Println("The market is open right now.")
		os.Exit(0)

	} else {
		fmt.Println("The market is not open right now.")
		os.Exit(1)
	}
}
