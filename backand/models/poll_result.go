package models

type PollResult struct {
    OptionText string `json:"option_text"`
    TotalVotes int    `json:"total_votes"`
}