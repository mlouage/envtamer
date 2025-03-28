package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newPullCmd() *cobra.Command {
	var filename string

	cmd := &cobra.Command{
		Use:   "pull DIRECTORY_NAME",
		Short: "Pull environment variables from the database to a local .env file",
		Long:  `This command retrieves stored environment variables for the specified directory from the database and writes them to a local .env file. If the file already exists, it will prompt for confirmation before overwriting.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("This is the pull command which will pull the env vars from %s to a local file called %s\n",
			args[0], filename)
		},
	}

	cmd.Flags().StringVarP(&filename, "filename", "f", ".env", "The name of the env file to create or update")
	return cmd
}
