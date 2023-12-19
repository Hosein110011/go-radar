package models

import (
	// "gorm.io/driver/postgres"
	"time"
	// pq "github.com/lib/pq"
)

type User struct {
	ID       string `gorm:"type:uuid;primary_key;"`
	Username string `gorm:"size:120;uniqueIndex"`
	IsOnline bool   `gorm:"default:false"`
	Nickname string `gorm:"size:122"`
	// Photo field is a URL or path to the image, stored as text
	Photo     string    `gorm:"column:photo;"`
	IsStaff   bool      `gorm:"default:false"`
	IsActive  bool      `gorm:"default:true"`
	Created   time.Time `gorm:"autoCreateTime"`
	Bio       string    `gorm:"column:bio;type:text"`
	Playing   string    `gorm:"column:playing;size:255"`
	LastSeen  time.Time
	IsDeleted bool `gorm:"default:false"`
	IsReady   bool `gorm:"default:false"`
	// For ManyToMany relationships, use a slice of pointers to the related struct
	Friends        []*User         `gorm:"many2many:account_user_friends;joinForeignKey:from_user_id;joinReferences:to_user_id"`
	Likes          []*User         `gorm:"many2many:account_user_likes;joinForeignKey:from_user_id;joinReferences:to_user_id"`
	Dislikes       []*User         `gorm:"many2many:account_user_dislikes;joinForeignKey:from_user_id;joinReferences:to_user_id"`
	// Rooms          []*Room         `gorm:"many2many:chat_room_member;"`
	FavouriteGames []FavouriteGame `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:user_id"`
}

func (User) TableName() string {
	return "account_user"
}

func GetAllUsers() []User {
	var Users []User
	db.Preload("Friends").Preload("Likes").Preload("Dislikes").Preload("Rooms").Find(&Users)
	return Users
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	// db.Where("username = ?", username).Preload("Friends").Preload("Likes").Preload("Dislikes").Preload("Rooms").First(&user)
	result := db.Preload("Friends").Preload("Likes").Preload("Dislikes").First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetFavouriteGamesByUserID(userID string) ([]Game, error) {
	var favouriteGames []FavouriteGame

	result := db.Debug().Preload("Game").Where("user_id = ? AND is_deleted = ?", userID, false).Find(&favouriteGames)
	if result.Error != nil {
		return nil, result.Error
	}

	var games []Game
	for _, favGame := range favouriteGames {
		games = append(games, favGame.Game)
	}

	return games, nil
}

func GetUserByID(userID string) (*User, error) {
	var user User
	result := db.Preload("Friends").Preload("Likes").Preload("Dislikes").Preload("FavouriteGames").First(&user, "id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}


func GetUserRoomIDs(userID string) ([]string, error) {
    var roomIDs []string

    // Assuming 'db' is your *gorm.DB instance
    result := db.Table("chat_room_member").Select("room_id").
              Where("user_id = ?", userID).
              Scan(&roomIDs)

    if result.Error != nil {
        return nil, result.Error
    }

    return roomIDs, nil
}