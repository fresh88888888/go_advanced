package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/application"
	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/domain"
	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/persistence/memory"
	"umbrella.github.com/advanced_go/advanced_8/domain-driver-design/web/controller"
)

func main() {
	userRepo := memory.NewUserRepository()
	userService := application.UserService{
		UserRespoitory: userRepo,
	}

	userController := controller.UserController{
		UserService: userService,
	}

	for i := 0; i < 10; i++ {
		userService.Create(&domain.User{
			Id:   int64(i),
			Name: fmt.Sprintf("User_%d", i),
		})
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/", userController.List)
	server := &http.Server{
		Addr:           ":8091",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
