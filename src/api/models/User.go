package models

import (
	"api/security"
	"time"
)

type User struct {
	ID        int       `gorm:"prinmary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:40;not null;" json:"name"`
	Email     string    `gorm:"size:60;not null;unique" json:"email"`
	Password  string    `gorm:"size:30;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
