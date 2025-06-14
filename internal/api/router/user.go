package router

import (
	"net/http"

	"github.com/gauravst/got/internal/api/handlers"
)

func RegisterUserRoutes(mux *http.ServeMux, services *AllServices) {
	mux.HandleFunc("GET /api/user", handlers.CreateUser(services.UserService))
}
