package config

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    var err error
    dsn := "root:@tcp(127.0.0.1:3307)/db_quickpoll?parseTime=true"
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Error opening database:", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("Error connecting to database:", err)
    }

    fmt.Println("Connected to MySQL database")
}