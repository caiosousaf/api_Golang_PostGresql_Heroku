package projetos

import "database/sql"

func NovoRepo(DB *sql.DB) IProjeto {
	return novoRepo(DB)
}