package main

import (
	"log"

	"github.com/gauravst/got/internal/cli"
	"github.com/gauravst/got/internal/config"
	"github.com/gauravst/got/internal/database"
)

func main() {
	cfg := config.ConfigMustLoad()

	// Initialize DB once
	database.InitDB(cfg.DatabaseUri)
	defer database.CloseDB()

	// Run CLI
	cmd := cli.NewCLI(cfg)
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
