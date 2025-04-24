package models

import "time"

type Vote struct {
    ID           string    `json:"id"`
    PollID       string    `json:"pollId"`
    OptionID     string    `json:"optionId"`
    UserIdentifier string  `json:"userIdentifier"`
    Fingerprint  string    `json:"fingerprint"`
    Timestamp    time.Time `json:"timestamp"`
}
