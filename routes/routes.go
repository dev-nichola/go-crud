package routes

import (
	"database/sql"
	"net/http"

	"github.com/dev-nichola/go-crud/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.NewHelloWorldController())

	server.HandleFunc("/employe", controller.NewIndexEmployeController())
	server.HandleFunc("/employe/create", controller.NewCreateEmployeController(db))

}
