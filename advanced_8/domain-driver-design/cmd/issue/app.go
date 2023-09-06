package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/application"
	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/domain"
	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/persistence/db"
	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/web/controller"
)

func main() {
	issueRepo := db.NewIssueRepository()

	issueService := application.IssueService{
		IssueRespository: *issueRepo,
	}

	issueController := controller.IssueController{
		IssueService: issueService,
	}

	for i := 0; i < 10; i += 1 {
		issueService.Create(
			&domain.Issue{
				IssueTitle:       fmt.Sprintf("Issue_%d", i),
				IssueDescription: "Task1",
				IssueOwnerId:     1,
				IssueProjectId:   1,
				IssueStatus:      domain.StatusDone,
				IssuePriority:    domain.PriorityHigh,
			})
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", issueController.List)

	server := &http.Server{
		Addr:           ":8092",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
