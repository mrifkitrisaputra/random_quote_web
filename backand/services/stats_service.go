package services

import (
    "backand/models"
    "backand/config"
)

func ExportPollResults(pollID string) ([]models.PollResult, error) {
    rows, err := config.DB.Query(`
        SELECT po.option_text, COUNT(v.id) AS total_votes
        FROM poll_options po
        LEFT JOIN votes v ON po.id = v.option_id
        WHERE po.poll_id = ?
        GROUP BY po.id`, pollID)

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []models.PollResult
    for rows.Next() {
        var res models.PollResult
        if err := rows.Scan(&res.OptionText, &res.TotalVotes); err != nil {
            return nil, err
        }
        results = append(results, res)
    }

    return results, nil
}