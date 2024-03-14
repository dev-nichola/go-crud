package routes

import (
	"net/http"

	"github.com/dev-nichola/go-crud/controller"
)

func MapRoutes(server *http.ServeMux) {
	server.HandleFunc("/", controller.NewHelloWorldController())

}
