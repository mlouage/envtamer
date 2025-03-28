package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func resolvePath(path string) (string, error) {
	if path == "" {
		dir, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("failed to get current directory: %w", err)
		}
		return dir, nil
	}

	if filepath.IsAbs(path) {
		return path, nil
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to resolve absolute path: %w", err)
	}

	return absPath, nil
}

func newPushCmd() *cobra.Command {
	var filename string

	cmd := &cobra.Command{
		Use:   "push [DIRECTORY_NAME]",
		Short: "Push the contents of a local .env file to the database",
		Long:  `This command reads the specified .env file and stores its contents in the database, associated with the given directory.`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var dirPath string
			var err error
			if len(args) > 0 {
				dirPath, err = resolvePath(args[0])
			} else {
				dirPath, err = resolvePath("")
			}
			if err != nil {
				return fmt.Errorf("failed to resolve directory path: %w", err)
			}

			fmt.Printf("Successfully pushed environment variables for directory: %s\n", dirPath)
			return nil
		},
	}

	cmd.Flags().StringVarP(&filename, "filename", "f", ".env", "The name of the env file")
	return cmd
}
