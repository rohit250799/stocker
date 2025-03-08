package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)


func check_all_market_open_status() bool {
	api_key := os.Getenv("apikey")
	if api_key == ""{
		log.Fatal("Api key not found in .env file")
	}
	
	baseUrl := "https://www.alphavantage.co/query"
	queryParams := "?function=MARKET_STATUS&apikey=" + api_key
	fullUrl := baseUrl + queryParams

	response, err := http.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
	return true
}

func get_market_data() {
	if check_all_market_open_status() {
		fmt.Println("The market is open right now.")
		os.Exit(0)

	} else {
		fmt.Println("The market is not open right now.")
		os.Exit(1)
	}
}
