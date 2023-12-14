package routes

import (
	"github.com/gorilla/mux"
	"github.com/Hosein110011/go-radar/pkg/controllers"
)


var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/users", controllers.GetProfile).Methods("GET")
}