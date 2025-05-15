package services

import (
    "fmt"

    "backand/config"
)

// ClosePoll menutup polling jika polling ditemukan
func ClosePoll(pollID string) error {
    var exists bool
    err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM polls WHERE id = ?)", pollID).Scan(&exists)
    if err != nil {
        return fmt.Errorf("failed to check poll existence: %v", err)
    }
    if !exists {
        return fmt.Errorf("poll not found")
    }

    _, err = config.DB.Exec("UPDATE polls SET status = 'closed' WHERE id = ?", pollID)
    if err != nil {
        return fmt.Errorf("failed to close poll: %v", err)
    }

    return nil
}

// DeletePoll menghapus polling lengkap dengan data terkait (votes, options)
func DeletePoll(pollID string) error {
    var exists bool
    err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM polls WHERE id = ?)", pollID).Scan(&exists)
    if err != nil {
        return fmt.Errorf("failed to check poll existence: %v", err)
    }
    if !exists {
        return fmt.Errorf("poll not found")
    }

    tx, err := config.DB.Begin()
    if err != nil {
        return err
    }

    // Hapus votes
    _, err = tx.Exec("DELETE FROM votes WHERE poll_id = ?", pollID)
    if err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to delete votes: %v", err)
    }

    // Hapus opsi polling
    _, err = tx.Exec("DELETE FROM poll_options WHERE poll_id = ?", pollID)
    if err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to delete poll options: %v", err)
    }

    // Hapus polling utama
    _, err = tx.Exec("DELETE FROM polls WHERE id = ?", pollID)
    if err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to delete poll: %v", err)
    }

    return tx.Commit()
}