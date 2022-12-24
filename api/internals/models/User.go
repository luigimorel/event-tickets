package models

import (
	"errors"
	"html"
	"strings"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName   string `gorm:"size:255;not null;" json:"firstName"`
	LastName    string `gorm:"size:255;not null;" json:"lastName"`
	PhoneNumber string `gorm:"type:text;not null;unique" json:"phoneNumber"`
	Email       string `gorm:"size:100;not null;unique" json:"email"`
	Password    string `gorm:"size:100;not null;" json:"password"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
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

func (user *User) Prepare() {
	user.ID = 0
	user.FirstName = html.EscapeString(strings.TrimSpace(user.FirstName))
	user.LastName = html.EscapeString(strings.TrimSpace(user.LastName))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.PhoneNumber = html.EscapeString(strings.TrimSpace(user.PhoneNumber))

}

func (user *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if user.FirstName == "" {
			return errors.New("first name is required")
		}
		if user.LastName == "" {
			return errors.New("last name is required")
		}
		if user.PhoneNumber == "" {
			return errors.New("phone number is required")
		}
		if user.Password == "" {
			return errors.New("password is required")
		}
		if user.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("email is invalid")
		}

		return nil

	case "login":
		if user.Password == "" {
			return errors.New("password is required")
		}
		if user.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("email is invalid")
		}
		return nil

	default:
		if user.FirstName == "" {
			return errors.New("first name is required")
		}
		if user.LastName == "" {
			return errors.New("last name is required")
		}
		if user.PhoneNumber == "" {
			return errors.New("phone number is required")
		}
		if user.Password == "" {
			return errors.New("password is required")
		}
		if user.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("email is invalid")
		}

		return nil
	}
}
