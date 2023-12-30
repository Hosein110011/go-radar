package routes

import (
	"github.com/Hosein110011/go-radar/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/users", controllers.GetProfile).Methods("GET")
	router.HandleFunc("/rooms", controllers.GetRooms).Methods("GET")
	router.HandleFunc("/jwt", controllers.JWT).Methods("GET")
	router.HandleFunc("/games", controllers.GetGames).Methods("GET")
	router.HandleFunc("/reqs", controllers.GetJoinReqs).Methods("GET")
	router.HandleFunc("/api/v1/profile/", controllers.GetUserProfile).Methods("GET")
	router.HandleFunc("/api/v1/squad/", controllers.GetSquad).Methods("GET")
	router.HandleFunc("/api/v1/messages/", controllers.GetMessages).Methods("GET")

}
