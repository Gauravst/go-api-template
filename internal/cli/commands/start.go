package commands

import (
	"fmt"

	"github.com/gauravst/got/internal/cli/middleware"
	"github.com/gauravst/got/internal/config"
	"github.com/spf13/cobra"
)

func Start(cfg *config.Config) *cobra.Command {
	var port int
	var debug bool

	cmd := &cobra.Command{
		Use:    "start",
		Short:  "Start the mycli server",
		PreRun: middleware.CheckSomething(),
		Run: func(cmd *cobra.Command, args []string) {
			// run you Start logic here with flag or without flag
			// Use the flag values
			fmt.Println("Start with port:", port)
			fmt.Println("Debug mode:", debug)
		},
	}

	// Add local flags
	cmd.Flags().IntVar(&port, "port", 8080, "Port to run the server on")
	cmd.Flags().BoolVar(&debug, "debug", false, "Enable debug mode")

	return cmd
}
