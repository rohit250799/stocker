package main

import (
	"database/sql"
	"fmt"
)

func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE demo_stock_info (
		id SERIAL PRIMARY KEY,
		stock_name varchar(30),
		stock_symbol varchar(10),
		closing_stock_price INT
	);`
    
	_, err := db.Exec(query);
	if err != nil {
		return err;
	}
	fmt.Println("Demo table has been created successfully");
	return nil
}

func perform_db_operations(user User) ([]User, error){

	db, err := connectDB();
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