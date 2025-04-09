package db

import (
	"database/sql"
	"fmt"
	"go-stocker/pkg"
	"log"
	"strings"
)

//for PG-admin
type User struct {
    Name string `db:"username"`
    Email string `db:"email"`
}

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
	FiftyTwoWeekHigh                float64 `json:"52WeekHigh,string"`
	FiftyTwoWeekLow                 float64 `json:"52WeekLow,string"`
	FiftyDayMovingAverage        float64 `json:"50DayMovingAverage,string"`
	TwoHundredDayMovingAverage       float64 `json:"200DayMovingAverage,string"`
	SharesOutstanding          int64   `json:"SharesOutstanding,string"`
	DividendDate               string  `json:"DividendDate"`
	ExDividendDate             string  `json:"ExDividendDate"`
}

type Demo_Company_overview struct {
	Symbol                     string  `json:"Symbol"`
	AssetType                  string  `json:"AssetType"`
	Name                       string  `json:"Name"`
}

func CreateCompanyOverviewTable(db *sql.DB) error {
	fmt.Println("Enter the name of the table: ")

	var get_table_name string
	fmt.Scanln(&get_table_name)

	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
    Symbol TEXT PRIMARY KEY,
    AssetType TEXT,
    Name TEXT,
    Description TEXT,
    CIK TEXT,
    Exchange TEXT,
    Currency TEXT,
    Country TEXT,
    Sector TEXT,
    Industry TEXT,
    Address TEXT,
    OfficialSite TEXT,
    FiscalYearEnd TEXT,
    LatestQuarter TEXT,
    MarketCapitalization BIGINT,
    EBITDA BIGINT,
    PERatio DOUBLE PRECISION,
    PEGRatio DOUBLE PRECISION,
    BookValue DOUBLE PRECISION,
    DividendPerShare TEXT,
    DividendYield TEXT,
    EPS DOUBLE PRECISION,
    RevenuePerShareTTM DOUBLE PRECISION,
    ProfitMargin DOUBLE PRECISION,
    OperatingMarginTTM DOUBLE PRECISION,
    ReturnOnAssetsTTM DOUBLE PRECISION,
    ReturnOnEquityTTM DOUBLE PRECISION,
    RevenueTTM BIGINT,
    GrossProfitTTM BIGINT,
    DilutedEPSTTM DOUBLE PRECISION,
    QuarterlyEarningsGrowthYOY DOUBLE PRECISION,
    QuarterlyRevenueGrowthYOY DOUBLE PRECISION,
    AnalystTargetPrice DOUBLE PRECISION,
    AnalystRatingStrongBuy INT,
    AnalystRatingBuy INT,
    AnalystRatingHold INT,
    AnalystRatingSell INT,
    AnalystRatingStrongSell INT,
    TrailingPE DOUBLE PRECISION,
    ForwardPE DOUBLE PRECISION,
    PriceToSalesRatioTTM DOUBLE PRECISION,
    PriceToBookRatio DOUBLE PRECISION,
    EVToRevenue DOUBLE PRECISION,
    EVToEBITDA DOUBLE PRECISION,
    Beta DOUBLE PRECISION,
    FiftyTwoWeekHigh DOUBLE PRECISION,
    FiftyTwoWeekLow DOUBLE PRECISION,
    FiftyDayMovingAverage DOUBLE PRECISION,
    TwoHundredDayMovingAverage DOUBLE PRECISION,
    SharesOutstanding BIGINT,
    DividendDate TEXT,
    ExDividendDate TEXT
	);`, get_table_name)
    
	_, err := db.Exec(query);
	if err != nil {
		return err;
	}
	fmt.Println("Demo table has been created successfully");
	return nil
}

func CreateTimeSeriesWeeklyTable(db *sql.DB) error {
	fmt.Println("Enter the name of the table: ")
	var getTableName string
	fmt.Scanln(&getTableName)

	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id SERIAL PRIMARY KEY,
		symbol TEXT NOT NULL,
		date DATE NOT NULL,
		open_price NUMERIC(10, 4),
		high_price NUMERIC(10, 4),
		low_price NUMERIC(10, 4),
		close_price NUMERIC(10, 4),
		volume BIGINT
	);`, getTableName)

	_, err := db.Exec(query);
	if err != nil {
		return err
	}
	fmt.Println("The table has been created")
	return nil
}

