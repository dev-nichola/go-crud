package main

import (
	"net/http"

	"github.com/dev-nichola/go-crud/database"
	"github.com/dev-nichola/go-crud/routes"
)

func main() {
	database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server)

	http.ListenAndServe(":8000", server)
}
