package controllers

import (
	"encoding/json"
	"fmt"

	// "github.com/gorilla/mux"
	"net/http"
	// "strconv"
	// "github.com/Hosein1100011/go-radar/pkg/utils"
	"github.com/Hosein110011/go-radar/pkg/models"
	"github.com/Hosein110011/go-radar/pkg/schema"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	queryParams := r.URL.Query()
	// Get the user_id query parameter
	userID := queryParams.Get("user_id")
	// Check if user_id is provided
	if claims, ok := token.Claims.(*utils.Claims); ok && token.Valid {
		// Use claims
		user, err := models.GetUserByUsername(claims.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(user)
		var requestedUser *models.User
		if (userID != "") && (user.ID != userID) {
			requestedUser, err = models.GetUserByID(userID)
			if err != nil {
				http.Error(w, "user not found.", http.StatusNotFound)
				return
			}
		} else {
			requestedUser = user
		}
		response, err := schema.CreateProfileResponse(user, requestedUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		res, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else {
		fmt.Println("Invalid token")
		return
	}

}

func GetSquad(w http.ResponseWriter, r *http.Request) {
	token, err := utils.GetTokenFromHeader(r)
	if err != nil {
		// Handle the error
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if claims, ok := token.Claims.(*utils.Claims); ok && token.Valid {
		// Use claims
		user, err := models.GetUserByUsername(claims.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response, err := schema.CreateSquadResponse(user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		res, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else {
		fmt.Println("Invalid token")
		http.Error(w, "Invalid token", http.StatusInternalServerError)
		return
	}
}

func GetJoinReqs(w http.ResponseWriter, r *http.Request) {
	reqs := models.GetAllJoinReqs()
	res, _ := json.Marshal(reqs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	token, err := utils.GetTokenFromHeader(r)
	if err != nil {
		// Handle the error
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	queryParams := r.URL.Query()
	// Get the user_id query parameter
	userID := queryParams.Get("user_id")
	// Check if user_id is provided
	if claims, ok := token.Claims.(*utils.Claims); ok && token.Valid {
		// Use claims
		user, err := models.GetUserByUsername(claims.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(user)
		var recipientUser *models.User
		if (userID != "") && (user.ID != userID) {
			recipientUser, err = models.GetUserByID(userID)
			if err != nil {
				http.Error(w, "user not found.", http.StatusNotFound)
				return
			}
		} else {
			http.Error(w, "user_id not provided or you are recipient too.", http.StatusNotFound)
			return
		}
		fmt.Println(recipientUser)
		session := models.GetSession()
		messages := models.GetSortedMessages(session, userID, user.ID, 11)

		res, err := json.Marshal(messages)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else {
		fmt.Println("Invalid token")
		return
	}

}
