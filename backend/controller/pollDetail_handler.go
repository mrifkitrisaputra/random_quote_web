package controller

import (
	"encoding/json"
	"net/http"

	"backend/service"

	"github.com/gorilla/mux"
)

func GetPollDetailHandler(w http.ResponseWriter, r *http.Request) {
    // Ambil pollID dari URL
    vars := mux.Vars(r)
    pollID := vars["pollID"]

    // Panggil service
    poll, err := service.GetPollDetails(pollID)
    if err != nil {
        http.Error(w, "Poll tidak ditemukan", http.StatusNotFound)
        return
    }

    // Kirim respons JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(poll)
}