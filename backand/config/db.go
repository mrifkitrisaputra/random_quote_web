package config

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
    var err error
    dsn := "root:@tcp(127.0.0.1:3307)/quiickpoll?parseTime=true"
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }

    err = DB.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Database connected")
}