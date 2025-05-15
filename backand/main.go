package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "backand/config"
	"backand/routes"
)

func main() {
    config.Connect() // ✅ Harus dijalankan dulu!

    r := mux.NewRouter()
    routes.SetupRoutes(r)

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000", "http://localhost:5173"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type"},
    })

    handler := c.Handler(r)

    fmt.Println("Server running on :8080")
    http.ListenAndServe(":8080", handler)
}