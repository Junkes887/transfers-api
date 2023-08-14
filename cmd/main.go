package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Junkes887/transfers-api/internal/adpters/database"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	database.NewConfigMySql()

	routes := chi.NewRouter()

	fmt.Println("Transfers API run port " + port)
	http.ListenAndServe(port, routes)
}
