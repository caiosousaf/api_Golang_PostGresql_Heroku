package users

import (
	modelApresentacao "gerenciadorDeProjetos/domain/users/model"
)
type IUser interface {
	NovoUsuario(req *modelApresentacao.ReqUser) (*modelApresentacao.ReqUser, error)
}