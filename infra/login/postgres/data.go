package postgres

import (
	"database/sql"
	modelApresentacao "gerenciadorDeProjetos/domain/login/model"
	modelData "gerenciadorDeProjetos/infra/login/model"
)

type DBLogin struct {
	DB *sql.DB
}

func (postgres *DBLogin) LoginUsuario(req *modelData.ReqLogin) (*modelApresentacao.Login, error) {
	sqlStatement := `SELECT * FROM users WHERE email = $1::VARCHAR(80);`
	user := modelApresentacao.Login{}
	row := postgres.DB.QueryRow(sqlStatement, req.Email)

	if err := row.Scan(&user.ID_Usuario,&user.Nome, &user.Email, &user.Password, &user.Data_Criacao); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}

	if user.Password != req.Password {
		return nil, nil
	}

	return &user, nil
}