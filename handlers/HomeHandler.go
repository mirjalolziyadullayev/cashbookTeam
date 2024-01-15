package handlers

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Home(w, r)
	case "POST":
	case "PUT":
	case "DELETE":

	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	
}
