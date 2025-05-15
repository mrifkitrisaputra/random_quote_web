package routes

import (
    "github.com/gorilla/mux"
    "backand/handlers"
)

func SetupRoutes(r *mux.Router) {
    // Poll Creation
    r.HandleFunc("/api/polls", handlers.CreatePollHandler).Methods("POST")

    // Dashboard
    r.HandleFunc("/api/dashboard", handlers.GetDashboardData).Methods("GET")
    r.HandleFunc("/api/polls/{id}", handlers.ViewPollPage).Methods("GET")

    // Voting
    r.HandleFunc("/api/polls/{id}/vote", handlers.CastVote).Methods("POST")

    // Admin Actions
    r.HandleFunc("/api/polls/{id}/close", handlers.ClosePoll).Methods("PUT")
    r.HandleFunc("/api/polls/{id}", handlers.DeletePoll).Methods("DELETE")
    r.HandleFunc("/api/polls/{id}/export", handlers.ExportPollResults).Methods("GET")
    r.HandleFunc("/api/polls/{id}/toggle-realtime", handlers.ToggleRealtimeVisibility).Methods("PUT")
}