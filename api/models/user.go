package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name      		string         	`gorm:"not null;" json:"name,omitempty"`
	Age       		sql.NullInt64  	`json:"age,omitempty"`
	Email     		string    	 	`gorm:"size:100;not null;unique" json:"email"`
	ProfilePhoto  	*string    		`gorm:"type:varchar(100);" json:"profile_photo,omitempty"`
	Password  		string    		`gorm:"size:100;not null;" json:"password,omitempty"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
