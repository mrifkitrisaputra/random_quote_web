package models

import "time"

type Link struct {
    Token     string    `json:"token"`
    PollID    string    `json:"pollId"`
    CreatedAt time.Time `json:"createdAt"`
    IsActive  bool      `json:"isActive"`
}
