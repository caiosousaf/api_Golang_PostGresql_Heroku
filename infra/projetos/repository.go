package projetos

import (
	"database/sql"
	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"
	modelData "gerenciadorDeProjetos/infra/projetos/model"
	"gerenciadorDeProjetos/infra/projetos/postgres"

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
func (r *repositorio) DeletarProjeto(id string) error {
	return r.Data.DeletarProjeto(id)
}