package models

import (
	// "gorm.io/driver/postgres"
	"time"
	"gorm.io/gorm"
)

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

func (Room) TableName() string {
	return "chat_room"
}

func GetAllRooms() []Room {
	var Rooms []Room
	db.Preload("Owner").Preload("Member").Find(&Rooms)
	return Rooms
}

func FindRoomByUser(userID string) (Room, error) {
	var room Room

	result := db.Debug().Preload("Game", func(db *gorm.DB) *gorm.DB {
		return db.Select("GameIDD", "ID")
		}).Preload("Owner", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Username")
		}).Preload("Member", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Username", "Nickname", "Photo", "IsReady")
		}).Joins("JOIN chat_room_member on chat_room_member.room_id = chat_room.id").
		Where("chat_room_member.user_id = ?", userID).
		Where("is_deleted = ?", false).
		First(&room)

	if result.Error != nil {
		return Room{}, result.Error
	}

	if err := db.Model(room).Association("Game").Find(&room.Game); err != nil {
        return Room{}, err
    }
	return room, nil
}
