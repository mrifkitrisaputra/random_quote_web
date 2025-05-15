package models

import "time"

// Vote merepresentasikan suara dari pengguna
type Vote struct {
    ID        int       `json:"id"`
    PollID    string    `json:"poll_id"`
    OptionID  int       `json:"option_id"`
    UserKey   string    `json:"user_key"`
    VotedAt   time.Time `json:"voted_at"`
}

// VoteResult adalah hasil perhitungan suara
type VoteResult struct {
    OptionText string `json:"option_text"`
    TotalVotes int    `json:"total_votes"`
}