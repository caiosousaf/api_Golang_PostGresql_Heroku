package login

import (
	"database/sql"
	"gerenciadorDeProjetos/infra/login/postgres"
	modelApresentacao "gerenciadorDeProjetos/domain/login/model"
	modelData "gerenciadorDeProjetos/infra/login/model"
)

type repositorio struct {
	Data *postgres.DBLogin
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBLogin{DB: novoDB},
	}
}

func (r *repositorio) LoginUsuario(req *modelApresentacao.Login) (*modelApresentacao.Login, error) {
	return r.Data.LoginUsuario(&modelData.ReqLogin{Email: req.Email, Password: req.Password})
}