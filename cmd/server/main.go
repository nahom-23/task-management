package main

import (
	"log"
	"net/http"

	"task-management/internal/config"
	"task-management/internal/handler"
	"task-management/internal/repository"
	"task-management/internal/service"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	cfg := config.LoadConfig()
	db := repository.ConnectDB(cfg)
	defer db.Close()
	
	r := mux.NewRouter()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	r.HandleFunc("/users/{national_id}", userHandler.GetUserByNationalID).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}