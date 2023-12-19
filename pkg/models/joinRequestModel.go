package models

import (
	"time"
	"fmt"
	"gorm.io/gorm"

)

type JoinRequest struct {
	ID         string `gorm:"column:id;type:uuid;primary_key"`
	FromUser   User `gorm:"foreignKey:FromUserID"`
	FromUserID string `gorm:"column:from_user_id;"`
	Status     string `gorm:"column:request_status;default:PENDING"`
	ToRoom     Room `gorm:"foreignKey:ToRoomID"`
	ToRoomID   string `gorm:"column:to_room_id;"`
	TimeStamp  time.Time `gorm:"column:timestamp"`
	IsDeleted  bool `gorm:"column:is_deleted;default:false"`
}

func (JoinRequest) TableName() string {
	return "request_joinrequest"
}

func GetAllJoinReqs() []JoinRequest {
	var JoinRequests []JoinRequest
	db.Preload("FromUser").Preload("ToRoom").Find(&JoinRequests)
	return JoinRequests
}

func GetJoinRequestByRoomID(RoomID string) ([]JoinRequest, error) {
	var joinRequests []JoinRequest

    // Assuming 'db' is your *gorm.DB instance
    result := db.Debug().Preload("FromUser", func(db *gorm.DB) *gorm.DB {
					return db.Select("ID", "Username", "Nickname", "Photo", "IsReady")
				}).
				Where("to_room_id = ? AND request_status = ? AND is_deleted = ?", RoomID, "PENDING", false).
				Find(&joinRequests)

    if result.Error != nil {
        fmt.Println("Error fetching join requests:", result.Error)
		return nil, result.Error
    }

    return joinRequests, nil
}
