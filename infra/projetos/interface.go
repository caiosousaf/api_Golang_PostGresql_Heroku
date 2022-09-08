package projetos

import (
	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"
)

type IProjeto interface {
	NovoProjeto(req *modelApresentacao.ReqProjeto) (*modelApresentacao.ReqProjetos, error)
	ListarProjetos() ([]modelApresentacao.ReqProjetos, error)
	ListarProjeto(id string) (*modelApresentacao.ReqProjetos, error)
	ListarProjetosComStatus(status string) ([]modelApresentacao.ReqStatusProjeto, error)
	DeletarProjeto(id string) error
}