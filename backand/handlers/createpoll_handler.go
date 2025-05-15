package handlers

import (
    "encoding/json"
    "net/http"

    "backand/services"
)

func CreatePollHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Title                string   `json:"title"`
        Description          string   `json:"description"`
        Options              []string `json:"options"`
        ExpiresAt            string   `json:"expires_at"`
        AllowRealtimeResults bool     `json:"allow_realtime_results"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    pollID, pollURL, err := services.CreatePoll(
        input.Title,
        input.Description,
        input.ExpiresAt,
        input.Options,
        input.AllowRealtimeResults,
    )
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "poll_id": pollID,
        "url":     pollURL,
    })
}