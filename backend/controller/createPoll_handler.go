package controller

import (
    "encoding/json"
    "net/http"
    "backend/service"
)

type CreatePollRequest struct {
    Title       string   `json:"title"`
    Description string   `json:"description"`
    Options     []string `json:"options"`
}

func CreatePollHandler(w http.ResponseWriter, r *http.Request) {
    var req CreatePollRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validate poll data
    if len(req.Title) == 0 || len(req.Options) < 2 {
        http.Error(w, "Title is required and at least two options are needed", http.StatusBadRequest)
        return
    }

    // Call service to create poll
    pollService := service.NewCreatePollService()
    pollDetails, err := pollService.CreatePoll(req.Title, req.Description, req.Options)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Return response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "pollID": pollDetails.PollID,
        "pollURL": pollDetails.PollURL,
    })
}