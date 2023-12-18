package models

import (
	// "gorm.io/driver/postgres"
	"time"

	"github.com/Hosein110011/go-radar/pkg/config"
	"gorm.io/gorm"
	// pq "github.com/lib/pq"
)

var db *gorm.DB

type Tabler interface {
	TableName() string
}

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
	Rooms          []Room          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:owner_id"`
	FavouriteGames []FavouriteGame `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:user_id"`
}

type Room struct {
	ID          string `gorm:"type:uuid;primary_key"`
	OwnerID     string
	Owner       User `gorm:"foreignKey:OwnerID"`
	GameID      string
	RoomName    string `gorm:"size:122"`
	IsPrivate   bool   `gorm:"default:false"`
	MemberLimit int    `gorm:"default:2"`
	Created     time.Time
	IsDeleted   bool `gorm:"default:false"`
	// Relationships
	Game   Game    `gorm:"foreignKey:GameID"`
	Member []*User `gorm:"many2many:chat_room_member;"`
}

type Game struct {
	ID          string `gorm:"column:id;type:uuid;primary_key"`
	GameName    string `gorm:"size:122;uniqueIndex;"`
	PackageName string `gorm:"column:package_name;size:222;"`
	Image       string `gorm:"column:image;size:222;"`
	Banner      string `gorm:"column:banner;size:222;"`
	Platform    string `gorm:"column:platform;size:122;"`
	IsDeleted   bool   `gorm:"column:isdeleted;default:false"`
	Rooms       []Room `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:Game"`
}

type FavouriteGame struct {
	ID        string
	GameID    string    `gorm:"column:game_id;type:uuid;primary_key"`
	Game      Game      `gorm:"foreignKey:GameID;references:ID"`
	UserID    string    `gorm:"column:user_id;type:uuid;primary_key"`
	LikedAt   time.Time `gorm:"column:liked_at;"`
	IsDeleted bool      `gorm:"column:isdeleted;default:false"`
}

func (User) TableName() string {
	return "account_user"
}

func (Room) TableName() string {
	return "chat_room"
}

func (Game) TableName() string {
	return "game_game"
}

func (FavouriteGame) TableName() string {
	return "account_favouritegame"
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func GetAllUsers() []User {
	var Users []User
	db.Preload("Friends").Preload("Likes").Preload("Dislikes").Preload("Rooms").Find(&Users)
	return Users
}

func GetAllRooms() []Room {
	var Rooms []Room
	db.Preload("Owner").Preload("Member").Find(&Rooms)
	return Rooms
}

func GetAllGames() []Game {
	var Games []Game
	db.Find(&Games)
	return Games
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	// db.Where("username = ?", username).Preload("Friends").Preload("Likes").Preload("Dislikes").Preload("Rooms").First(&user)
	result := db.Preload("Friends").Preload("Likes").Preload("Dislikes").Preload("Rooms").First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByID(userID string) (*User, error) {
	var user User
	result := db.Preload("Friends").Preload("Likes").Preload("Dislikes").Preload("Rooms").Preload("FavouriteGames").First(&user, "id = ?", userID)
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
