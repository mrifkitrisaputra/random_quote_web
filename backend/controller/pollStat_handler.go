package controller

import (
	"encoding/json"
	"net/http"

	"backend/service"

	"github.com/gorilla/mux"
)

func HandlePollStatsRequest(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    pollID := vars["pollID"]

    stats, err := service.GetPollStats(pollID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(stats)
}