package routes

import (
    "github.com/gorilla/mux"
    "backend/controller"
)

func SetupRoutes(r *mux.Router) {
    r.HandleFunc("/api/polls/create", controller.CreatePollHandler).Methods("POST")
	r.HandleFunc("/api/dashboard", controller.HandleDashboardRequest).Methods("GET")
	r.HandleFunc("/api/polls/{pollID}/stats", controller.HandlePollStatsRequest).Methods("GET")
    r.HandleFunc("/api/polls/{pollID}", controller.GetPollDetailHandler).Methods("GET")
    r.HandleFunc("/api/polls/{pollID}/close", controller.ClosePollHandler).Methods("PUT")
}