package tasks

import "database/sql"

func NovoRepo(DB *sql.DB) ITask {
	return novoRepo(DB)
}