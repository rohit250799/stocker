package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"strconv"

	//"strings"
	"encoding/json"
)

type Company_overview struct {
	Symbol                     string  `json:"Symbol"`
	AssetType                  string  `json:"AssetType"`
	Name                       string  `json:"Name"`
	Description                string  `json:"Description"`
	CIK                        string  `json:"CIK"`
	Exchange                   string  `json:"Exchange"`
	Currency                   string  `json:"Currency"`
	Country                    string  `json:"Country"`
	Sector                     string  `json:"Sector"`
	Industry                   string  `json:"Industry"`
	Address                    string  `json:"Address"`
	OfficialSite               string  `json:"OfficialSite"`
	FiscalYearEnd              string  `json:"FiscalYearEnd"`
	LatestQuarter              string  `json:"LatestQuarter"`
	MarketCapitalization       int64   `json:"MarketCapitalization,string"`
	EBITDA                     int64   `json:"EBITDA,string"`
	PERatio                    float64 `json:"PERatio,string"`
	PEGRatio                   float64 `json:"PEGRatio,string"`
	BookValue                  float64 `json:"BookValue,string"`
	DividendPerShare           string  `json:"DividendPerShare"`
	DividendYield              string  `json:"DividendYield"`
	EPS                        float64 `json:"EPS,string"`
	RevenuePerShareTTM         float64 `json:"RevenuePerShareTTM,string"`
	ProfitMargin               float64 `json:"ProfitMargin,string"`
	OperatingMarginTTM         float64 `json:"OperatingMarginTTM,string"`
	ReturnOnAssetsTTM          float64 `json:"ReturnOnAssetsTTM,string"`
	ReturnOnEquityTTM          float64 `json:"ReturnOnEquityTTM,string"`
	RevenueTTM                 int64   `json:"RevenueTTM,string"`
	GrossProfitTTM             int64   `json:"GrossProfitTTM,string"`
	DilutedEPSTTM              float64 `json:"DilutedEPSTTM,string"`
	QuarterlyEarningsGrowthYOY float64 `json:"QuarterlyEarningsGrowthYOY,string"`
	QuarterlyRevenueGrowthYOY  float64 `json:"QuarterlyRevenueGrowthYOY,string"`
	AnalystTargetPrice         float64 `json:"AnalystTargetPrice,string"`
	AnalystRatingStrongBuy     int     `json:"AnalystRatingStrongBuy,string"`
	AnalystRatingBuy           int     `json:"AnalystRatingBuy,string"`
	AnalystRatingHold          int     `json:"AnalystRatingHold,string"`
	AnalystRatingSell          int     `json:"AnalystRatingSell,string"`
	AnalystRatingStrongSell    int     `json:"AnalystRatingStrongSell,string"`
	TrailingPE                 float64 `json:"TrailingPE,string"`
	ForwardPE                  float64 `json:"ForwardPE,string"`
	PriceToSalesRatioTTM       float64 `json:"PriceToSalesRatioTTM,string"`
	PriceToBookRatio           float64 `json:"PriceToBookRatio,string"`
	EVToRevenue                float64 `json:"EVToRevenue,string"`
	EVToEBITDA                 float64 `json:"EVToEBITDA,string"`
	Beta                       float64 `json:"Beta,string"`
	FiftyTwoWeekHigh           float64 `json:"52WeekHigh,string"`
	FiftyTwoWeekLow            float64 `json:"52WeekLow,string"`
	FiftyDayMovingAverage      float64 `json:"50DayMovingAverage,string"`
	TwoHundredDayMovingAverage float64 `json:"200DayMovingAverage,string"`
	SharesOutstanding          int64   `json:"SharesOutstanding,string"`
	DividendDate               string  `json:"DividendDate"`
	ExDividendDate             string  `json:"ExDividendDate"`
}

type TimeSeriesWeeklyData struct {
	Id          int       `json:"-"`
	Symbol      string    `json:"Symbol"`
	Date        time.Time `json:"Date"`
	Open_price  float64   `json:"Open_Price"`
	High_price  float64   `json:"High_Price"`
	Low_price   float64   `json:"Low_Price"`
	Close_price float64   `json:"Close_Price"`
	Volume      int64     `json:"Volume"`
}

