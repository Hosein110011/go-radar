package schema

import (
	"github.com/Hosein110011/go-radar/pkg/models"
	"time"
	"fmt"
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
	Photo        string            `json:"photo"` // Use pointer to allow null
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

func CreateProfileResponse(user, requestedUser *models.User) (ApiResponse, error) {
	var userResult UserResult
	var Games []models.Game
	var FavouriteGames []GameApiResponse

	userResult.ID = requestedUser.ID
	userResult.Nickname = requestedUser.Nickname
	userResult.Username = requestedUser.Username
	userResult.IsReady = requestedUser.IsReady
	userResult.IsOnline = requestedUser.IsOnline
	Games, err := models.GetFavouriteGamesByUserID(requestedUser.ID)
	if err != nil {
		return ApiResponse{}, err
	}
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
		userResult.Photo = ConvertPhotoUrl(requestedUser.Photo)
	}

	if user.ID == requestedUser.ID {
		userResult.IsMine = true
		userResult.IsFriend = false
	} else {
		userResult.IsMine = false
		for _, friend := range requestedUser.Friends {
			if friend.ID == user.ID {
				userResult.IsFriend = true
				break
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
	Rooms, err := models.GetUserRoomIDs(requestedUser.ID)
	if err != nil {
		return ApiResponse{}, err
	}
	if len(Rooms) != 0 {
		userResult.RoomID = Rooms[0]
	}
	userResult.LikeCount = len(requestedUser.Likes)
	userResult.DislikeCount = len(requestedUser.Dislikes)
	// userResult.LikedGames = []string{} // Populate as needed

	return ApiResponse{
		Message:    "The user.",
		IsSuccess:  true,
		StatusCode: 200,
		Result:     userResult,
	} , nil
}

//---------------------------------------------

type SquadApiResponse struct {
	ID           string                   `json:"id"`
	RoomName     string                   `json:"room_name"`
	Owner        string                   `json:"owner"`
	Game         string                   `json:"game"`
	MemberLimit  int                      `json:"member_limit"`
	Member       []AccountApiResponse     `json:"member"`
	IsOwner      bool                     `json:"is_owner"`
	Created      time.Time                `json:"created"`
	UserID       string                   `json:"user_id"`
	JoinRequests []JoinRequestApiResponse `json:"join_requests"`
}

type AccountApiResponse struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Photo    string `json:"photo"`
	IsReady  bool   `json:"is_ready"`
}

type JoinRequestApiResponse struct {
	ID         string
	FromUser   AccountApiResponse
}

func CreateSquadResponse(userID string) (ApiResponse, error) {
	var room models.Room
	var SquadResult SquadApiResponse
	var requests []JoinRequestApiResponse

	room, err := models.FindRoomByUser(userID)
	if err != nil {
		return ApiResponse{}, err
	}

	// fmt.Println(room)
	SquadResult.ID = room.ID
	SquadResult.RoomName = room.RoomName
	SquadResult.Owner = room.Owner.Username
	SquadResult.Game = room.Game.GameIDD
	SquadResult.MemberLimit = room.MemberLimit
	SquadResult.Created = room.Created
	fmt.Println(room.Game)
	for _, member := range room.Member {
		var Account AccountApiResponse
		Account.ID = member.ID
		Account.Nickname = member.Nickname
		Account.Username = member.Username
		Account.Photo = ConvertPhotoUrl(member.Photo)
		Account.IsReady = member.IsReady
		SquadResult.Member = append(SquadResult.Member, Account)
	}
	if userID == room.Owner.ID {
		SquadResult.IsOwner = true
		joinRequests, err := models.GetJoinRequestByRoomID(room.ID)
		if err != nil {
			return ApiResponse{}, err
		}
		// fmt.Println(joinRequests)
		for _, req := range joinRequests {
			var Account AccountApiResponse
			var JoinRequest JoinRequestApiResponse
			JoinRequest.ID = req.ID
			user := req.FromUser
			Account.ID = user.ID
			Account.Nickname = user.Nickname
			Account.Username = user.Username
			Account.Photo = ConvertPhotoUrl(user.Photo)
			Account.IsReady = user.IsReady
			JoinRequest.FromUser = Account
			requests = append(requests, JoinRequest)
		}
		SquadResult.JoinRequests = requests
	} else {
		SquadResult.IsOwner = false
	}
	SquadResult.UserID = userID
	return ApiResponse{
		Message:    "Squad",
		IsSuccess:  true,
		StatusCode: 200,
		Result:     SquadResult,
	} , nil
	

}


func ConvertPhotoUrl(Url string) string {
	if Url == "" {
		return ""
	}
	NewUrl := "https://tz.radar.game/media/" + Url
	return NewUrl
}