package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql" // MySQL driver
)

func main() {
    // Database connection string
    // Format: "user:password@tcp(host:port)/dbname"
    db, err := sql.Open("mysql", "root:changeme@tcp(127.0.0.1:3306)/myapp")
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
    defer db.Close() // Close the connection when main exits

    // Test the connection
    err = db.Ping()
    if err != nil {
        log.Fatal("Error pinging the database:", err)
    }
    fmt.Println("Successfully connected to the database!")

    // Query the database (example: selecting from a 'users' table)
    rows, err := db.Query("SELECT person_id  FROM person order by person_id")
    if err != nil {
        log.Fatal("Error executing query:", err)
    }
    defer rows.Close() // Close the result set when done

    // Iterate over the results
    for rows.Next() {
        var person_id int
        err := rows.Scan(&person_id)
        if err != nil {
            log.Fatal("Error scanning row:", err)
        }
        fmt.Printf("PERSON_ID: %d \n", person_id)
    }

    // Check for errors from iterating over rows
    if err = rows.Err(); err != nil {
        log.Fatal("Error during row iteration:", err)
    }
}
