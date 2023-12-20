package models

import "time"

// "gorm.io/driver/postgres"

// pq "github.com/lib/pq"

type Game struct {
	ID          string `json:"id" gorm:"column:id;primary_key"`   
	GameIDD      string `json:"game_id" gorm:"column:game_id;"`
	GameName    string `json:"game_name" gorm:"size:122;uniqueIndex;"`
	PackageName string `json:"package_name" gorm:"column:package_name;size:222;"`
	Image       string `json:"image" gorm:"column:image;size:222;"`
	Banner      string `json:"banner" gorm:"column:banner;size:222;"`
	Platform    string `json:"platform" gorm:"column:platform;size:122;"`
	IsDeleted   bool   `json:"is_deleted" gorm:"column:isdeleted;default:false"`
	Rooms       []Room `json:"rooms" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:Game"`
}

type FavouriteGame struct {
	ID        string    `json:"id"`
	GameID    string    `json:"game_id" gorm:"column:game_id;type:uuid;primary_key"`
	Game      Game      `json:"game" gorm:"foreignKey:GameID;references:ID"`
	UserID    string    `json:"user_id" gorm:"column:user_id;type:uuid;primary_key"`
	LikedAt   time.Time `json:"liked_at" gorm:"column:liked_at;"`
	IsDeleted bool      `json:"is_deleted" gorm:"column:isdeleted;default:false"`
}

func (FavouriteGame) TableName() string {
	return "account_favouritegame"
}

func (Game) TableName() string {
	return "game_game"
}

func GetAllGames() []Game {
	var Games []Game
	db.Find(&Games)
	return Games
}

