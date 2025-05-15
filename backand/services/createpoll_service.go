package services

import (
    "crypto/rand"
    "fmt"
    "time"

    "backand/config"
)

func CreatePoll(title, description, expiresAt string, options []string, allowRealtime bool) (string, string, error) {
    pollID := generatePollID()
    now := time.Now().Format("2006-01-02 15:04:05")

    expires, err := time.Parse("2006-01-02 15:04:05", expiresAt)
    if err != nil {
        return "", "", fmt.Errorf("invalid expiration date")
    }

    tx, err := config.DB.Begin()
    if err != nil {
        return "", "", err
    }

    _, err = tx.Exec(`
        INSERT INTO polls (
            id, title, description, created_at, expires_at, status, allow_realtime_results
        ) VALUES (?, ?, ?, ?, ?, ?, ?)`,
        pollID, title, description, now, expires, "open", allowRealtime,
    )
    if err != nil {
        tx.Rollback()
        return "", "", err
    }

    for i, opt := range options {
        _, err = tx.Exec("INSERT INTO poll_options (poll_id, option_text, order_index) VALUES (?, ?, ?)",
            pollID, opt, i)
        if err != nil {
            tx.Rollback()
            return "", "", err
        }
    }

    err = tx.Commit()
    if err != nil {
        return "", "", err
    }

    url := fmt.Sprintf("https://quickpoll/polls/%s", pollID)
    return pollID, url, nil
}

func generatePollID() string {
    b := make([]byte, 8)
    rand.Read(b)
    return fmt.Sprintf("%x", b)
}