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