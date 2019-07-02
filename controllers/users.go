package controllers

import (
	"math/rand"
	models "usersApi/models"
	u "usersApi/utils"

	"net/http"

	"github.com/Pallinder/go-randomdata"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	data := generateUsers()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func generateUsers() []models.User {
	users := make([]models.User, 0)
	for i := 0; i < 100000; i++ {
		users = append(users, models.User{Name: randomdata.SillyName(), Age: rand.Intn(99)})
	}
	return users
}
