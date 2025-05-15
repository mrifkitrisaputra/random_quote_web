package models

// PollStats adalah statistik polling
type PollStats struct {
    PollID     string `json:"poll_id"`
    TotalVotes int    `json:"total_votes"`
    LastUpdated string `json:"last_updated"`
}