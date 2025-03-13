package pkg

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

func Get_company_overview(company_symbol string) string {
	api_key := os.Getenv("apikey")
	if api_key == ""{
		log.Fatal("Api key not found in .env file")
	}

	if company_symbol == "" {
		fmt.Println("Company symbol cannot be empty")
		os.Exit(1)
	}

	baseUrl := "https://www.alphavantage.co/query"
	queryParams_1 := "?function=OVERVIEW&symbol=" + company_symbol
	queryParams_2 := "&apikey=" + api_key

	fullUrl := baseUrl + queryParams_1 + queryParams_2
	
	response, err := http.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(responseData)
}


func Check_company_existence() bool {

	var company_symbol string 

	fmt.Println("Enter the symbol of the company to search: ")
	fmt.Scanf("%s", &company_symbol)

	if Get_company_overview(company_symbol) == "" || Get_company_overview(company_symbol) == "{}" {
		fmt.Println("No, the company does not exist")
		return false
	} else {
		fmt.Println("Yes, the company exists")
		return true
	}
}

// func display_company_info() string {
		
// }