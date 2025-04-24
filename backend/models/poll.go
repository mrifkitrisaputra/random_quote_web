package models

import "time"

type PollStatus string

const (
    Open   PollStatus = "open"
    Closed PollStatus = "closed"
)

type Poll struct {
    ID          string     `json:"id"`
    Title       string     `json:"title"`
    Description string     `json:"description"`
    CreatorID   string     `json:"creatorId"`
    CreatedAt   time.Time  `json:"createdAt"`
    ExpiresAt   time.Time  `json:"expiresAt"`
    Status      PollStatus `json:"status"`
}

type PollSettings struct {
    PollID                string `json:"pollId"`
    IsPrivate             bool   `json:"isPrivate"`
    ShowResultsBeforeClose bool  `json:"showResultsBeforeClose"`
    AllowMultipleAnswers  bool   `json:"allowMultipleAnswers"`
}