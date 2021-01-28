package auto

import (
	"api/models"
)

var users = []models.User{
	models.User{Name: "sushil", Email: "sushil12@gmail.com", ProfilePic: "ha.jpg", Password: "123456789", UserType: "admin", Status: "Activated"},
}

var blogs = []models.Blog{
	models.Blog{
		Title:   "frist title",
		Content: "frist content",
	},
}
