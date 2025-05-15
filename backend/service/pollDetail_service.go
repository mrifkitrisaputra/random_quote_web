package service

import (
	"backend/config"
	"backend/models"
	"encoding/json"
)

func GetPollDetails(pollID string) (*models.PollModel, error) {
    var poll models.PollModel
    var optionsJSON []byte

    query := `
        SELECT id, title, description, options, is_open, created_at
        FROM polls
        WHERE id = ?
    `

    row := config.DB.QueryRow(query, pollID)
    err := row.Scan(
        &poll.ID,
        &poll.Title,
        &poll.Description,
        &optionsJSON,
        &poll.IsOpen,
        &poll.CreatedAt,
    )
    if err != nil {
        return nil, err
    }

    json.Unmarshal(optionsJSON, &poll.Options)

    return &poll, nil
}