package command

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mlouage/envtamer-go/internal/storage"
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

func parseEnvFile(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	envVars := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove quotes if present
		if len(value) > 1 && (value[0] == '"' || value[0] == '\'') && value[0] == value[len(value)-1] {
			value = value[1 : len(value)-1]
		}

		envVars[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return envVars, nil
}

func newPushCmd() *cobra.Command {
	var filename string

	cmd := &cobra.Command{
		Use:   "push [DIRECTORY_NAME]",
		Short: "Push the contents of a local .env file to the database",
		Long:  `This command reads the specified .env file and stores its contents in the database, associated with the given directory.`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Resolve directory path
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

			// Parse .env file
			envFilePath := filepath.Join(dirPath, filename)
			envVars, err := parseEnvFile(envFilePath)
			if err != nil {
				return fmt.Errorf("failed to parse env file: %w", err)
			}

			// Save to database
			db, err := storage.New()
			if err != nil {
				return fmt.Errorf("failed to create storage: %w", err)
			}
			defer db.Close()

			if err := db.SaveEnvVars(dirPath, envVars); err != nil {
				return fmt.Errorf("failed to save env vars: %w", err)
			}

			fmt.Printf("Successfully pushed %d environment variables for directory: %s\n", len(envVars), dirPath)
			return nil
		},
	}

	cmd.Flags().StringVarP(&filename, "filename", "f", ".env", "The name of the env file")
	return cmd
}
