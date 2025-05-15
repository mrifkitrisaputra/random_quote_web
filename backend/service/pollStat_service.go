package service

import (
	"encoding/json"
	"fmt"

	"backend/config"
	"backend/models"
)

func GetPollStats(pollID string) (*models.PollStatsModel, error) {
    var pollData models.PollModel
    var votes []struct {
        VoterID        string `json:"voter_id"`
        SelectedOption string `json:"selected_option"`
    }

    // Fetch poll details
    query := `
        SELECT id, title, description, options, is_open, created_at
        FROM polls
        WHERE id = ?
    `
    row := config.DB.QueryRow(query, pollID)

    var optionsJSON []byte
    err := row.Scan(
        &pollData.ID,
        &pollData.Title,
        &pollData.Description,
        &optionsJSON,
        &pollData.IsOpen,
        &pollData.CreatedAt,
    )
    if err != nil {
        return nil, fmt.Errorf("error fetching poll details: %w", err)
    }

    json.Unmarshal(optionsJSON, &pollData.Options)

    // Fetch votes for this poll
    voteQuery := `
        SELECT voter_id, selected_option
        FROM votes
        WHERE poll_id = ?
    `
    rows, err := config.DB.Query(voteQuery, pollID)
    if err != nil {
        return nil, fmt.Errorf("error fetching votes: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var vote struct {
            VoterID        string `json:"voter_id"`
            SelectedOption string `json:"selected_option"`
        }

        err := rows.Scan(&vote.VoterID, &vote.SelectedOption)
        if err != nil {
            return nil, fmt.Errorf("error scanning vote row: %w", err)
        }
        votes = append(votes, vote)
    }

    return &models.PollStatsModel{
        PollData: pollData,
        Votes:    votes,
    }, nil
}