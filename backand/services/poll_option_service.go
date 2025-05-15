package services

import (
    "fmt"

    "backand/config"
	"backand/models"
)

func FetchPollOptions(pollID string) ([]models.PollOption, error) {
    rows, err := config.DB.Query("SELECT id, option_text, order_index FROM poll_options WHERE poll_id = ?", pollID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var options []models.PollOption
    for rows.Next() {
        var o models.PollOption
        if err := rows.Scan(&o.ID, &o.Text, &o.Order); err != nil {
            return nil, fmt.Errorf("failed to scan poll options: %v", err)
        }
        options = append(options, o)
    }

    return options, nil
}