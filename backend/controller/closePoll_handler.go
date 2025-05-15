package controller

import (
	"backend/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func ClosePollHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    pollID := vars["pollID"]

    userID := r.Header.Get("X-User-ID") // Ambil dari header auth (opsional)

    err := service.ClosePoll(pollID, userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]bool{"success": true})
}