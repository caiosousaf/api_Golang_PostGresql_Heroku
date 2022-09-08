package tasks

import (
	"database/sql"
	"gerenciadorDeProjetos/infra/tasks/postgres"
)

type repositorio struct {
	Data *postgres.DBTasks
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBTasks{DB: novoDB},
	}
}
