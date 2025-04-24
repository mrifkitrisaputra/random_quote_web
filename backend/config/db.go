package config

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/quickpoll")
    if err != nil {
        log.Fatal("Database connection failed: ", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Database ping failed: ", err)
    }
    log.Println("Database connected")
}
