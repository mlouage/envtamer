package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

// Storage represents a SQLite database connection
type Storage struct {
	db *sql.DB
	empty bool
}

// New creates a new Storage instance
func New() (*Storage, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("🛑 Failed to get user's home directory: %w", err)
	}

	dbDir := filepath.Join(home, ".envtamer")
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return nil, fmt.Errorf("🛑 Failed to create the database directory: %w", err)
	}

	dbPath := filepath.Join(dbDir, "envtamer.db")

	newDb := false

	if _, err := os.Stat(dbPath); errors.Is(err, os.ErrNotExist) {
		newDb = true
		fmt.Println("🗄️ Empty database file created successfully")
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("🛑 Failed to open database: %w", err)
	}

	return &Storage{db: db, empty: newDb}, nil
}

func (s *Storage) Init() error {
	if s.empty {
		fmt.Println("⏳ Running migrations...")
		_, err := s.db.Exec(`
			CREATE TABLE IF NOT EXISTS "EnvVariables" (
							"Directory" TEXT NOT NULL,
							"Key" TEXT NOT NULL,
							"Value" TEXT NOT NULL,
							CONSTRAINT "PK_EnvVariables" PRIMARY KEY ("Directory", "Key")
					);
		`)
		if err != nil {
			return fmt.Errorf("🛑 Failed to initialize database table: %w", err)
		}

		fmt.Println("✅ Migrations applied successfully")
		fmt.Println("🚀 Ready to push and pull env files")
	} else {
		fmt.Println("🛑 Database file already exists. Initialization skipped")
	}

	return nil
}

// Close closes the database connection
func (s *Storage) Close() error {
	return s.db.Close()
}

// SaveEnvVars saves environment variables for a directory
func (s *Storage) SaveEnvVars(directory string, envVars map[string]string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Delete existing env vars for this directory
	_, err = tx.Exec("DELETE FROM EnvVariables WHERE Directory = ?", directory)
	if err != nil {
		return fmt.Errorf("failed to delete existing env vars: %w", err)
	}

	// Insert new env vars
	stmt, err := tx.Prepare("INSERT INTO EnvVariables (Directory, Key, Value) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	for key, value := range envVars {
		_, err = stmt.Exec(directory, key, value)
		if err != nil {
			return fmt.Errorf("failed to insert env var: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
