package service

import (
	"encoding/json"
	"fmt"

	"backend/config"
	"backend/models"
)

func GetUserPolls(userID string) ([]models.PollModel, error) {
    var polls []models.PollModel

    query := `
        SELECT id, title, description, options, is_open, created_at
        FROM polls
        WHERE user_id = ?
    `
    rows, err := config.DB.Query(query, userID)
    if err != nil {
        return nil, fmt.Errorf("error fetching user polls: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var poll models.PollModel
        var optionsJSON []byte

        err := rows.Scan(
            &poll.ID,
            &poll.Title,
            &poll.Description,
            &optionsJSON,
            &poll.IsOpen,
            &poll.CreatedAt,
        )
        if err != nil {
            return nil, fmt.Errorf("error scanning row: %w", err)
        }

        json.Unmarshal(optionsJSON, &poll.Options)
        polls = append(polls, poll)
    }

    return polls, nil
}