func Insert_Company_Overview_data(comp string) (string, error) {
	db_connection_pointer, error := ConnectDB()
	if error != nil {
		log.Fatal(error)
	}
	//var id int64

	company := pkg.Get_company_overview(comp)

	query := `INSERT INTO company_overview (Symbol, AssetType, Name, Description, CIK, Exchange, Currency, Country, Sector, Industry,
    Address, OfficialSite, FiscalYearEnd, LatestQuarter, MarketCapitalization, EBITDA, PERatio, PEGRatio, BookValue, DividendPerShare, DividendYield, EPS,RevenuePerShareTTM, ProfitMargin, OperatingMarginTTM, ReturnOnAssetsTTM, ReturnOnEquityTTM, RevenueTTM, GrossProfitTTM, DilutedEPSTTM, QuarterlyEarningsGrowthYOY, QuarterlyRevenueGrowthYOY, AnalystTargetPrice, AnalystRatingStrongBuy, AnalystRatingBuy, AnalystRatingHold, AnalystRatingSell, AnalystRatingStrongSell, TrailingPE, ForwardPE, PriceToSalesRatioTTM, PriceToBookRatio, EVToRevenue, EVToEBITDA, Beta, FiftyTwoWeekHigh, FiftyTwoWeekLow, FiftyDayMovingAverage, 
    TwoHundredDayMovingAverage, SharesOutstanding, DividendDate, ExDividendDate) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52) ON CONFLICT (Symbol) DO UPDATE SET MarketCapitalization = EXCLUDED.MarketCapitalization;`

	//executing the query with struct values
	_, err := db_connection_pointer.Exec(query,
		company.Symbol, company.AssetType, company.Name, company.Description, company.CIK, company.Exchange, 
		company.Currency, company.Country, company.Sector, company.Industry, company.Address, company.OfficialSite, 
		company.FiscalYearEnd, company.LatestQuarter, company.MarketCapitalization, company.EBITDA, company.PERatio, 
		company.PEGRatio, company.BookValue, company.DividendPerShare, company.DividendYield, company.EPS, 
		company.RevenuePerShareTTM, company.ProfitMargin, company.OperatingMarginTTM, company.ReturnOnAssetsTTM, 
		company.ReturnOnEquityTTM, company.RevenueTTM, company.GrossProfitTTM, company.DilutedEPSTTM, 
		company.QuarterlyEarningsGrowthYOY, company.QuarterlyRevenueGrowthYOY, company.AnalystTargetPrice, 
		company.AnalystRatingStrongBuy, company.AnalystRatingBuy, company.AnalystRatingHold, company.AnalystRatingSell, 
		company.AnalystRatingStrongSell, company.TrailingPE, company.ForwardPE, company.PriceToSalesRatioTTM, 
		company.PriceToBookRatio, company.EVToRevenue, company.EVToEBITDA, company.Beta, 
		company.FiftyTwoWeekHigh, company.FiftyTwoWeekLow, company.FiftyDayMovingAverage, 
		company.TwoHundredDayMovingAverage, company.SharesOutstanding, company.DividendDate, company.ExDividendDate,
	)


	if err != nil {
		log.Fatalf("Failed to insert data %v", err)
	}	

	fmt.Println("Data inserted successfully!")
	return company.Symbol, nil
	
}

func Insert_Time_Series_Weekly_data(comp string) (string, error) {
	dbConnectionPointer, error := ConnectDB()
	if error != nil {
		log.Fatal(error)
	}

	company, geterror := pkg.Get_Time_Series_Weekly_data(comp)
	if geterror != nil {
		fmt.Println("There has been some error: ", geterror)
		log.Fatal(geterror)
	}

	valueStrings := []string{}
	valueArgs := []interface{}{}
	i := 1
	
	for _, rec := range company {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d)", i, i+1, i+2, i+3, i+4, i+5, i+6))
		valueArgs = append(valueArgs,
			rec.Symbol,
			rec.Date,
			rec.OpenPrice,
			rec.HighPrice, 
			rec.LowPrice, 
			rec.ClosePrice,
			rec.Volume,
		)
		i += 7
	}

	query := fmt.Sprintf(
		`INSERT INTO time_series_weekly (Symbol, Date, Open_Price, High_Price, Low_Price, Close_Price, Volume) VALUES %s`,
		strings.Join(valueStrings, ","),

	)
	
	_, err := dbConnectionPointer.Exec(query, valueArgs...)
	if err != nil {
		log.Fatalf("Failed to bulk insert data %v", err)
	}

	return "data inserted successfully", nil
	
}