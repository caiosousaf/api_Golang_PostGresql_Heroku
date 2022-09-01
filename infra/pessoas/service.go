package pessoas

import (
	"database/sql"
)

func NovoRepo(DB *sql.DB) IPessoa {
	return novoRepo(DB)
}