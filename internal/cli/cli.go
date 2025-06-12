package cli

import (
	"fmt"
	"os"

	"github.com/gauravst/got/internal/cli/commands"
	"github.com/gauravst/got/internal/cli/middleware"
	"github.com/gauravst/got/internal/config"
	"github.com/spf13/cobra"
)

var (
	configPath string
	debug      bool
)

func NewCLI(cfg *config.Config) *cobra.Command {
	root := &cobra.Command{
		Use:    "mycli",
		Short:  "CLI Tool",
		PreRun: middleware.CheckAuth(),
		// Global Flags function
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if debug {
				fmt.Println("Debug mode is ON")
			}

			if configPath != "" {
				os.Setenv("CONFIG_PATH", configPath)
				fmt.Println("Using config:", configPath)
			}
		},
	}

	// Persistent (Global) Flags, work with all commands
	root.PersistentFlags().StringVar(&configPath, "config", "", "Path to config file")
	root.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug mode")

	// Pass config to all commands
	root.AddCommand(commands.Start(cfg))

	return root
}
