package models

import "time"

// PollModel represents a poll in the system
type PollModel struct {
    ID          string    `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Options     []string  `json:"options"`
    IsOpen      bool      `json:"is_open"`
    CreatedAt   time.Time `json:"created_at"`
}

// PollStatsModel represents poll statistics
type PollStatsModel struct {
    PollData PollModel `json:"poll_data"`
    Votes    []struct {
        VoterID        string `json:"voter_id"`
        SelectedOption string `json:"selected_option"`
    } `json:"votes"`
}