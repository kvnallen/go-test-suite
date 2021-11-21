//go:generate swag init -g main.go -o docs\swagger
package main

import (
	"log"
	"meu.pack/internal/infra"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/swaggo/files"
	_ "meu.pack/docs/swagger"
)

var (
	router      = chi.NewRouter()
	controllers = []infra.HttpController{
		infra.NewUserController(router),
	}
)

// @title Swagger Example API
// @version 1.0
// @description API de users MELI
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@docs.io

// @host /
// @BasePath /
func main() {
	for _, controller := range controllers {
		controller.RegisterRoutes()
	}

	log.Fatal(http.ListenAndServe(":3000", router))
}
