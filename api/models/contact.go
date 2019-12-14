package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

//An Contact represent the contacts table in the database
type Contact struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"not nul" json:"name"`
	Nickname  string    `json:"nickname"`
	Email     string    `gorm:"size:100;not null;" json:"email"`
	Phone     string    `gorm:"size:20" json:"phone"`
	Address   string    `gorm:"size:255" json:"address"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

//validations
func (contact *Contact) Validate() error {

	if contact.Name == "" {
		return errors.New("El nombre es requerido")
	}
	if contact.Email == "" {
		return errors.New("El correo electrónico es requerido")
	}
	if contact.UserID < 1 {
		return errors.New("El usuario es requerido")
	}
	return nil
}

//store contact instance
func (contact *Contact) SaveContact(db *gorm.DB) (*Contact, error) {
	var err error
	err = db.Debug().Model(&Contact{}).Create(&contact).Error
	if err != nil {
		return &Contact{}, err
	}
	if contact.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", contact.UserID).Error
		if err != nil {
			return &Contact{}, err
		}
	}
	return contact, nil
}
