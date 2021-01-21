package models

import (
	"api/security"
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"gorm.io/gorm"
)

type User struct {
	ID         int       `gorm:"prinmary_key;auto_increment" json:"id"`
	Name       string    `gorm:"size:40;not null;" json:"name"`
	Email      string    `gorm:"size:60;not null;unique" json:"email"`
	Password   string    `gorm:"size:30;not null" json:"password"`
	ProfilePic string    `gorm:"column:profile_pic" json:"profile_pic"`
	UserType   string    `gorm:"column:user_type;default:'user'" json:"user_type"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hashedPassword, err := security.Hash(u.Password)
	fmt.Println(hashedPassword)
	fmt.Println(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}

// func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
// 	hashedPassword, err := security.Hash(u.Password)
// 	fmt.Println(hashedPassword)
// 	fmt.Println(u.Password)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = hashedPassword
// 	return nil
// }

func (u *User) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *User) Validate(action string) error {
	switch action {
	case "update":
		if u.Name == "" {
			return errors.New("required name")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
		if u.Status == "" {
			return errors.New("required status")
		}
		return nil
	case "login":
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}

		if u.Password == "" {
			return errors.New("required password")
		}

		return nil
	default:
		if u.Name == "" {
			return errors.New("required name")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
		if u.Status == "" {
			return errors.New("required status")
		}
		return nil
	}
	return nil
}
