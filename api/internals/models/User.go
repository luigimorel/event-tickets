package models

import (
	"encoding/json"
	"errors"
	"html"
	"io"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"size:255;not null;" validate:"required" json:"name"`
	PhoneNumber string `gorm:"size:20;type:text;not null;unique" validate:"required" json:"phone"`
	Role        string `gorm:"size:20;default:user" json:"role"`
	Email       string `gorm:"size:100;not null;unique" validate:"required" json:"email"`
	Password    string `gorm:"size:100;not null;" validate:"required" json:"password"`
	Verified    bool   `gorm:"default:false;" json:"verified"`
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uint32    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) FromJSON(r io.Reader) error {
	err := json.NewDecoder(r)
	return err.Decode(u)
}

func (u *User) ValidateJSON() error {
	validate := validator.New()
	return validate.Struct(u)
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *User) BeforeSave(*gorm.DB) error {
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) Prepare() {
	user.ID = 0
	user.Name = html.EscapeString(strings.TrimSpace(user.Name))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.PhoneNumber = html.EscapeString(strings.TrimSpace(user.PhoneNumber))

}

func (user *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if user.Name == "" {
			return errors.New("name is required")
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
		if user.Name == "" {
			return errors.New("first name is required")
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
