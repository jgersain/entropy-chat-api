package models

import "time"

//An Conversation represent the conversations table in the database
type Conversation struct {
	ID         uint64 `gorm:"primary_key;auto_increment" json:"id"`
	SenderId   uint32 `gorm:"not null" json:"sender_id"`
	ReceiverId uint32 `gorm:"not null" json:"receiver_id"`
	Messages   []Message
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
