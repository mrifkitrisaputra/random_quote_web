package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "backand/services"
)

func ToggleRealtimeVisibility(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    pollID := vars["id"]

    var input struct {
        AllowRealtime bool `json:"allow_realtime"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    err := services.UpdateAllowRealtime(pollID, input.AllowRealtime)
    if err != nil {
        http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Realtime visibility updated",
    })
}