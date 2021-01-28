package controllers

import (
	"api/auth"
	"api/models"
	"api/responses"
	"api/security"
	"api/utils/email"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Login fuction is use to login for user  .
func Login(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := auth.SignIn(user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

// Registration function is use for registration of new user
func Registration(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	user := models.User{}
	err := json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := auth.SignUp(user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

//  AdminLogin function is only use for admin login
func AdminLogin(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body) // it reads the data for the api asign to the body
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user) // here the data in body is store at the user's memory address
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := auth.AdminSignIn(user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var fPass models.Password
	err = json.Unmarshal(body, &fPass)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := auth.EmailPassword(fPass.Email)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// VErify Email in db
	var otp string
	if user == true {
		fmt.Println("user exist")
		otp, err = security.GenerateCode("sakldgofsagofiusahf", time.Now())
		// otp = security.GenerateOTP()
		email.SendEmail2(fPass.Email, otp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(otp)
	}
	if user == false {
		fmt.Println("user dose not exist !!")
		responses.JSON(w, http.StatusBadRequest, nil)
	}
	// Generate OTP and send email'

	responses.JSON(w, http.StatusOK, otp)
}

func NewPassword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var fPass models.Password
	err = json.Unmarshal(body, &fPass)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	password := fPass.Password
	email := fPass.Email
	otp := fPass.Otp
	// fmt.Println(password)
	// fmt.Println(email)
	// fmt.Println(otp)

	// 1. Check email
	users := auth.EmailPassword(fPass.Email)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	conform := security.Validate(otp, "sakldgofsagofiusahf")
	fmt.Println(users)
	fmt.Println(conform)
	if users == true && conform == true {
		user := auth.SetPassword(email, password)
		fmt.Println(user)
		fmt.Println("Password update")
	}

	// user := models.User{}

	responses.JSON(w, http.StatusOK, conform)
}
