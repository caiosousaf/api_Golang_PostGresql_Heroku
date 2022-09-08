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
	
	str := *req.Nome_Pessoa

	req.Nome_Pessoa = &str

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

func ListarTarefasPessoa(id string) ([]modelApresentacao.ReqTarefaPessoa, error) {
	db := database.Conectar()
	defer db.Close()

	pessoasRepo := pessoas.NovoRepo(db)
	return pessoasRepo.ListarTarefasPessoa(id)
}

func AtualizarPessoa(id string, req *modelApresentacao.ReqAtualizarPessoa) (*modelApresentacao.ReqAtualizarPessoa, error) {
	db := database.Conectar()
	defer db.Close()

	pessoasRepo := pessoas.NovoRepo(db)

	str := *req.Nome_Pessoa

	req.Nome_Pessoa = &str

	return pessoasRepo.AtualizarPessoa(id, req)
}

func DeletarPessoa(id string)(err error) {
	db := database.Conectar()
	defer db.Close()
	
	pessoasRepo := pessoas.NovoRepo(db)
	
	dados,err := pessoasRepo.ListarPessoa(id)
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