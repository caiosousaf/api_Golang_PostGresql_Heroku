package pessoas

import (
	"fmt"
	"gerenciadorDeProjetos/config/database"
	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
	"gerenciadorDeProjetos/infra/pessoas"
	utils "gerenciadorDeProjetos/utils/params"
)

func NovaPessoa(req *modelApresentacao.ReqPessoa) (*modelApresentacao.ReqPessoa, error) {
	db := database.Conectar()
	defer db.Close()
	pessoasRepo := pessoas.NovoRepo(db)

	return pessoasRepo.NovaPessoa(req)
}

func ListarPessoas() (*modelApresentacao.ListarGetPessoa, error) {
	db := database.Conectar()
	defer db.Close()

	pessoasRepo := pessoas.NovoRepo(db)
	return pessoasRepo.ListarPessoas()
}

func ListarPessoa(id *int64) (res *modelApresentacao.ReqGetPessoa,err error) {
	db := database.Conectar()
	defer db.Close()

	pessoasRepo := pessoas.NovoRepo(db)
	return pessoasRepo.ListarPessoa(id)
}

func ListarTarefasPessoa(id *int64) (res []modelApresentacao.ReqTarefaPessoa, err error) {
	db := database.Conectar()
	defer db.Close()

	pessoasRepo := pessoas.NovoRepo(db)

	dados, err := pessoasRepo.ListarPessoa(id)
	//if len(dados) == 0 {
	if err != nil {
		return res, fmt.Errorf("person does not exist")
	}

	if dados == nil {
		return res, fmt.Errorf("unrecognized error")
	}

	res, err = pessoasRepo.ListarTarefasPessoa(id)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch tasks from a person")
	}
	return
}

func AtualizarPessoa(id *int64, req *modelApresentacao.ReqAtualizarPessoa) (res *modelApresentacao.ReqAtualizarPessoa, err error) {
	db := database.Conectar()
	defer db.Close()

	pessoasRepo := pessoas.NovoRepo(db)

	dados, err := pessoasRepo.ListarPessoa(id)
	//if len(dados) == 0 {
	if err != nil {
		return res, fmt.Errorf("unable to update: Person does not exist")
	}

	if dados == nil {
		return res, fmt.Errorf("unrecognized error")
	}

	res, err = pessoasRepo.AtualizarPessoa(id, req)
	if err != nil {
		return nil, fmt.Errorf("unable to update: Team does not exist")
	}
	return
}

func DeletarPessoa(id *int64) (err error) {
	db := database.Conectar()
	defer db.Close()

	pessoasRepo := pessoas.NovoRepo(db)

	dados, err := pessoasRepo.ListarPessoa(id)
	//if len(dados) == 0 {
	if err != nil {
		return fmt.Errorf("person does not exist")
	}

	if dados == nil {
		return fmt.Errorf("unrecognized error")
	}
	err = pessoasRepo.DeletarPessoa(id)
	return
}

func ListarPessoasFiltro(params *utils.RequestParams) (res *modelApresentacao.ListarGetPessoa, err error) {
	db := database.Conectar()
	defer db.Close()
	pessoasRepo := pessoas.NovoRepo(db)

	res, err = pessoasRepo.ListarPessoasFiltro(params)
	if err != nil {
		return nil, fmt.Errorf("usuarios nao listados // " + err.Error())
	}

	return
}