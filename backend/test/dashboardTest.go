package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

const dashboardBaseURL = "http://localhost:8080"

type DashboardPollRequest struct {
    Title       string   `json:"title"`
    Description string   `json:"description"`
    Options     []string `json:"options"`
}

func sendCreatePoll() {
    fmt.Println("\n🧪 [Test] Membuat poll baru")

    requestBody := DashboardPollRequest{
        Title:       "Apa hobi kamu?",
        Description: "Pilih satu hobi yang paling kamu sukai",
        Options:     []string{"Membaca", "Olahraga", "Menonton Film"},
    }

    requestJSON, _ := json.MarshalIndent(requestBody, "", "  ")
    fmt.Printf("📤 Sending Request:\n%s\n", string(requestJSON))

    bodyJSON, _ := json.Marshal(requestBody)
    req, err := http.NewRequest("POST", dashboardBaseURL+"/api/polls/create", bytes.NewBuffer(bodyJSON))
    if err != nil {
        fmt.Printf("❌ Error creating request: %v\n", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("❌ Error sending request: %v\n", err)
        return
    }
    defer resp.Body.Close()

    bodyBytes, _ := io.ReadAll(resp.Body)
    bodyString := string(bodyBytes)

    fmt.Printf("🔹 Status Code: %d\n", resp.StatusCode)
    fmt.Printf("📦 Response Body:\n%s\n", bodyString)
}

func sendGetUserPolls() {
    fmt.Println("\n🧪 [Test] Mengambil daftar poll milik user")
    fmt.Println("📤 Request: GET /api/dashboard?user_id=user123")

    req, err := http.NewRequest("GET", dashboardBaseURL+"/api/dashboard?user_id=user123", nil)
    if err != nil {
        fmt.Printf("❌ Error creating request: %v\n", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("❌ Error sending request: %v\n", err)
        return
    }
    defer resp.Body.Close()

    bodyBytes, _ := io.ReadAll(resp.Body)
    bodyString := string(bodyBytes)

    fmt.Printf("🔹 Status Code: %d\n", resp.StatusCode)
    fmt.Printf("📦 Response Body:\n%s\n", bodyString)
}

func showMenuDashboard() int {
    fmt.Println("\n==============================")
    fmt.Println("🧪 Test Menu: Dashboard")
    fmt.Println("==============================")
    fmt.Println("1. Buat Poll Baru")
    fmt.Println("2. Ambil Daftar Poll Milik User")
    fmt.Println("0. Kembali ke Menu Utama")
    fmt.Print("Masukkan pilihan Anda: ")

    var choice int
    _, err := fmt.Scan(&choice)
    if err != nil {
        fmt.Println("❌ Input tidak valid. Harap masukkan angka.")
        return -1
    }

    return choice
}

func runDashboardTester() {
    for {
        choice := showMenuDashboard()
        if choice == -1 {
            continue
        }

        switch choice {
        case 1:
            sendCreatePoll()
        case 2:
            sendGetUserPolls()
        case 0:
            fmt.Println("👋 Kembali ke menu utama.")
            return
        default:
            fmt.Println("⚠️ Pilihan tidak dikenali.")
        }
    }
}