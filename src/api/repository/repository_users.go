package repository

import (
	"api/models"
)

type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	Admin() (models.User, error)
	User(uint32) (models.User, error)
	FindUserById(uint32) (models.User, error)
	FindUserByEmail(string) (models.User, error)
	Update(uint32, models.User) (string, error)
	UploadPic(uint32, models.User) (string, error)
	ResetPassword(uint32, models.User) (string, error)
	UpdateByAdmin(uint32, models.User) (string, error)
	// Delete(uint32) (int64, error)
	DeleteByAdmin(uint32, models.User) (string, error)
}
