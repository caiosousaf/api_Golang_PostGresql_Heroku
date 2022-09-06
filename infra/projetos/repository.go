package projetos

import (
	"database/sql"
	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"
	modelData "gerenciadorDeProjetos/infra/projetos/model"
	"gerenciadorDeProjetos/infra/projetos/postgres"

	"github.com/gin-gonic/gin"
)

type repositorio struct {
	Data *postgres.DBProjetos
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBProjetos{DB: novoDB},
	}
}

func (r *repositorio) NovoProjeto(req *modelApresentacao.ReqProjeto, c *gin.Context) {
	r.Data.NovoProjeto(&modelData.ReqProjeto{Nome_Projeto: req.Nome_Projeto, Descricao_Projeto: req.Descricao_Projeto, Equipe_ID: req.Equipe_ID, Prazo: req.Prazo}, c)
}
func (r *repositorio) ListarProjetos() ([]modelApresentacao.ReqProjetos, error) {
	return r.Data.ListarProjetos()
}
func (r *repositorio) ListarProjeto(id string) (*modelApresentacao.ReqProjetos, error) {
	return r.Data.ListarProjeto(id)
}