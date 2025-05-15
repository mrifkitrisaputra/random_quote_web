package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "backand/services"

)

func ExportPollResults(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    pollID := vars["id"]

    results, err := services.ExportPollResults(pollID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}