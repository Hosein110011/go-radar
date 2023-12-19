package schema

import (
	"github.com/Hosein110011/go-radar/pkg/models"
)

type ApiResponse struct {
	Message    string      `json:"message"`
	IsSuccess  bool        `json:"isSuccess"`
	StatusCode int         `json:"statusCode"`
	Result     interface{} `json:"result"`
}

type UserResult struct {
	ID           string            `json:"id"`
	Nickname     string            `json:"nickname"`
	Username     string            `json:"username"`
	Photo        *string           `json:"photo"` // Use pointer to allow null
	IsReady      bool              `json:"is_ready"`
	IsOnline     bool              `json:"is_online"`
	IsMine       bool              `json:"is_mine"`
	IsFriend     bool              `json:"is_friend"`
	Bio          string            `json:"bio"`
	RoomID       string            `json:"room_id"` // Use pointer to allow null
	LikeStatus   string            `json:"like_status"`
	LikeCount    int               `json:"like_count"`
	DislikeCount int               `json:"dislike_count"`
	LikedGames   []GameApiResponse `json:"liked_games"`
}

type GameApiResponse struct {
	ID          string `json:"id"`
	GameName    string `json:"game_name"`
	PackageName string `json:"package_name"`
	Image       string `json:"image"`
	Banner      string `json:"banner"`
	Platform    string `json:"platform"`
	IsDeleted   bool   `json:"is_deleted"`
}

func CreateProfileResponse(user, requestedUser *models.User) ApiResponse {
	var userResult UserResult
	var Games []models.Game
	var FavouriteGames []GameApiResponse

	userResult.ID = requestedUser.ID
	userResult.Nickname = requestedUser.Nickname
	userResult.Username = requestedUser.Username
	userResult.IsReady = requestedUser.IsReady
	userResult.IsOnline = requestedUser.IsOnline
	Games, _ = models.GetFavouriteGamesByUserID(requestedUser.ID)

	for _, game := range Games {
		var FavouriteGame GameApiResponse
		FavouriteGame.ID = game.ID
		FavouriteGame.GameName = game.GameName
		FavouriteGame.PackageName = game.PackageName
		FavouriteGame.Image = game.Image
		FavouriteGame.Banner = game.Banner
		FavouriteGame.Platform = game.Platform
		FavouriteGame.IsDeleted = game.IsDeleted
		FavouriteGames = append(FavouriteGames, FavouriteGame)
	}

	userResult.LikedGames = FavouriteGames

	// Assuming bio, photo, and room_id are nullable
	if requestedUser.Bio != "" {
		userResult.Bio = requestedUser.Bio
	}

	if requestedUser.Photo != "" {
		userResult.Photo = &requestedUser.Photo
	}

	if user.ID == requestedUser.ID {
		userResult.IsMine = true
		userResult.IsFriend = false
	} else {
		userResult.IsMine = false
		for _, friend := range requestedUser.Friends {
			if friend.ID == user.ID {
				userResult.IsFriend = true
			} else {
				userResult.IsFriend = false
			}
		}
	}
	for _, like := range requestedUser.Likes {
		if user.ID == like.ID {
			userResult.LikeStatus = "liked"
			break
		}
	}
	if userResult.LikeStatus == "" {
		for _, dislike := range requestedUser.Dislikes {
			if user.ID == dislike.ID {
				userResult.LikeStatus = "disliked"
				break
			} else {
				userResult.LikeStatus = "none"
				break
			}
		}
	}

	if len(requestedUser.Rooms) != 0 {
		userResult.RoomID = requestedUser.Rooms[0].ID
	}
	userResult.LikeCount = len(requestedUser.Likes)
	userResult.DislikeCount = len(requestedUser.Dislikes)
	// userResult.LikedGames = []string{} // Populate as needed

	return ApiResponse{
		Message:    "The user.",
		IsSuccess:  true,
		StatusCode: 200,
		Result:     userResult,
	}
}

//---------------------------------------------

// type SquadApiResponse struct {
// 	ID           string                   `json:"id"`
// 	RoomName     string                   `json:"room_name"`
// 	Owner        string                   `json:"owner"`
// 	Game         string                   `json:"game"`
// 	MemberLimit  int                      `json:"member_limit"`
// 	Member       []AccountApiResponse     `json:"member"`
// 	IsOwner      bool                     `json:"is_owner"`
// 	Created      time.Time                `json:"created"`
// 	UserID       string                   `json:"user_id"`
// 	JoinRequests []JoinRequestApiResponse `json:"join_requests"`
// }

// type AccountApiResponse struct {
// 	ID       string `json:"id"`
// 	Nickname string `json:"nickname"`
// 	Username string `json:"username"`
// 	Photo    string `json:"photo"`
// }

// func CreateSquadResponse(userID string) (ApiResponse, error) {
// 	var rooms []models.Room

// 	rooms, err := models.FindRoomByUser(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// }
