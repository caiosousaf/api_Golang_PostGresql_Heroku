package projetos

import (
	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"
)

type IProjeto interface {
	NovoProjeto(req *modelApresentacao.ReqProjeto) (*modelApresentacao.ReqProjetos, error)
	ListarProjetos() ([]modelApresentacao.ReqProjetos, error)
	ListarProjeto(id string) (*modelApresentacao.ReqProjetos, error)
	ListarProjetosComStatus(status string) ([]modelApresentacao.ReqStatusProjeto, error)
	ListarTasksProjeto(id string) ([]modelApresentacao.ReqTasksProjeto, error)
	DeletarProjeto(id string) error
	AtualizarProjeto(id string, req *modelApresentacao.ReqAtualizarProjeto) (*modelApresentacao.ReqAtualizarProjeto, error)
	AtualizarStatusProjeto(id string, req *modelApresentacao.ReqAtualizarProjeto) (*modelApresentacao.ReqAtualizarProjeto, error)
}