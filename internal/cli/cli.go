package cli

import (
	"github.com/gauravst/got/internal/cli/commands"
	"github.com/spf13/cobra"
)

func NewCLI() *cobra.Command {
	root := &cobra.Command{
		Use:   "mydb",
		Short: "MyDB CLI Tool",
	}

	root.AddCommand(commands.Start())
	return root
}
