package controller

import (
	"encoding/json"
	"net/http"

	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/application"
)

type IssueController struct {
	IssueService application.IssueService
}

func (c IssueController) List(w http.ResponseWriter, r *http.Request) {
	issues, err := c.IssueService.Issues()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	issuesJson, err := json.Marshal(issues)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(issuesJson)
}

func (c IssueController) Show(w http.ResponseWriter, r *http.Request) {

}

func (c IssueController) Create(w http.ResponseWriter, r *http.Request) {

}

func (c IssueController) Delete(w http.ResponseWriter, r *http.Request) {

}
