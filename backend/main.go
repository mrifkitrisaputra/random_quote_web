package main

import (
    "backend/config"
    "backend/routes"
    "log"
)

func main() {
    // Inisialisasi koneksi ke database
    config.InitDB()

    // Setup router
    r := routes.SetupRouter()

    // Jalankan server
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Server failed to start: ", err)
    }
}
