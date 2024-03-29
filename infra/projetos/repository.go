package projetos

import (
	"database/sql"
	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"
	modelData "gerenciadorDeProjetos/infra/projetos/model"
	"gerenciadorDeProjetos/infra/projetos/postgres"
	utils "gerenciadorDeProjetos/utils/params"
)

type repositorio struct {
	Data *postgres.DBProjetos
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBProjetos{DB: novoDB},
	}
}

func (r *repositorio) NovoProjeto(req *modelApresentacao.ReqProjeto) (*modelApresentacao.ReqProjetos, error) {
	return r.Data.NovoProjeto(&modelData.ReqProjeto{Nome_Projeto: req.Nome_Projeto, Descricao_Projeto: req.Descricao_Projeto, Equipe_ID: req.Equipe_ID, Prazo: req.Prazo})
}
func (r *repositorio) ListarProjetos() ([]modelApresentacao.ReqProjetos, error) {
	return r.Data.ListarProjetos()
}
func (r *repositorio) ListarProjeto(id string) (*modelApresentacao.ReqProjetos, error) {
	return r.Data.ListarProjeto(id)
}
func (r *repositorio) ListarProjetosComStatus(status string) ([]modelApresentacao.ReqStatusProjeto, error) {
	return r.Data.ListarProjetosComStatus(status)
}
func (r *repositorio) ListarTasksProjeto(id string) ([]modelApresentacao.ReqTasksProjeto, error) {
	return r.Data.ListarTasksProjeto(id)
}
func (r *repositorio) DeletarProjeto(id string) error {
	return r.Data.DeletarProjeto(id)
}
func (r *repositorio) AtualizarProjeto(id string, req *modelApresentacao.ReqAtualizarProjeto) (*modelApresentacao.ReqAtualizarProjeto, error) {
	return r.Data.AtualizarProjeto(id, &modelData.ReqAtualizarProjetoData{Nome_Projeto: req.Nome_Projeto, Equipe_ID: req.EquipeID, Descricao_Projeto: req.Descricao_Projeto})
}
func (r *repositorio) AtualizarStatusProjeto(id string, req *modelApresentacao.ReqAtualizarProjeto) (*modelApresentacao.ReqAtualizarProjeto, error){
	return r.Data.AtualizarStatusProjeto(id, &modelData.ReqUpdateStatusProjeto{Status: req.Status})
} 
func (r *repositorio) ListarProjetosFiltro(params *utils.RequestParams) ([]modelApresentacao.ReqProjetos, error) {
	return r.Data.ListarProjetosFiltro(params)
}