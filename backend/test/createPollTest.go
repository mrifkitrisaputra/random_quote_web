package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

const baseURL = "http://localhost:8080"

// Struct untuk request body
type CreatePollRequest struct {
    Title       string   `json:"title"`
    Description string   `json:"description"`
    Options     []string `json:"options"`
}

func sendCreatePollRequest(reqBody CreatePollRequest) {
    // Tampilkan request payload ke user
    requestJSON, _ := json.MarshalIndent(reqBody, "", "  ")
    fmt.Printf("\n📤 Sending Request:\n%s\n", string(requestJSON))

    // Kirim request ke server
    bodyJSON, _ := json.Marshal(reqBody)
    req, err := http.NewRequest("POST", baseURL+"/api/polls/create", bytes.NewBuffer(bodyJSON))
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

func runCreatePollTestCases() {
    for {
        fmt.Println("\n==============================")
        fmt.Println("🧪 Test Case: Create Poll")
        fmt.Println("==============================")
        fmt.Println("1. Input Valid")
        fmt.Println("2. Tanpa Judul")
        fmt.Println("3. Hanya Satu Opsi")
        fmt.Println("0. Kembali ke Menu Utama")
        fmt.Print("Masukkan pilihan Anda: ")

        var choice int
        _, err := fmt.Scan(&choice)
        if err != nil {
            fmt.Println("❌ Input tidak valid. Harap masukkan angka.")
            continue
        }

        switch choice {
        case 1:
            sendCreatePollRequest(CreatePollRequest{
                Title:       "Apa hobi kamu?",
                Description: "Pilih satu hobi yang paling kamu sukai",
                Options:     []string{"Membaca", "Olahraga", "Menonton Film"},
            })
        case 2:
            sendCreatePollRequest(CreatePollRequest{
                Title:       "",
                Description: "Pilih satu hobi yang paling kamu sukai",
                Options:     []string{"Opsi 1", "Opsi 2"},
            })
        case 3:
            sendCreatePollRequest(CreatePollRequest{
                Title:       "Poll tanpa cukup opsi",
                Description: "",
                Options:     []string{"Hanya satu opsi"},
            })
        case 0:
            fmt.Println("👋 Kembali ke menu utama.")
            return
        default:
            fmt.Println("⚠️ Pilihan tidak dikenali.")
        }
    }
}