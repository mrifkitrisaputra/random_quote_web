package service

import (
    "fmt"
    "backend/config"
)

func ClosePoll(pollID string, userID string) error {
    // Verifikasi pemilik poll
    isOwner, err := verifyPollOwnership(pollID, userID)
    if err != nil {
        return fmt.Errorf("gagal memverifikasi kepemilikan poll: %w", err)
    }

    if !isOwner {
        return fmt.Errorf("anda bukan pemilik poll")
    }

    // Tutup poll
    err = closePoll(pollID)
    if err != nil {
        return fmt.Errorf("gagal menutup poll: %w", err)
    }

    return nil
}

func verifyPollOwnership(pollID string, userID string) (bool, error) {
    // Query database untuk cek pemilik poll
    query := `
        SELECT user_id FROM polls
        WHERE id = ?
    `

    row := config.DB.QueryRow(query, pollID)
    var ownerID string
    err := row.Scan(&ownerID)
    if err != nil {
        return false, fmt.Errorf("gagal mengambil data pemilik poll: %w", err)
    }

    return ownerID == userID, nil
}

func closePoll(pollID string) error {
    // Update status poll menjadi is_open = false
    query := `
        UPDATE polls
        SET is_open = false
        WHERE id = ?
    `

    result, err := config.DB.Exec(query, pollID)
    if err != nil {
        return fmt.Errorf("gagal menutup poll: %w", err)
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("poll tidak ditemukan")
    }

    return nil
}