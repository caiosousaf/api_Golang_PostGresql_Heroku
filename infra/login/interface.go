package login

import (
	modelApresentacao "gerenciadorDeProjetos/domain/login/model"
)
type ILogin interface {
	LoginUsuario(req *modelApresentacao.Login) (*modelApresentacao.Login, error)
}