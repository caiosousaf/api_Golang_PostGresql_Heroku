package postgres

import (
	"database/sql"

	modelApresentacao "gerenciadorDeProjetos/domain/users/model"
	modelData "gerenciadorDeProjetos/infra/users/model"
)

type DBUsers struct {
	DB *sql.DB
}

func (postgres *DBUsers) NovoUsuario(req *modelData.ReqUser) (*modelApresentacao.ReqUser, error) {
	sqlStatement := `INSERT INTO users(nome, email, password) VALUES($1, $2, $3) RETURNING id_usuario, nome, email`

	var user = &modelApresentacao.ReqUser{}

	row := postgres.DB.QueryRow(sqlStatement, req.Nome, req.Email, req.Password)

	if err := row.Scan(&user.ID_Usuario, &user.Nome, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}

	return user, nil
}
