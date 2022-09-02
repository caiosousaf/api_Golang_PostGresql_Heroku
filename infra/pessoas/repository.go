package pessoas

import (
	"database/sql"

	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
	modelData "gerenciadorDeProjetos/infra/pessoas/model"
	"gerenciadorDeProjetos/infra/pessoas/postgres"
	"github.com/gin-gonic/gin"
)

type repositorio struct {
	Data *postgres.DBPessoas
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBPessoas{DB: novoDB},
	}
}

func (r *repositorio) NovaPessoa(req *modelApresentacao.ReqPessoa, c *gin.Context) {
	r.Data.NovaPessoa(&modelData.ReqPessoa{Nome_Pessoa: req.Nome_Pessoa, 
		Funcao_Pessoa: req.Funcao_Pessoa, Equipe_ID: req.Equipe_ID}, c)
}
func (r *repositorio) ListarPessoas() ([]modelApresentacao.ReqGetPessoa, error) {
	return r.Data.ListarPessoas()
}

