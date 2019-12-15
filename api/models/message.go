package models

import "time"

type Message struct {
	ID             uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Body           string `sql:"type:text;" json:"body"`
	ConversationID uint
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
