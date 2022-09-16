package pessoas

import (
	"fmt"
	"gerenciadorDeProjetos/config/database"
	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
	"gerenciadorDeProjetos/infra/pessoas"
)

func NovaPessoa(req *modelApresentacao.ReqPessoa) (*modelApresentacao.ReqPessoa, error) {
	db := database.Conectar()
	defer db.Close()
	pessoasRepo := pessoas.NovoRepo(db)

	return pessoasRepo.NovaPessoa(req)
}

func ListarPessoas() ([]modelApresentacao.ReqGetPessoa, error) {
	db := database.Conectar()
	defer db.Close()

	pessoasRepo := pessoas.NovoRepo(db)
	return pessoasRepo.ListarPessoas()
}

func ListarPessoa(id string) (*modelApresentacao.ReqGetPessoa, error) {
	db := database.Conectar()
	defer db.Close()

	pessoasRepo := pessoas.NovoRepo(db)
	return pessoasRepo.ListarPessoa(id)
}

func ListarTarefasPessoa(id string) (res []modelApresentacao.ReqTarefaPessoa, err error) {
	db := database.Conectar()
	defer db.Close()

	pessoasRepo := pessoas.NovoRepo(db)

	dados, err := pessoasRepo.ListarPessoa(id)
	//if len(dados) == 0 {
	if err != nil {
		return res, fmt.Errorf("pessoa não Encontrada")
	}

	if dados == nil {
		return res, fmt.Errorf("em Aberto")
	}

	res, err = pessoasRepo.ListarTarefasPessoa(id)
	if err != nil {
		return nil, fmt.Errorf("não foi possivel buscar tarefas de uma pessoa")
	}
	return
}

func AtualizarPessoa(id string, req *modelApresentacao.ReqAtualizarPessoa) (res *modelApresentacao.ReqAtualizarPessoa, err error) {
	db := database.Conectar()
	defer db.Close()

	pessoasRepo := pessoas.NovoRepo(db)

	dados, err := pessoasRepo.ListarPessoa(id)
	//if len(dados) == 0 {
	if err != nil {
		return res, fmt.Errorf("person does not exist")
	}

	if dados == nil {
		return res, fmt.Errorf("em Aberto")
	}

	res, err = pessoasRepo.AtualizarPessoa(id, req)
	if err != nil {
		return nil, fmt.Errorf("team does not exist")
	}
	return
}

func DeletarPessoa(id string) (err error) {
	db := database.Conectar()
	defer db.Close()

	pessoasRepo := pessoas.NovoRepo(db)

	dados, err := pessoasRepo.ListarPessoa(id)
	//if len(dados) == 0 {
	if err != nil {
		return fmt.Errorf("pessoa não Encontrada")
	}

	if dados == nil {
		return fmt.Errorf("em Aberto")
	}
	err = pessoasRepo.DeletarPessoa(id)
	return
}