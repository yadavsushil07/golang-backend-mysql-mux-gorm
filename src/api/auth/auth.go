package auth

import (
	"api/database"
	"api/models"
	"api/security"
	"api/utils/channels"
	"fmt"

	"gorm.io/gorm"
)

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
		// defer db.Close()

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

func SignUp(email, password string) (string, error) {
	user := models.User{}
	user.Email = email
	user.Password = password
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

		err = db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		// pass := security.VerifyPassword(user.Password, password)
		// if pass == false {
		// 	ch <- false
		// 	return
		// }
		ch <- true
	}(done)

	if channels.OK(done) {
		return CreateToken(uint32(user.ID), user.UserType)
	}
	fmt.Println(err)
	return "", err
}

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

		// fmt.Println("Random secret:", gotp.RandomSecret(16))
		// security.DefaultTOTPUsage()

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

// func ResetPassword(email string)(link string, error){
// 	user := models.User{}
// 	var err error
// 	var db *gorm.DB
// 	done := make(chan bool)

// 	fmt.Println("Hete 1")

// 	go func(ch chan<- bool) {
// 		defer close(ch)
// 		db, err = database.Connect()
// 		if err != nil {
// 			ch <- false
// 			return
// 		}
// 		// defer db.Close()

// 		err = db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
// 		if err != nil {
// 			ch <- false
// 			return
// 		}
// }
