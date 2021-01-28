package auth

import (
	"api/database"
	"api/models"
	"api/security"
	"api/utils/channels"
	"fmt"

	"gorm.io/gorm"
)

// this fuction is use to signIn only if user is active
// and it return token as an output
func SignIn(email, password string) (string, error) {

	user := models.User{}
	var err error
	var db *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		db, err = database.Connect()
		if err != nil {
			ch <- false
			return
		}
		err = db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		if user.Status == "Deactivated" {
			ch <- false
			return
		}
		pass := security.VerifyPassword(user.Password, password)
		if pass == false {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return CreateToken(uint32(user.ID), user.UserType)
	}
	fmt.Println(err)
	return "", err
}

//this function is use to create new user when user is registering himself
// and it return the token of the user
func SignUp(email, password string) (string, error) {

	user := models.User{}
	user.Email = email
	user.Password = password
	user.Status = "Activated"
	var err error
	var db *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		db, err = database.Connect()
		if err != nil {
			ch <- false
			return
		}

		err = db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return CreateToken(uint32(user.ID), user.UserType)
	}
	fmt.Println(err)
	return "", err
}

//this is signIn Api for admin
func AdminSignIn(email, password string) (string, error) {

	user := models.User{}
	var err error
	var db *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		db, err = database.Connect()
		if err != nil {
			ch <- false
			return
		}
		// defer db.Close()

		err = db.Debug().Model(models.User{}).Where("email = ? AND user_type = ? ", email, "admin").Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		pass := security.VerifyPassword(user.Password, password)
		if pass == false {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return CreateToken(uint32(user.ID), user.UserType)
	}
	fmt.Println(err)
	return "", err
}

func EmailPassword(email string) bool {

	// This fuction checks the user email at the time of forgot password api that user exist or not

	user := models.User{}
	var err error
	var db *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		db, err = database.Connect()
		if err != nil {
			ch <- false
			return
		}
		// defer db.Close()

		err = db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return true
	}
	return false
}

func SetPassword(email, password string) bool {

	//this function is use set new password after user is sent opt on its email

	var err error
	var db *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		db, err = database.Connect()
		if err != nil {
			ch <- false
			return
		}
		// defer db.Close()
		hashedPassword, _ := security.Hash(password)
		db = db.Debug().Model(models.User{}).Where("email = ?", email).UpdateColumns(
			map[string]interface{}{
				"password": hashedPassword,
			},
		)
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return true
	}
	return false
}
