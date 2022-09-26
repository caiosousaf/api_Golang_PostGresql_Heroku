package pessoas

import (
	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"

	utils "gerenciadorDeProjetos/utils/params"
)
type IPessoa interface {
	NovaPessoa(req *modelApresentacao.ReqPessoa) (*modelApresentacao.ReqPessoa, error)
	ListarPessoas() (*modelApresentacao.ListarGetPessoa, error)
	ListarPessoa(id string) (*modelApresentacao.ReqGetPessoa, error)
	ListarTarefasPessoa(id string) ([]modelApresentacao.ReqTarefaPessoa, error)
	AtualizarPessoa(id string, req *modelApresentacao.ReqAtualizarPessoa) (*modelApresentacao.ReqAtualizarPessoa, error)
	DeletarPessoa(id string) error
	ListarPessoasFiltro(params *utils.RequestParams) (*modelApresentacao.ListarGetPessoa, error)
}