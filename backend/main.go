package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "backend/config"
    "backend/routes"
)

func main() {
    config.InitDB()

    r := mux.NewRouter()
    routes.SetupRoutes(r)

    // Middleware CORS
    corsHandler := func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
            w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

            if r.Method == "OPTIONS" {
                w.WriteHeader(http.StatusOK)
                return
            }

            next.ServeHTTP(w, r)
        })
    }

    // Pastikan kamu kirim router Gorilla Mux ke corsHandler
    log.Println("Server running on :8080")
    err := http.ListenAndServe(":8080", corsHandler(r))
    if err != nil {
        log.Fatal(err)
    }
}