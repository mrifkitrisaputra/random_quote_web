package models

import "encoding/json"

type Poll struct {
    ID          string    `json:"id"`
	UserID      string    `json:"-"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Options     []string  `json:"options"`
    IsOpen      bool      `json:"is_open"`
    CreatedAt   string    `json:"created_at"`
}

type PollDetails struct {
    PollID  string `json:"poll_id"`
    PollURL string `json:"poll_url"`
}

// Fungsi bantu untuk serialize opsi ke JSON
func SerializeOptions(options []string) ([]byte, error) {
    return json.Marshal(options)
}

// Fungsi bantu untuk deserialize dari JSON ke []string
func DeserializeOptions(data []byte) ([]string, error) {
    var options []string
    err := json.Unmarshal(data, &options)
    return options, err
}