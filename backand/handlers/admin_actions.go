package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "backand/services"
)

func ClosePoll(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    pollID := vars["id"]

    err := services.ClosePoll(pollID)
    if err != nil {
        if err.Error() == "poll not found" {
            http.Error(w, `{"error": "poll not found"}`, http.StatusNotFound)
        } else {
            http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
        }
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Poll closed successfully",
    })
}

func DeletePoll(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    pollID := vars["id"]

    err := services.DeletePoll(pollID)
    if err != nil {
        if err.Error() == "poll not found" {
            http.Error(w, `{"error": "poll not found"}`, http.StatusNotFound)
        } else {
            http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
        }
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Poll deleted successfully",
    })
}