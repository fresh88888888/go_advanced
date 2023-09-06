package controller

import (
	"encoding/json"
	"net/http"

	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/application"
)

// Controller for User model
type UserController struct {
	UserService application.UserService
}

func (u *UserController) List(w http.ResponseWriter, r *http.Request) {
	users, err := u.UserService.Users()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}
func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {

}

func (u *UserController) Show(w http.ResponseWriter, r *http.Request) {

}

func (u *UserController) Delete(w http.ResponseWriter, r *http.Request) {
}
