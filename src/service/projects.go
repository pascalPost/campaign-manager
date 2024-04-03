package service

import (
	"database/sql"
	log2 "github.com/labstack/gommon/log"
	"log/slog"
)

type Project struct {
	id   uint
	name string
}

func CreateProjectsTable(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS project (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log2.Fatal(err)
	}
}

// QueryProjects returns all projects from the table (or an empty slice in case of an error)
func QueryProjects(db *sql.DB) []Project {
	rows, err := db.Query("SELECT * FROM project")
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	projects := make([]Project, 0)
	if err != nil {
		slog.Error("query projects", err)
		return projects
	}

	for rows.Next() {
		p := Project{}
		err = rows.Scan(&p.id, &p.name)
		if err != nil {
			slog.Error("query projects scan", err)
			return projects
		}
		projects = append(projects, p)
	}

	err = rows.Err()
	if err != nil {
		slog.Error("query projects scan", err)
	}

	return projects
}

//func CreateProjectStatusChangesTable(db *sql.DB) {
//	sqlStmt := `
//	CREATE TABLE IF NOT EXISTS project_status_changes (
//		id INTEGER PRIMARY KEY AUTOINCREMENT,
//		project_id INTEGER NOT NULL,
//		status TEXT NOT NULL,
//		changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//		FOREIGN KEY(project_id) REFERENCES projects(id)
//	);
//	`
//	_, err := db.Exec(sqlStmt)
//	if err != nil {
//		log2.Fatal(err)
//	}
//}
//
//func CreateProjectParametersTable(db *sql.DB) {
//	sqlStmt := `
// CREATE TABLE IF NOT EXISTS project_parameters (
//  id INTEGER PRIMARY KEY AUTOINCREMENT,
//  project_id INTEGER NOT NULL,
//  parameter_name TEXT NOT NULL,
//  parameter_value TEXT NOT NULL,
//  FOREIGN KEY(project_id) REFERENCES projects(id)
// );
// `
//	_, err := db.Exec(sqlStmt)
//	if err != nil {
//		log2.Fatal(err)
//	}
//}
//
//func StoreConfigMap() {
//
//}
