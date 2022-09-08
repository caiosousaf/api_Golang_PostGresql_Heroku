package postgres

import "database/sql"

type DBTasks struct {
	DB *sql.DB
}