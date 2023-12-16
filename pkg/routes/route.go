package routes

import (
	"github.com/gorilla/mux"
	"github.com/Hosein110011/go-radar/pkg/controllers"
)


var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/users", controllers.GetProfile).Methods("GET")
	router.HandleFunc("/rooms", controllers.GetRooms).Methods("GET")
	router.HandleFunc("/jwt", controllers.JWT).Methods("GET")
	router.HandleFunc("/games", controllers.GetGames).Methods("GET")
	router.HandleFunc("/profile", controllers.GetUserProfile).Methods("GET")
}