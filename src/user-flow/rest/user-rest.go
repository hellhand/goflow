package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"user-flow/messaging"
	"user-flow/model"
	"user-flow/user-repository"
)

// our main function
func UserRest(router *mux.Router) {

	go func(router *mux.Router) {
		router.HandleFunc("/", Index).Methods("GET")
		router.HandleFunc("/users", GetUsers).Methods("GET").Headers("Content-Type", "application/json")
		router.HandleFunc("/users/{id}", GetUser).Methods("GET").Headers("Content-Type", "application/x-www-form-urlencoded")
		router.HandleFunc("/users", CreateUser).Methods("POST").Headers("Content-Type", "application/json")
		router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT").Headers("Content-Type", "application/json")
		router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE").Headers("Content-Type", "application/x-www-form-urlencoded")
	}(router)
}
func Index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Welcome!")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := user_repository.FindAll()
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	_ = json.NewEncoder(w).Encode(users)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	user, err := user_repository.Find(userId)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	_ = json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	messaging.Write(&user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	messaging.Write(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	user_repository.Delete(userId)
}
