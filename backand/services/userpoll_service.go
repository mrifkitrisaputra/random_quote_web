package services

import (
    "backand/models"
    "backand/config"
)

func FetchPollData(pollID string) (*models.Poll, error) {
    var p models.Poll
    row := config.DB.QueryRow(`
        SELECT id, title, description, expires_at, status, allow_realtime_results 
        FROM polls WHERE id = ?`, pollID)

    if err := row.Scan(&p.ID, &p.Title, &p.Description, &p.ExpiresAt, &p.Status, &p.AllowRealtimeResults); err != nil {
        return nil, err
    }

    return &p, nil
}