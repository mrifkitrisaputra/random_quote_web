package services

import (
    "fmt"

    "backand/config"
)

func CastVote(pollID string, optionID int, userKey string) error {
    var count int
    err := config.DB.QueryRow("SELECT COUNT(*) FROM poll_options WHERE id = ? AND poll_id = ?", optionID, pollID).Scan(&count)
    if err != nil {
        return err
    }
    if count == 0 {
        return fmt.Errorf("invalid option_id for this poll")
    }

    err = config.DB.QueryRow("SELECT COUNT(*) FROM votes WHERE poll_id = ? AND user_key = ?", pollID, userKey).Scan(&count)
    if err != nil {
        return err
    }
    if count > 0 {
        return fmt.Errorf("already voted")
    }

    _, err = config.DB.Exec("INSERT INTO votes (poll_id, option_id, user_key) VALUES (?, ?, ?)", pollID, optionID, userKey)
    return err
}