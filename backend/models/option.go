package models

type Option struct {
    ID     string `json:"id"`
    PollID string `json:"pollId"`
    Text   string `json:"text"`
}
