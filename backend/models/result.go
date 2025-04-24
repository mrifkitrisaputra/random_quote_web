package models

type Result struct {
    PollID     string                    `json:"pollId"`
    TotalVotes int                       `json:"totalVotes"`
    OptionData map[string]OptionResult   `json:"optionData"`
}
