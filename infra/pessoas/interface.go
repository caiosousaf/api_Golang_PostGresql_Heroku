package pessoas

import (
	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
)
type IPessoa interface {
	NovaPessoa(req *modelApresentacao.ReqPessoa) (*modelApresentacao.ReqPessoa, error)
	ListarPessoas() (*modelApresentacao.ListarGetPessoa, error)
	ListarPessoa(id *int64) (*modelApresentacao.ReqGetPessoa, error)
	ListarTarefasPessoa(id *int64) ([]modelApresentacao.ReqTarefaPessoa, error)
	AtualizarPessoa(id *int64, req *modelApresentacao.ReqAtualizarPessoa) (*modelApresentacao.ReqAtualizarPessoa, error)
	DeletarPessoa(id *int64) error
}