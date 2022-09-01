package equipes

import (
	"database/sql"

	modelApresentacao "gerenciadorDeProjetos/domain/equipes/model"
	modelData "gerenciadorDeProjetos/infra/equipes/model"
	"gerenciadorDeProjetos/infra/equipes/postgres"
	"github.com/gin-gonic/gin"
)

type repositorio struct {
	Data *postgres.DBEquipes
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBEquipes{DB: novoDB},
	}
}

func (r *repositorio) NovaEquipe(req *modelApresentacao.ReqEquipe, c *gin.Context) {
	r.Data.NovaEquipe(&modelData.Equipe{Nome_Equipe: req.Nome_Equipe}, c)
}
func (r *repositorio) ListarEquipes() ([]modelApresentacao.ReqEquipe, error) {
	return r.Data.ListarEquipes()
}
func (r *repositorio) BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error) {
	return r.Data.BuscarEquipe(id)
}
func (r *repositorio) BuscarMembrosDeEquipe(id string) ([]modelApresentacao.ReqEquipeMembros, error) {
	return r.Data.BuscarMembrosDeEquipe(id)
}
func (r *repositorio) BuscarProjetosDeEquipe(id string) ([]modelApresentacao.ReqEquipeProjetos, error) {
	return r.Data.BuscarProjetosDeEquipe(id)
}
func (r *repositorio) DeletarEquipe(id string) (*modelApresentacao.ReqEquipe, error) {
	return r.Data.DeletarEquipe(id)
}