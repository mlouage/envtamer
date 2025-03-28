package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [DIRECTORY_NAME]",
		Short: "List stored directories or environment variables",
		Long:  `If no directory is specified, this command lists all directories stored in the database. If a directory is provided, it lists all environment variables stored for that directory.`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if (len(args) == 0) {
				fmt.Println("This is the list command")
			} else {
				fmt.Printf("This is the list command which will list the env for directory %s\n", args[0])
			}
		},
	}

	return cmd
}
