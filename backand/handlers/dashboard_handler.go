package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "backand/services"
)

func GetDashboardData(w http.ResponseWriter, r *http.Request) {
    polls, err := services.GetPollsForDashboard()
    if err != nil {
        http.Error(w, `{"error": "Gagal memuat polling"}`, http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "polls": polls,
    })
}

func GetPollDetails(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    pollID := vars["id"]

    poll, err := services.GetPollDetails(pollID)
    if err != nil {
        http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(poll)
}