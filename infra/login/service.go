package login

import "database/sql"

func NovoRepo(DB *sql.DB) ILogin {
	return novoRepo(DB)
}