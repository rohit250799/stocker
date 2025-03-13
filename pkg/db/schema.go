package db

import (
	"database/sql"
	"fmt"
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

func CreateTable(db *sql.DB) error {
	fmt.Println("Enter the name of the table: ")

	var get_table_name string
	fmt.Scanln(&get_table_name)

	fmt.Println("Enter the number of fields: ")
	var num_of_fields int
	fmt.Scanln(&num_of_fields)

	for i := 0; i < num_of_fields; i++ {
		num_of_fields ++
	}

	query := fmt.Sprintf(`
	CREATE TABLE %s (
		id SERIAL PRIMARY KEY,
		stock_name varchar(30),
		stock_symbol varchar(10),
		closing_stock_price INT
	);,`, get_table_name)
    
	_, err := db.Exec(query);
	if err != nil {
		return err;
	}
	fmt.Println("Demo table has been created successfully");
	return nil
}

func perform_db_operations(user User) ([]User, error){

	db, err := ConnectDB();
	if err != nil {
		fmt.Println("There were some internal problems")
		//return err
	}
	//performing operations in db
	// db_custom_users := User{};
	rows, err := db.Query("SELECT username, email FROM USERS");
	if err != nil {
		return nil, err;
	}

	defer rows.Close();

	var db_custom_users []User;

	for rows.Next() {
		var database_user User;
		if err := rows.Scan(&database_user.Name, &database_user.Email); err != nil {
			return db_custom_users, err; 
		}
		db_custom_users = append(db_custom_users, database_user)
	}

	if err = rows.Err(); err != nil {
		return db_custom_users, err
	}

	return db_custom_users, nil
}