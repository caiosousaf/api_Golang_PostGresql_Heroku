package users

import "database/sql"

func NovoRepo(DB *sql.DB) IUser {
	return novoRepo(DB)
}