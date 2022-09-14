package users

import (
	"gerenciadorDeProjetos/config/database"
	modelApresentacao "gerenciadorDeProjetos/domain/users/model"
	"gerenciadorDeProjetos/infra/users"
	"gerenciadorDeProjetos/config/services"
)

func NovoUsuario(req *modelApresentacao.ReqUser) (*modelApresentacao.ReqUser, error) {
	db := database.Conectar()
	defer db.Close()
	usersRepo := users.NovoRepo(db)

	req.Password = services.SHAR256Encoder(req.Password)
	return usersRepo.NovoUsuario(req)
}