type Weekly_Time_Series_API_Response struct {
	MetaData           map[string]string               `json:"Meta Data"`
	Weekly_Time_Series map[string]map[string]string 			`json:"Weekly Time Series"`
}

type DemoResponse struct {
	Symbol    string `json:"symbol"`
	AssetType string `json:"assettype"`
	Name      string `json:"name"`
}

func check_all_market_open_status() bool {
	api_key := os.Getenv("apikey")
	if api_key == "" {
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

func Get_company_overview(company_symbol string) Company_overview {
	api_key := os.Getenv("apikey")
	if api_key == "" {
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

	var company_overview_api_response Company_overview
	err = json.Unmarshal([]byte(responseData), &company_overview_api_response)
	if err != nil {
		fmt.Println("Error parsing JSON", err)
	}

	symbol := company_overview_api_response.Symbol
	assetType := company_overview_api_response.AssetType
	name := company_overview_api_response.Name
	description := company_overview_api_response.Description
	cik := company_overview_api_response.CIK
	exchange := company_overview_api_response.Exchange
	currency := company_overview_api_response.Currency
	country := company_overview_api_response.Country
	sector := company_overview_api_response.Sector
	industry := company_overview_api_response.Industry
	address := company_overview_api_response.Address
	officialSite := company_overview_api_response.OfficialSite
	fiscalYearEnd := company_overview_api_response.FiscalYearEnd
	latestQuarter := company_overview_api_response.LatestQuarter
	marketCapitalization := company_overview_api_response.MarketCapitalization
	ebitda := company_overview_api_response.EBITDA
	peRatio := company_overview_api_response.PERatio
	pegRatio := company_overview_api_response.PEGRatio
	bookValue := company_overview_api_response.BookValue
	dividendPerShare := company_overview_api_response.DividendPerShare
	dividendYield := company_overview_api_response.DividendYield
	eps := company_overview_api_response.EPS
	revenuePerShareTTM := company_overview_api_response.RevenuePerShareTTM
	profitMargin := company_overview_api_response.ProfitMargin
	operatingMarginTTM := company_overview_api_response.OperatingMarginTTM
	returnOnAssetsTTM := company_overview_api_response.ReturnOnAssetsTTM
	returnOnEquityTTM := company_overview_api_response.ReturnOnEquityTTM
	revenueTTM := company_overview_api_response.RevenueTTM
	grossProfitTTM := company_overview_api_response.GrossProfitTTM
	dilutedEPSTTM := company_overview_api_response.DilutedEPSTTM
	quarterlyEarningsGrowthYOY := company_overview_api_response.QuarterlyEarningsGrowthYOY
	quarterlyRevenueGrowthYOY := company_overview_api_response.QuarterlyRevenueGrowthYOY
	analystTargetPrice := company_overview_api_response.AnalystTargetPrice
	analystRatingStrongBuy := company_overview_api_response.AnalystRatingStrongBuy
	analystRatingBuy := company_overview_api_response.AnalystRatingBuy
	analystRatingHold := company_overview_api_response.AnalystRatingHold
	analystRatingSell := company_overview_api_response.AnalystRatingSell
	analystRatingStrongSell := company_overview_api_response.AnalystRatingStrongSell
	trailingPE := company_overview_api_response.TrailingPE
	forwardPE := company_overview_api_response.ForwardPE
	priceToSalesRatioTTM := company_overview_api_response.PriceToSalesRatioTTM
	priceToBookRatio := company_overview_api_response.PriceToBookRatio
	evToRevenue := company_overview_api_response.EVToRevenue
	evToEBITDA := company_overview_api_response.EVToEBITDA
	beta := company_overview_api_response.Beta
	fiftyTwoWeekHigh := company_overview_api_response.FiftyTwoWeekHigh
	fiftyTwoWeekLow := company_overview_api_response.FiftyTwoWeekLow
	fiftyDayMovingAverage := company_overview_api_response.FiftyDayMovingAverage
	twoHundredDayMovingAverage := company_overview_api_response.TwoHundredDayMovingAverage
	sharesOutstanding := company_overview_api_response.SharesOutstanding
	dividendDate := company_overview_api_response.DividendDate
	exDividendDate := company_overview_api_response.ExDividendDate

	return Company_overview{symbol, assetType, name, description, cik, exchange, currency, country, sector, industry, address, officialSite, fiscalYearEnd, latestQuarter, marketCapitalization, ebitda, peRatio, pegRatio, bookValue, dividendPerShare, dividendYield, eps, revenuePerShareTTM, profitMargin, operatingMarginTTM,
		returnOnAssetsTTM, returnOnEquityTTM, revenueTTM, grossProfitTTM, dilutedEPSTTM, quarterlyEarningsGrowthYOY, quarterlyRevenueGrowthYOY, analystTargetPrice, analystRatingStrongBuy, analystRatingBuy, analystRatingHold, analystRatingSell, analystRatingStrongSell, trailingPE, forwardPE, priceToSalesRatioTTM, priceToBookRatio, evToRevenue, evToEBITDA, beta, fiftyTwoWeekHigh, fiftyTwoWeekLow, fiftyDayMovingAverage, twoHundredDayMovingAverage, sharesOutstanding, dividendDate, exDividendDate}

}

func Get_Time_Series_Weekly_data(company_symbol string) TimeSeriesWeeklyData {
	api_key := os.Getenv("apikey")
	if api_key == "" {
		log.Fatal("Api key not found in .env file")
	}

	if company_symbol == "" {
		fmt.Println("Company symbol cannot be empty")
		os.Exit(1)
	}

	baseUrl := "https://www.alphavantage.co/query"
	queryParams_1 := "?function=TIME_SERIES_WEEKLY&symbol=" + company_symbol
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

	var timeSeriesWeeklyApiResponse Weekly_Time_Series_API_Response
	err = json.Unmarshal([]byte(responseData), &timeSeriesWeeklyApiResponse)
	if err != nil {
		fmt.Println("Error parsing JSON", err)
	}

	id := 0
	symbol := timeSeriesWeeklyApiResponse.MetaData["2. Symbol"]
	date := timeSeriesWeeklyApiResponse.MetaData["3. Date"]
	open_price := timeSeriesWeeklyApiResponse.MetaData["4. Open Price"]
	high_price := timeSeriesWeeklyApiResponse.MetaData["5. High Price"]
	low_price := timeSeriesWeeklyApiResponse.MetaData["6. Low Price"]
	close_price := timeSeriesWeeklyApiResponse.MetaData["7. Close Price"]
	volume := timeSeriesWeeklyApiResponse.MetaData["8. Volume"]

	//layout := "2006-03-27"
	parseDate := time.Now()
	if date != "" {
		var err error
		parseDate, err = time.Parse("2006-01-02", date)
		if err != nil {
			fmt.Println("Error parsing date: ", err)
			return TimeSeriesWeeklyData{}
		}
	}
	// parseDate, err := time.Parse(layout, date)
	// if err != nil {
	// 	fmt.Println("Error parsing date: ", err)
	// 	return TimeSeriesWeeklyData{}
	// }

	openPrice, err := strconv.ParseFloat(open_price, 64)
	if err != nil {
		fmt.Println("Error parsing open price: ", err)
		return TimeSeriesWeeklyData{}
	}

	highPrice, err := strconv.ParseFloat(high_price, 64)
	if err != nil {
		fmt.Println("Error parsing high price: ", err)
		return TimeSeriesWeeklyData{}
	}

	lowPrice, err := strconv.ParseFloat(low_price, 64)
	if err != nil {
		fmt.Println("Error parsing low price: ", err)
		return TimeSeriesWeeklyData{}
	}

	closePrice, err := strconv.ParseFloat(close_price, 64)
	if err != nil {
		fmt.Println("Error parsing close price: ", err)
		return TimeSeriesWeeklyData{}
	}

	volumeInt, err := strconv.ParseInt(volume, 10, 64)
	if err != nil {
		fmt.Println("Error parsing volume: ", volumeInt)
		return TimeSeriesWeeklyData{}
	}


	return TimeSeriesWeeklyData{id, symbol, parseDate, openPrice, highPrice, lowPrice, closePrice, volumeInt}
}

func Demo_company_overview(company_symbol string) []string {
	api_key := os.Getenv("apikey")
	if api_key == "" {
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

	var apiResponse DemoResponse
	err = json.Unmarshal([]byte(responseData), &apiResponse)
	if err != nil {
		fmt.Println("Error parsing JSON", err)
	}

	name := apiResponse.Name
	symbol := apiResponse.AssetType
	assettype := apiResponse.Symbol

	return []string{name, assettype, symbol}

}
