package models

import "time"

type Poll struct {
    ID                   string    `json:"id"`
    Title                string    `json:"title"`
    Description          string    `json:"description"`
    CreatedAt            time.Time `json:"created_at"`
    ExpiresAt            time.Time `json:"expires_at"`
    Status               string    `json:"status"` // open / closed
    AllowRealtimeResults bool      `json:"allow_realtime"`
}