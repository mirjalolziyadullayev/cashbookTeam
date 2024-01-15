package handlers

import (
	"cashbookTeam/models"
	"encoding/json"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		CreateUser(w, r)
	case "POST":
	case "PUT":
	case "DELETE":
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.UserModel
	json.NewDecoder(r.Body).Decode(&newUser)
	
	
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	
}