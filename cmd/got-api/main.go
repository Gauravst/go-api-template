package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gauravst/got/internal/api/router"
	"github.com/gauravst/got/internal/config"
	"github.com/gauravst/got/internal/database"
	"github.com/gauravst/got/internal/repositories"
	"github.com/gauravst/got/internal/services"
)

func main() {
	// load config
	cfg := config.ConfigMustLoad()

	// database setup
	database.InitDB(cfg.DatabaseUri)
	defer database.CloseDB()

	// Repositories and Services
	userRepo := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepo)

	authRepo := repositories.NewAuthRepository(database.DB)
	authService := services.NewAuthService(authRepo)

	// Grouped services for injection
	allServices := &router.AllServices{
		UserService: userService,
		AuthService: authService,
		// add more services here
	}

	// Create router
	finalHandler := router.NewRouter(cfg, allServices)

	// Setup server
	addr := cfg.Address
	if cfg.HTTPServer.Port != 0 {
		addr = "0.0.0.0:" + strconv.Itoa(cfg.HTTPServer.Port)
	}

	server := &http.Server{
		Addr:    addr,
		Handler: finalHandler,
	}

	slog.Info("server started", slog.String("address", cfg.Address))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("faild to Shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server Shutdown successfully")
}
