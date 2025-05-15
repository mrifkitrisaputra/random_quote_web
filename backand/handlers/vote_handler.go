package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "backand/services"
)

func CastVote(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    pollID := vars["id"]

    var input struct {
        OptionID int    `json:"option_id"`
        UserKey  string `json:"user_key"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    err := services.CastVote(pollID, input.OptionID, input.UserKey)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Vote successful"})
}