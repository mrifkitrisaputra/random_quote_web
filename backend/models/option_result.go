package models

type OptionResult struct {
    OptionID   string  `json:"optionId"`
    Text       string  `json:"text"`
    Votes      int     `json:"votes"`
    Percentage float64 `json:"percentage"`
}
