package handler 

import (
	"cashbookTeam/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Home(w, r)
	case "POST":
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "You cannot send request with method 'POST'!")
	case "PUT":
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "You cannot send request with method 'PUT'!")
	case "DELETE":
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "You cannot send request with method 'DELETE'!")
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	var UserData []models.UserModel
	byteData, _ := os.ReadFile("db/all.json")
	json.Unmarshal(byteData, &UserData)

	template := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>doc</title>
		<style>
		body {
			margin: 0;
			padding: 0;
			background-color: whitesmoke;
		}
		.container {
			text-align: center;
			margin: 0px;
			padding: 10px;
		}
		.list {
			text-align: left;
			max-width: 400px;
			border: 1px solid grey;
			border-radius: 4px;
			margin: 20px;
			padding: 10px;
			background-color: lightblue;
			font-family: sans-serif;
		}
		.list h1,h2,h3,h4,h5,h6 {
			margin-right: 5px;
			color: grey;
			margin: 10px;
			padding: 4px;
			border: 1px solid grey;
			border-radius: 4px;
		}
		.list i {
			color: black;
			font-family: calibri;
			font-size: 16px;
		}
		</style>
	</head>
	<body>
	<div class='container'>
	<h3>Users list</h3>` 
	for i := 0; i < len(UserData); i++ {
		div := "<div class='list'>"
		Id := "<h4>User ID:<i>" + fmt.Sprint(UserData[i].Id) + "</i></h4>"
		firstname := "<h4>Firstname:<i>" + UserData[i].Firstname + "</i></h4>"
		lastname := "<h4>Lastname:<i>" + UserData[i].Lastname + "</i></h4>"
		CreatedAt := "<h4>Created at:<i>" + fmt.Sprint(UserData[i].CreatedAt) + "</i></h4>"
		UpdatedAt := "<h4>Updated at:<i>" + fmt.Sprint(UserData[i].UpdatedAt) + "</i></h4>"
		div += Id + firstname + lastname + CreatedAt + UpdatedAt
		div += "</div>"
		template += div
	}
	template +=  `
	</div>
	</body>
	</html>
	`
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, template)
}
