package pessoas

import (
	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"

	utils "gerenciadorDeProjetos/utils/params"
)
type IPessoa interface {
	NovaPessoa(req *modelApresentacao.ReqPessoa) (*modelApresentacao.ReqPessoa, error)
	ListarPessoas() (*modelApresentacao.ListarGetPessoa, error)
<<<<<<< HEAD
	ListarPessoa(id string) (*modelApresentacao.ReqGetPessoa, error)
	ListarTarefasPessoa(id string) ([]modelApresentacao.ReqTarefaPessoa, error)
	AtualizarPessoa(id string, req *modelApresentacao.ReqAtualizarPessoa) (*modelApresentacao.ReqAtualizarPessoa, error)
	DeletarPessoa(id string) error
	ListarPessoasFiltro(params *utils.RequestParams) (*modelApresentacao.ListarGetPessoa, error)
=======
	ListarPessoa(id *int64) (*modelApresentacao.ReqGetPessoa, error)
	ListarTarefasPessoa(id *int64) ([]modelApresentacao.ReqTarefaPessoa, error)
	AtualizarPessoa(id *int64, req *modelApresentacao.ReqAtualizarPessoa) (*modelApresentacao.ReqAtualizarPessoa, error)
	DeletarPessoa(id *int64) error
>>>>>>> 91ec8327816946fa99cde1dd4e53e69a84ca1097
}