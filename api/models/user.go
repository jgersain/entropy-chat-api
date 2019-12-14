package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"time"
)

//An User represent the users table in the database
type User struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name         string    `gorm:"not null;" json:"name"`
	Age          int       `json:"age,omitempty,string"`
	Email        string    `gorm:"size:100;not null;unique" json:"email"`
	ProfilePhoto *string   `gorm:"type:varchar(100);" json:"profile_photo,omitempty"`
	Password     string    `gorm:"size:100;not null;" json:"password,omitempty"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

//returns the bcrypt hash of the user password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//validate some user fields
func (user *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if user.Password == "" {
			return errors.New("La contraseña es necesaria")
		}
		if user.Email == "" {
			return errors.New("El email es necesario")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Email invalido")
		}
		return nil

	default:
		if user.Name == "" {
			return errors.New("El nombre es necesario")
		}
		if user.Password == "" {
			return errors.New("La contraseña es necesaria")
		}
		if user.Email == "" {
			return errors.New("El correo electrónico es necesario")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("El correo electrónico no es válido")
		}
		return nil
	}
}

//save user instance on database
func (user *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *User) BeforeSave() error {
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := user.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":      user.Password,
			"name":          user.Name,
			"age":           user.Age,
			"email":         user.Email,
			"profile_photo": user.ProfilePhoto,
			"update_at":     time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}

	// updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
