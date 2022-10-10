package equipes

import (
	modelApresentacao "gerenciadorDeProjetos/domain/equipes/model"
	modelPessoa "gerenciadorDeProjetos/domain/pessoas/model"
	utils "gerenciadorDeProjetos/utils/params"
)

type IEquipe interface {
	NovaEquipe(req *modelApresentacao.ReqEquipe) (*modelApresentacao.ReqEquipe, error)
	ListarEquipes() ([]modelApresentacao.ReqEquipe, error)
	BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error)
	BuscarMembrosDeEquipe(id string) ([]modelPessoa.ReqMembros, error)
	BuscarProjetosDeEquipe(id string) ([]modelApresentacao.ReqEquipeProjetos, error)
	BuscarTasksDeEquipe(id string) ([]modelApresentacao.ReqTasksbyTeam, error)
	DeletarEquipe(id string) error
	AtualizarEquipe(id string, req *modelApresentacao.ReqEquipe) (*modelApresentacao.ReqEquipe, error)
	ListarEquipesFiltro(params *utils.RequestParams) ([]modelApresentacao.ReqEquipe, error)
}