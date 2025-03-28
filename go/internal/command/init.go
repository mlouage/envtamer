package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize an empty database in the user's home folder",
		Long:  `This command creates an empty SQLite database file named 'envtamer-go.db' in the '.envtamer-go' directory of the user's home folder. If the file already exists, the command will not overwrite it.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is the init command")
		},
	}

	return cmd
}
