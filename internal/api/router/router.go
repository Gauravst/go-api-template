package router

import (
	"net/http"

	"github.com/gauravst/got/internal/api/middleware"
	"github.com/gauravst/got/internal/config"
	"github.com/gauravst/got/internal/services"
)

type AllServices struct {
	UserService services.UserService
	AuthService services.AuthService
}

func NewRouter(cfg *config.Config, services *AllServices) http.Handler {
	publicRouter := http.NewServeMux()
	protectedRouter := http.NewServeMux()

	// Register protected routes
	RegisterUserRoutes(protectedRouter, services)

	// Register public router

	// add middleware if needed according to route
	mainRouter := http.NewServeMux()
	mainRouter.Handle("/api/auth/", publicRouter)
	mainRouter.Handle("/", middleware.Auth(cfg, services.AuthService)(protectedRouter))

	return middleware.CORS(cfg)(mainRouter)
}
