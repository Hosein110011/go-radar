package controllers

import (
	"encoding/json"
	"fmt"
	// "github.com/gorilla/mux"
	"net/http"
	// "strconv"
	// "github.com/Hosein1100011/go-radar/pkg/utils"
	"github.com/Hosein110011/go-radar/pkg/models"
	"github.com/Hosein110011/go-radar/pkg/utils"

)


func GetProfile(w http.ResponseWriter, r *http.Request) {
	profiles := models.GetAllUsers()
	res, _ := json.Marshal(profiles)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms := models.GetAllRooms()
	res, _ := json.Marshal(rooms)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func GetGames(w http.ResponseWriter, r *http.Request) {
	games := models.GetAllGames()
	res, _ := json.Marshal(games)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}


func JWT(w http.ResponseWriter, r *http.Request) {
    token, err := utils.GetTokenFromHeader(r)
    if err != nil {
        // Handle the error
        fmt.Println(err)
        return
    }

    if claims, ok := token.Claims.(*utils.Claims); ok && token.Valid {
        // Use claims
        fmt.Println(claims.Username)
    } else {
        fmt.Println("Invalid token")
    }
}


func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	token, err := utils.GetTokenFromHeader(r)
    if err != nil {
        // Handle the error
        fmt.Println(err)
        return
    }

    if claims, ok := token.Claims.(*utils.Claims); ok && token.Valid {
        // Use claims
        user := models.GetUserByUsername(claims.Username)
		res, _ := json.Marshal(user)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
    } else {
        fmt.Println("Invalid token")
		return
    }

}