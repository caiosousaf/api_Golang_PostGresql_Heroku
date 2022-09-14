package login

import (
	"gerenciadorDeProjetos/config/database"
	"gerenciadorDeProjetos/config/services"
	"gerenciadorDeProjetos/infra/login"
	modelApresentacao "gerenciadorDeProjetos/domain/login/model"
)

func LoginUsuario(req *modelApresentacao.Login) (*modelApresentacao.Login, error) {
	db := database.Conectar()
	defer db.Close()

	req.Password = services.SHAR256Encoder(req.Password)

	loginRepo := login.NovoRepo(db)
	return loginRepo.LoginUsuario(req) 
}