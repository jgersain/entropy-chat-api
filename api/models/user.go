package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

//An User represent the users table in the database
type User struct {
	gorm.Model
	Name         string  `gorm:"not null;" json:"name"`
	Age          int     `json:"age,omitempty,string"`
	Email        string  `gorm:"size:100;not null;unique" json:"email"`
	ProfilePhoto *string `gorm:"type:varchar(100);" json:"profile_photo,omitempty"`
	Password     string  `gorm:"size:100;not null;" json:"password"`
}

//returns the bcrypt hash of the user password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//validate some user fields
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if u.Password == "" {
			return errors.New("La contraseña es necesaria")
		}
		if u.Email == "" {
			return errors.New("El email es necesario")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email invalido")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("El nombre es necesario")
		}
		if u.Password == "" {
			return errors.New("La contraseña es necesaria")
		}
		if u.Email == "" {
			return errors.New("El correo electrónico es necesario")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("El correo electrónico no es válido")
		}
		return nil
	}
}

//save user instance on database
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":      u.Password,
			"name":          u.Name,
			"age":           u.Age,
			"email":         u.Email,
			"profile_photo": u.ProfilePhoto,
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
