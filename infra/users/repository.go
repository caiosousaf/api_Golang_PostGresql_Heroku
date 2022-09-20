package users

import (
	"database/sql"
	"gerenciadorDeProjetos/infra/users/postgres"
	modelApresentacao "gerenciadorDeProjetos/domain/users/model"
	modelData "gerenciadorDeProjetos/infra/users/model"
)

type repositorio struct {
	Data *postgres.DBUsers
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBUsers{DB: novoDB},
	}
}

func (r *repositorio) NovoUsuario(req *modelApresentacao.ReqUser) (*modelApresentacao.ReqUser, error) {
	return r.Data.NovoUsuario(&modelData.ReqUser{Nome: req.Nome, Email: req.Email, Password: req.Password})
}