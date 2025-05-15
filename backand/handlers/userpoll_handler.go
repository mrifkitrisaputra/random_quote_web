package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "backand/models"
    "backand/services"
)

func ViewPollPage(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    pollID := vars["id"]

    // 1. Ambil data polling
    poll, err := services.FetchPollData(pollID)
    if err != nil {
        http.Error(w, `{"error": "poll not found"}`, http.StatusNotFound)
        return
    }

    // 2. Ambil daftar opsi jawaban
    options, err := services.FetchPollOptions(pollID)
    if err != nil {
        http.Error(w, `{"error": "failed to fetch options"}`, http.StatusInternalServerError)
        return
    }

    // 3. Cek apakah peserta bisa lihat hasil
    canSeeResults := poll.AllowRealtimeResults || poll.Status == "closed"

    // 4. Jika boleh → ambil hasil voting
    var results []models.PollResult
    if canSeeResults {
        results, err = services.ExportPollResults(pollID)
        if err != nil {
            http.Error(w, `{"error": "failed to fetch results"}`, http.StatusInternalServerError)
            return
        }
    }

    // 5. Siapkan response
    response := models.PollDetails{
        Poll:            *poll,
        Options:         options,
        Results:         results,
        CanSeeResults:   canSeeResults,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}