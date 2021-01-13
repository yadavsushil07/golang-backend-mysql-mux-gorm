package crud

import (
	"api/models"
	"api/utils/channels"
	"fmt"

	"gorm.io/gorm"
)

type repositoryUsersCRUD struct {
	db *gorm.DB
}

func NewRepositoryUsersCURD(db *gorm.DB) *repositoryUsersCRUD {
	return &repositoryUsersCRUD{db}
}

func (r *repositoryUsersCRUD) Save(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}

func (r *repositoryUsersCRUD) FindAll() ([]models.User, error) {
	var err error
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Limit(10).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		fmt.Println(users)
		return users, nil
	}
	return nil, err
}
