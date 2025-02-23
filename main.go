package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"database/sql"

	_ "github.com/lib/pq"
)

//for PG-admin
type User struct {
    Name string `db:"username"`
    Email string `db:"email"`
}

func connectDB() (*sql.DB, error) {
    db_host := os.Getenv("DB_HOST");
    db_port := os.Getenv("DB_PORT");
    db_user := os.Getenv("POSTGRES_USER")
    db_password := os.Getenv("POSTGRES_PASSWORD")
    db_name := os.Getenv("POSTGRES_DB")

    

    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        db_host, db_port, db_user, db_password, db_name);
    db, err := sql.Open("postgres", connStr);
    if err != nil {
        log.Fatalln(err);
    }

    //defer db.Close();

    //Testing connection to db
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    } else {
        log.Println("Successfully connected to db");
    }

    return db, nil;
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


func main() {

    //to find the .env file
    err:= godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading to .env file: %s", err)
    }

    //connectDB()

    db, err := connectDB();
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()


    user := User{Name: "hackerrohit6", Email: "hackerrohit6@gmail.com"}

    db_custom_users, err := perform_db_operations(user);
    if err != nil {
        fmt.Println("Error performing operations: ", err);
        return
    }

    fmt.Println("All database users: ", user.Name);

    for _, database_user := range db_custom_users {
        fmt.Printf("Email: %s, Name: %s", database_user.Email, database_user.Name);
    }
}

