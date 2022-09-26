package pessoas

import (
	"database/sql"

	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
	modelData "gerenciadorDeProjetos/infra/pessoas/model"
	"gerenciadorDeProjetos/infra/pessoas/postgres"
	utils "gerenciadorDeProjetos/utils/params"
)

type repositorio struct {
	Data *postgres.DBPessoas
}

func novoRepo(novoDB *sql.DB) *repositorio {
	return &repositorio{
		Data: &postgres.DBPessoas{DB: novoDB},
	}
}

func (r *repositorio) NovaPessoa(req *modelApresentacao.ReqPessoa) (*modelApresentacao.ReqPessoa, error) {
	return r.Data.NovaPessoa(&modelData.ReqPessoa{Nome_Pessoa: req.Nome_Pessoa, 
		Funcao_Pessoa: req.Funcao_Pessoa, Equipe_ID: req.Equipe_ID})
}
func (r *repositorio) ListarPessoas() (*modelApresentacao.ListarGetPessoa, error) {
	return r.Data.ListarPessoas()
}
func (r *repositorio) ListarPessoa(id *int64) (*modelApresentacao.ReqGetPessoa, error) {
	return r.Data.ListarPessoa(id)
}
func (r *repositorio) ListarTarefasPessoa(id *int64) ([]modelApresentacao.ReqTarefaPessoa, error) {
	return r.Data.ListarTarefasPessoa(id)
}
func (r *repositorio) AtualizarPessoa(id *int64, req *modelApresentacao.ReqAtualizarPessoa) (*modelApresentacao.ReqAtualizarPessoa, error) {
	return r.Data.AtualizarPessoa(id, &modelData.ReqPessoa{Nome_Pessoa: req.Nome_Pessoa, Funcao_Pessoa: req.Funcao_Pessoa, Equipe_ID: req.Equipe_ID})
}
func (r *repositorio) DeletarPessoa(id *int64) error {
	return r.Data.DeletarPessoa(id)
}
func (r *repositorio) ListarPessoasFiltro(params *utils.RequestParams) (*modelApresentacao.ListarGetPessoa, error) {
	return r.Data.ListarPessoasFiltro(params)
}