package repository

import (
	"api/models"
)

type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	Admin() (models.User, error)
	FindUserById(uint32) (models.User, error)
	Update(uint32, models.User) (int64, error)
	UploadPic(uint32, models.User) (int64, error)
	ResetPassword(uint32, models.User) (int64, error)
	UpdateByAdmin(uint32, models.User) (int64, error)
	Delete(uint32) (int64, error)
	DeleteByAdmin(uint32, models.User) (int64, error)
}
