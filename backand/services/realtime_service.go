package services

import (
    "fmt"

    "backand/config"
)

func UpdateAllowRealtime(pollID string, allow bool) error {
    // Cek apakah polling masih terbuka
    var status string
    err := config.DB.QueryRow("SELECT status FROM polls WHERE id = ?", pollID).Scan(&status)
    if err != nil {
        return fmt.Errorf("poll not found")
    }

    if status != "open" {
        return fmt.Errorf("only open polls can change realtime visibility")
    }

    // Update kolom allow_realtime_results
    _, err = config.DB.Exec("UPDATE polls SET allow_realtime_results = ? WHERE id = ?", allow, pollID)
    if err != nil {
        return err
    }

    return nil
}