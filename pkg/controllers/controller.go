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
		response := schema.CreateProfileResponse(requestedUser)
		if user.ID == requestedUser.ID {
			response.Result.IsMine = true
			response.Result.IsFriend = false
		} else {
			response.Result.IsMine = false
			for _, friend := range requestedUser.Friends {
				if friend.ID == user.ID {
					response.Result.IsFriend = true
				} else {
					response.Result.IsFriend = false
				}
			}
		}
		for _, like := range requestedUser.Likes {
			if user.ID == like.ID {
				response.Result.LikeStatus = "liked"
				break
			}
		}
		if response.Result.LikeStatus == "" {
			for _, dislike := range requestedUser.Dislikes {
				if user.ID == dislike.ID {
					response.Result.LikeStatus = "disliked"
					break
				} else {
					response.Result.LikeStatus = "none"
					break
				}
			}
		}
		res, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else {
		fmt.Println("Invalid token")
		return
	}

}