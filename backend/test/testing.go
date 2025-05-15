package main

import "fmt"

func showMainMenu() int {
    fmt.Println("\n==============================")
    fmt.Println("🧪 QuickPoll CLI Tester ")
    fmt.Println("==============================")
    fmt.Println("Pilih fitur yang ingin di-test:")
    fmt.Println("1. Test Create Poll")
    fmt.Println("2. Test Dashboard")
    fmt.Println("0. Keluar")
    fmt.Print("Masukkan pilihan Anda: ")

    var choice int
    _, err := fmt.Scan(&choice)
    if err != nil {
        fmt.Println("❌ Input tidak valid. Harap masukkan angka.")
        return -1
    }

    return choice
}

func main() {
    for {
        choice := showMainMenu()
        if choice == -1 {
            continue
        }

        switch choice {
        case 1:
            runCreatePollTestCases()
        case 2:
            runDashboardTester()
        case 0:
            fmt.Println("👋 Terima kasih. Keluar dari program.")
            return
        default:
            fmt.Println("⚠️ Pilihan tidak dikenali.")
        }
    }
}