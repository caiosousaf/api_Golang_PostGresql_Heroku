package equipes

import (
	"database/sql"
)

func NovoRepo(DB *sql.DB) IEquipe {
	return novoRepo(DB)
}