package db

import (
	"database/sql"
	"errors"
	types "github.com/campaign-manager/src/types"
	"log"
	"log/slog"
)

// CreateSettingsTable creates the setting table
func CreateSettingsTable(db *sql.DB) {
	// limit to single row for now
	// TODO remove plain text password. Will be replaced with general token.
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS setting (
		id INTEGER PRIMARY KEY CHECK (id = 0),
		workingDir TEXT NOT NULL,
		lsfUsername TEXT NOT NULL,
		lsfPassword TEXT NOT NULL
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}

// GetSettings returns the settings from the table
func GetSettings(db *sql.DB) (*types.Settings, error) {
	var workingDir string
	var lsfUserName string
	var lsfPassword string

	id := 1
	err := db.QueryRow("SELECT workingDir, lsfUsername, lsfPassword FROM setting WHERE id = ?;", id).Scan(workingDir, lsfUserName, lsfPassword)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return types.NewSettings(workingDir, lsfUserName, lsfPassword), nil
}

// SetSettings saves the given settings to the db
func SetSettings(db *sql.DB, s *types.Settings) error {
	id := 1
	sqlStmt := "INSERT OR REPLACE INTO setting (id, workingDir, lsfUsername, lsfPassword) VALUES (?, ?, ?, ?);"

	_, err := db.Exec(sqlStmt, id, s.WorkingDir(), s.LSFUsername(), s.LSFPassword())

	slog.Error("Error in setting settings in the db", s, err)

	return err
}
