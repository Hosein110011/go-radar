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
    ID        string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
    Username  string `gorm:"size:120;uniqueIndex"`
    IsOnline  bool
    Nickname  string `gorm:"size:122"`
    // Photo field is a URL or path to the image, stored as text
    Photo     string
    IsStaff   bool `gorm:"default:false"`
    IsActive  bool `gorm:"default:true"`
    Created   time.Time `gorm:"autoCreateTime"`
    Bio       string `gorm:"type:text"`
    Playing   string `gorm:"size:255"`
    LastSeen  time.Time
    IsDeleted bool `gorm:"default:false"`
    IsReady   bool `gorm:"default:false"`
    // For ManyToMany relationships, use a slice of pointers to the related struct
    Friends   []*User `gorm:"many2many:user_friends;"`
    Likes     []*User `gorm:"many2many:user_likes;"`
    Dislikes  []*User `gorm:"many2many:user_dislikes;"`

}

func (User) TableName() string {
	return "account_user"
}	

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}
