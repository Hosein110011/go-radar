package controllers

import (
	"encoding/json"
	// "fmt"
	// "github.com/gorilla/mux"
	"net/http"
	// "strconv"
	// "github.com/Hosein1100011/go-radar/pkg/utils"
	"github.com/Hosein110011/go-radar/pkg/models"
)


func GetProfile(w http.ResponseWriter, r *http.Request) {
	profiles := models.GetAllUsers()
	res, _ := json.Marshal(profiles)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}