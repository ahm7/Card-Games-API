package Routers

import (
	"main/Controller"

	"github.com/gorilla/mux"
)


func DefineRoutes(r *mux.Router) {

	// routes
	r.HandleFunc("/api/Deck", Controller.Create_Deck).Methods("POST")
	r.HandleFunc("/api/Deck/{id}",Controller.Open_Deck).Methods("GET")
	r.HandleFunc("/api/Deck/{id}/draw", Controller.Draw_Card).Methods("Post")
}