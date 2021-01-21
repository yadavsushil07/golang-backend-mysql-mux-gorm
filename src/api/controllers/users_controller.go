package controllers

import (
	"api/auth"
	"api/database"
	"api/models"
	"api/repository"
	"api/repository/crud"
	"api/responses"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {
	//this function returns the filename(to save in database) of the saved file or an error if it occurs

	name, err := FileUpload(r)
	if err != nil {
		responses.ERROR(w, http.StatusNoContent, err)
		return
	}
	user := models.User{}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	userID, _, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	repo := crud.NewRepositoryUsersCURD(db)
	fmt.Println(user.ID)

	func(UserRepository repository.UserRepository) {
		user.ProfilePic = name
		user, err := UserRepository.UploadPic(uint32(userID), user)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI))
		responses.JSON(w, http.StatusCreated, user)
	}(repo)
	fmt.Println(name)
	fmt.Println(err)

}

func FileUpload(r *http.Request) (string, error) {

	r.ParseMultipartForm(32 << 10)

	file, handler, err := r.FormFile("file") //retrieve the file from form data

	fmt.Println(handler)
	if err != nil {
		return "", err
	}
	defer file.Close() //close the file when we finish

	f, err := os.OpenFile("upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()
	io.Copy(f, file)

	return handler.Filename, nil
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
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
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// defer db.Close()

	repo := crud.NewRepositoryUsersCURD(db)

	func(UserRepository repository.UserRepository) {
		user, err := UserRepository.Save(user)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
		responses.JSON(w, http.StatusCreated, user)
	}(repo)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {

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

	userID, _, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// defer db.Close()

	repo := crud.NewRepositoryUsersCURD(db)

	fmt.Println("hello ", user.ID)
	func(UserRepository repository.UserRepository) {
		rows, err := UserRepository.ResetPassword(userID, user)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, rows)
	}(repo)

}

// func ForgetPassword(w http.ResponseWriter, r *http.Request) {
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	opt := security.RandStringBytes(6)
// 	fmt.Println(opt)
// }

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

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

	userID, _, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// defer db.Close()

	repo := crud.NewRepositoryUsersCURD(db)
	if userID != uint32(id) {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if userID == uint32(id) {
		func(UserRepository repository.UserRepository) {
			rows, err := UserRepository.Update(uint32(id), user)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			responses.JSON(w, http.StatusOK, rows)
		}(repo)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userID, userType, err := auth.ExtractClaim(r)

	fmt.Println("USER :", userID)
	fmt.Println(userType)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	fmt.Println("USER :", userID)

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// defer db.Close()

	repo := crud.NewRepositoryUsersCURD(db)
	if userID != uint32(user.ID) {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	if userType == "user" {
		func(UserRepository repository.UserRepository) {
			_, err := UserRepository.Delete(uint32(id))
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}
			w.Header().Set("Entity", fmt.Sprintf("%d", id))
			responses.JSON(w, http.StatusNoContent, "")
		}(repo)
	}
}

//ADMIN PART

// func VerifyOpt(w http.ResponseWriter, r *http.Request){
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// }

func AdminGetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// defer db.Close()

	user_id, user_type, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	fmt.Println("USER :", user_id)
	fmt.Println("USER :", user_type)
	repo := crud.NewRepositoryUsersCURD(db)
	if user_type == "admin" {
		func(UserRepository repository.UserRepository) {
			users, err := UserRepository.FindAll()
			if err != nil {
				responses.ERROR(w, http.StatusUnprocessableEntity, err)
				return
			}

			responses.JSON(w, http.StatusOK, users)
		}(repo)
	}
	if user_type == "user" {
		responses.JSON(w, http.StatusUnauthorized, err)
		return
	}
}

func AdminUploadImage(w http.ResponseWriter, r *http.Request) {
	//this function returns the filename(to save in database) of the saved file or an error if it occurs

	name, err := FileUpload(r)
	user := models.User{}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	_, user_type, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	repo := crud.NewRepositoryUsersCURD(db)
	if user_type == "admin" {
		func(UserRepository repository.UserRepository) {
			user.ProfilePic = name
			user, err := UserRepository.Update(uint32(1), user)
			if err != nil {
				responses.ERROR(w, http.StatusUnprocessableEntity, err)
				return
			}
			w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI))
			responses.JSON(w, http.StatusCreated, user)
		}(repo)
		fmt.Println(name)
		fmt.Println(err)
	}
	if user_type == "user" {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
}

func AdminProfile(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// defer db.Close()
	_, user_type, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	repo := crud.NewRepositoryUsersCURD(db)
	if user_type == "admin" {
		func(UserRepository repository.UserRepository) {
			admin, err := UserRepository.Admin()
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			responses.JSON(w, http.StatusOK, admin)
		}(repo)
	}
	if user_type == "user" {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// defer db.Close()
	_, user_type, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	repo := crud.NewRepositoryUsersCURD(db)

	if user_type == "admin" {
		func(UserRepository repository.UserRepository) {
			user, err := UserRepository.FindUserById(uint32(id))
			fmt.Println("hello", err)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			responses.JSON(w, http.StatusOK, user)
		}(repo)
	}
	if user_type == "user" {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
}

func AdminUpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

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

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// defer db.Close()
	_, user_type, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	repo := crud.NewRepositoryUsersCURD(db)
	if user_type == "admin" {
		func(UserRepository repository.UserRepository) {
			rows, err := UserRepository.UpdateByAdmin(uint32(id), user)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			responses.JSON(w, http.StatusOK, rows)
		}(repo)
	}
	if user_type == "user" {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
}

func DeleteUserByAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
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

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// defer db.Close()
	_, user_type, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	repo := crud.NewRepositoryUsersCURD(db)
	if user_type == "admin" {
		func(UserRepository repository.UserRepository) {
			rows, err := UserRepository.DeleteByAdmin(uint32(id), user)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			responses.JSON(w, http.StatusOK, rows)
		}(repo)
	}
	if user_type == "user" {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
}

// func DeleteUserByAdmin(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.ParseUint(vars["id"], 10, 32)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	user := models.User{}
// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	// defer db.Close()
// 	_, user_type, err := auth.ExtractClaim(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// 	repo := crud.NewRepositoryUsersCURD(db)
// 	if user_type == "admin" {
// 		func(UserRepository repository.UserRepository) {
// 			_, err := UserRepository.DeleteByAdmin(uint32(id), user)
// 			if err != nil {
// 				responses.ERROR(w, http.StatusBadRequest, err)
// 				return
// 			}
// 			w.Header().Set("Entity", fmt.Sprintf("%d", id))
// 			responses.JSON(w, http.StatusNoContent, "")
// 		}(repo)
// 	}
// 	if user_type == "user" {
// 		responses.ERROR(w, http.StatusUnauthorized, err)
// 		return
// 	}
// }
