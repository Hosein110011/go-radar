package models

import (
	// "gorm.io/driver/postgres"
	"time"
  	"gorm.io/gorm"
	"github.com/Hosein110011/go-radar/pkg/config"
	// pq "github.com/lib/pq"
)

var db *gorm.DB

type Tabler interface {
	TableName() string
  }
  

type User struct {
    ID        string `gorm:"type:uuid;primary_key;"`
    Username  string `gorm:"size:120;uniqueIndex"`
    IsOnline  bool   `gorm:"default:false"`
    Nickname  string `gorm:"size:122"`
    // Photo field is a URL or path to the image, stored as text
    Photo     string `gorm:"column:photo;"`
    IsStaff   bool `gorm:"default:false"`
    IsActive  bool `gorm:"default:true"`
    Created   time.Time `gorm:"autoCreateTime"`
    Bio       string `gorm:"column:bio;type:text"`
    Playing   string `gorm:"column:playing;size:255"`
    LastSeen  time.Time
    IsDeleted bool `gorm:"default:false"`
    IsReady   bool `gorm:"default:false"`
    // For ManyToMany relationships, use a slice of pointers to the related struct
    Friends   []*User `gorm:"many2many:account_user_friends;joinForeignKey:id;joinReferences:to_user_id"`
    Likes     []*User `gorm:"many2many:account_user_likes;"`
    Dislikes  []*User `gorm:"many2many:account_user_dislikes;"`
    Rooms     []Room  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:Owner"`
}


type Room struct {
    ID        string `gorm:"type:uuid;primary_key"`
    OwnerID     string 
    Owner       User `gorm:"foreignKey:OwnerID"`
    GameID      string 
    RoomName    string `gorm:"size:122"`
    IsPrivate   bool   `gorm:"default:false"`
    MemberLimit int    `gorm:"default:2"`
    Created     time.Time
    IsDeleted   bool `gorm:"default:false"`
    // Relationships
    Game        Game    `gorm:"foreignKey:GameID"`
    Member      []*User `gorm:"many2many:chat_room_member;"`
}


type Game struct {
    ID            string `gorm:"column:game_id;type:uuid;primary_key"`
    GameName      string `gorm:"size:122;uniqueIndex;"`
    PackageName   string `gorm:"column:package_name;size:222;"`
    Image         string `gorm:"column:image;size:222;"`
    Banner        string `gorm:"column:banner;size:222;"`
    Platform      string `gorm:"column:platform;size:122;"`
    IsDeleted     bool   `gorm:"column:isdeleted;default:false"`
    Rooms         []Room  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:Game"`

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

func GetUserByUsername(username string) User {
    var user User
    db.Where("username = ?", username).Preload("Friends").Preload("Likes").Preload("Dislikes").Preload("Rooms").First(&user)
    return user
}