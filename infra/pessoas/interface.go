package pessoas

import (
	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
)
type IPessoa interface {
	NovaPessoa(req *modelApresentacao.ReqPessoa) (*modelApresentacao.ReqPessoa, error)
	ListarPessoas() ([]modelApresentacao.ReqGetPessoa, error)
	ListarPessoa(id string) (*modelApresentacao.ReqGetPessoa, error)
	ListarTarefasPessoa(id string) ([]modelApresentacao.ReqTarefaPessoa, error)
	AtualizarPessoa(id string, req *modelApresentacao.ReqAtualizarPessoa) (*modelApresentacao.ReqAtualizarPessoa, error)
	DeletarPessoa(id string) error
}