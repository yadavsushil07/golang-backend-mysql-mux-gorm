package crud

import (
	"api/models"
	"api/security"
	"api/utils/channels"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type repositoryUsersCRUD struct {
	db *gorm.DB
}

// NewRepositoryUsersCURD function is use to  update data in the database
func NewRepositoryUsersCURD(db *gorm.DB) *repositoryUsersCRUD {

	return &repositoryUsersCRUD{db}
}

func (r *repositoryUsersCRUD) Save(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
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
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Where("user_type = ?", "user").Limit(10).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return users, nil
	}
	return nil, err
}

func (r *repositoryUsersCRUD) User(id uint32) (models.User, error) {
	var err error
	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Where("id = ?", id).Limit(1).Find(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	// if gorm.ErrRecordNotFound.Error(err) {
	// 	return models.User{}, errors.New("Record not found")
	// }
	return models.User{}, err
}

func (r *repositoryUsersCRUD) Admin() (models.User, error) {
	var err error
	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Where("user_type = ?", "admin").Limit(1).Find(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	// if gorm.ErrRecordNotFound.Error(err) {
	// 	return models.User{}, errors.New("Record not found")
	// }
	return models.User{}, err
}

func (r *repositoryUsersCRUD) FindUserByEmail(email string) (models.User, error) {

	// this fuction is use to find user by email and it use while password reset

	var err error
	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Where("email = ?", email).Limit(1).Find(&user).Error
		if err != nil {
			fmt.Println("hello1")
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	return user, errors.New("Record not found")
}

// FindUserById function is use to find user by id and this function is only call by admin
func (r *repositoryUsersCRUD) FindUserById(id uint32) (models.User, error) {

	var err error
	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Where("id = ?", id).Limit(1).Find(&user).Error

		if err != nil {
			ch <- false
			return
		}
		if user.ID == 0 {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {

		return user, nil
	}
	return user, errors.New("Record not found")
}

func (r *repositoryUsersCRUD) UpdateByAdmin(id uint32, user models.User) (string, error) {

	// This function is use to upadte admin name

	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", id).UpdateColumn("name", user.Name)
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return "", rs.Error
		}
		return "successfully updated", nil
	}
	return "", rs.Error
}

func (r *repositoryUsersCRUD) ResetPassword(id uint32, user models.User) (string, error) {

	// This function is use to reset Password for user when he forgot password

	var rs *gorm.DB
	hashedPassword, _ := security.Hash(user.Password)
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", id).UpdateColumn("password", hashedPassword)
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return "", rs.Error
		}
		return "successfully change the password", nil
	}
	return "", rs.Error
}

func (r *repositoryUsersCRUD) Update(id uint32, user models.User) (string, error) {

	// This function is use to upadte user name

	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", id).UpdateColumn("name", user.Name)
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return "", rs.Error
		}
		return " successfully changed the name", nil
	}
	return "", rs.Error
}

func (r *repositoryUsersCRUD) UploadPic(id uint32, user models.User) (string, error) {

	// This function is use to upadte user profile pic

	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", id).UpdateColumn("profile_pic", user.ProfilePic)
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return "", rs.Error
		}
		return "successfully changed the profile-pic", nil
	}
	return "", rs.Error
}

func (r *repositoryUsersCRUD) DeleteByAdmin(id uint32, user models.User) (string, error) {

	// this function is use to deactivate user and it is only perform by admin

	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", id).UpdateColumn("status", user.Status)
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return "", rs.Error
		}
		return "successful", nil
	}
	return "", rs.Error
}

// func (r *repositoryUsersCRUD) Delete(id uint32) (int64, error) {
// 	var rs *gorm.DB
// 	done := make(chan bool)
// 	go func(ch chan<- bool) {
// 		defer close(ch)
// 		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", id).Take(&models.User{}).Delete(&models.User{})
// 		ch <- true
// 	}(done)
// 	if channels.OK(done) {
// 		if rs.Error != nil {
// 			return 0, rs.Error
// 		}
// 		return rs.RowsAffected, nil
// 	}
// 	return 0, rs.Error
// }
