package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Junkes887/transfers-api/internal/adpters/database"
	"github.com/Junkes887/transfers-api/internal/adpters/database/repository"
	"github.com/Junkes887/transfers-api/internal/adpters/web"
	"github.com/Junkes887/transfers-api/internal/domain/usecase"
	"github.com/Junkes887/transfers-api/pkg/middleware"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	configMySql := database.NewConfigMySql()
	repository := repository.NewRepository(configMySql)
	useCase := usecase.NewUseCase(repository, repository)
	handler := web.NewHandler(useCase, useCase, useCase)

	routes := chi.NewRouter()

	routes.Use(middleware.ValidateRoute)

	routes.Get("/accounts", handler.GetAllAccount)
	routes.Post("/accounts", handler.CreateAccount)
	routes.Get("/accounts/{account_id}/balance", handler.GetBalance)

	routes.Post("/login", handler.Login)

	routes.Post("/transfers", handler.CreateTransfer)
	routes.Get("/transfers", handler.GetTransfer)

	fmt.Println("Transfers API run port " + port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), routes)
}
