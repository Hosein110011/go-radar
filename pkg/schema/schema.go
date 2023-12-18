package schema

import (
	// "fmt"

	"github.com/Hosein110011/go-radar/pkg/models"
)

type ApiResponse struct {
	Message    string     `json:"message"`
	IsSuccess  bool       `json:"isSuccess"`
	StatusCode int        `json:"statusCode"`
	Result     UserResult `json:"result"`
}

type UserResult struct {
	ID           string        `json:"id"`
	Nickname     string        `json:"nickname"`
	Username     string        `json:"username"`
	Photo        *string       `json:"photo"` // Use pointer to allow null
	IsReady      bool          `json:"is_ready"`
	IsOnline     bool          `json:"is_online"`
	IsMine       bool          `json:"is_mine"`
	IsFriend     bool          `json:"is_friend"`
	Bio          string        `json:"bio"`
	RoomID       string        `json:"room_id"` // Use pointer to allow null
	LikeStatus   string        `json:"like_status"`
	LikeCount    int           `json:"like_count"`
	DislikeCount int           `json:"dislike_count"`
	LikedGames   []models.Game `json:"liked_games"`
}

func CreateProfileResponse(userData *models.User) ApiResponse {
	var userResult UserResult

	userResult.ID = userData.ID
	userResult.Nickname = userData.Nickname
	userResult.Username = userData.Username
	userResult.IsReady = userData.IsReady
	userResult.IsOnline = userData.IsOnline
	userResult.LikedGames, _ = models.GetFavouriteGamesByUserID(userData.ID)

	// Assuming bio, photo, and room_id are nullable
	if userData.Bio != "" {
		userResult.Bio = userData.Bio
	}

	if userData.Photo != "" {
		userResult.Photo = &userData.Photo
	}

	// Assuming logic for isMine, isFriend, likeStatus, likeCount, dislikeCount, likedGames
	// You need to replace these with actual logic or data
	// userResult.IsMine = true
	// userResult.IsFriend = false
	// userResult.LikeStatus = "disliked"
	if len(userData.Rooms) != 0 {
		userResult.RoomID = userData.Rooms[0].ID
	}
	userResult.LikeCount = len(userData.Likes)
	userResult.DislikeCount = len(userData.Dislikes)
	// userResult.LikedGames = []string{} // Populate as needed

	return ApiResponse{
		Message:    "The user.",
		IsSuccess:  true,
		StatusCode: 200,
		Result:     userResult,
	}
}
