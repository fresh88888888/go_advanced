package controller

import (
	"net/http"

	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/domain"
)

// Controller for Project model
type ProjectController struct {
	ProjectService domain.ProjectService
}

func (c ProjectController) List(w http.ResponseWriter, r *http.Request) {

}

func (c ProjectController) Show(w http.ResponseWriter, r *http.Request) {

}

func (c ProjectController) Create(w http.ResponseWriter, r *http.Request) {

}

func (c ProjectController) Delete(w http.ResponseWriter, r *http.Request) {

}
