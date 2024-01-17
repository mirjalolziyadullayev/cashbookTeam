package handler

import (
	"cashbookTeam/helper"
	"cashbookTeam/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetUser(w, r)
	case "POST":
		CreateUser(w, r)
	case "PUT":
		UpdateUser(w, r)
	case "DELETE":
		DeleteUser(w, r)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.UserModel
	json.NewDecoder(r.Body).Decode(&newUser)

	var userData []models.UserModel
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &userData)

	newUser.Id = helper.MaxIdUser(userData)

	if newUser.Firstname == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "User's Firstname cannot be an empty string!")
		return
	} else if newUser.Lastname == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "User's Lastname cannot be an empty string!")
		return
	}

	userData = append(userData, newUser)

	res, _ := json.Marshal(userData)
	os.WriteFile("db/all.json", res, 0)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, "User created with ID: %d\n", newUser.Id)

	json.NewEncoder(w).Encode(newUser)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser models.UserModel
	json.NewDecoder(r.Body).Decode(&updateUser)

	var userData []models.UserModel
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &userData)

	var userFound bool
	for i := 0; i < len(userData); i++ {
		if userData[i].Id == updateUser.Id {
			if updateUser.Firstname == "" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, "User's Firstname cannot be an empty string!")
				return
			} else if updateUser.Lastname == "" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, "User's Lastname cannot be an empty string!")
				return
			}
			userData[i] = updateUser
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User not found with ID: %d\n", updateUser.Id)
		return
	}

	res, _ := json.Marshal(userData)
	os.WriteFile("db/all.json", res, 0)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, "User updated with ID: %d\n", updateUser.Id)

	json.NewEncoder(w).Encode(updateUser)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var deleteUser models.UserModel
	json.NewDecoder(r.Body).Decode(&deleteUser)

	var userData []models.UserModel
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &userData)

	var userFound bool
	for i := 0; i < len(userData); i++ {
		if userData[i].Id == deleteUser.Id {
			userData = append(userData[:i], userData[i+1:]...)
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User not found with ID: %d\n", deleteUser.Id)
		return
	}

	res, _ := json.Marshal(userData)
	os.WriteFile("db/all.json", res, 0)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, "User deleted with ID: %d\n", deleteUser.Id)

	json.NewEncoder(w).Encode(deleteUser)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	var getUser models.UserModel
	json.NewDecoder(r.Body).Decode(&getUser)

	var userData []models.UserModel
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &userData)

	var userFound bool
	var TakeUser models.UserModel

	for i := 0; i < len(userData); i++ {
		if userData[i].Id == getUser.Id {
			TakeUser = userData[i]
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User not found with ID: %d\n", getUser.Id)
		return
	}

	res, _ := json.Marshal(userData)
	os.WriteFile("db/all.json", res, 0)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, "User found with ID: %d\n", getUser.Id)

	json.NewEncoder(w).Encode(TakeUser)
}
