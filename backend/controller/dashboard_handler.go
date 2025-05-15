package controller

import (
    "encoding/json"
    "net/http"
    "backend/service"
    "backend/models"
)

func HandleDashboardRequest(w http.ResponseWriter, r *http.Request) {
    userID := "user123"
    polls, err := service.GetUserPolls(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if len(polls) == 0 {
        json.NewEncoder(w).Encode([]models.PollModel{})
        return
    }
    
    json.NewEncoder(w).Encode(polls)
}