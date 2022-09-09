package projetos

import (
	"fmt"
	"gerenciadorDeProjetos/config/database"
	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"
	"gerenciadorDeProjetos/infra/projetos"
)

func NovoProjeto(req *modelApresentacao.ReqProjeto) (*modelApresentacao.ReqProjetos, error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	return projetosRepo.NovoProjeto(req)
}

func ListarProjetos() ([]modelApresentacao.ReqProjetos, error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	return projetosRepo.ListarProjetos()
}

func ListarProjeto(id string) (*modelApresentacao.ReqProjetos, error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	return projetosRepo.ListarProjeto(id)
}

func ListarTasksProjeto(id string) (res []modelApresentacao.ReqTasksProjeto, err error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	dados, err := projetosRepo.ListarProjeto(id)

	if err != nil {
		return nil, fmt.Errorf("projeto não Encontrado")
	}

	if dados == nil {
		return nil, fmt.Errorf("em Aberto")
	}
	res, err = projetosRepo.ListarTasksProjeto(id)
	if err != nil {
		return nil, err
	}
	return
}

func ListarProjetosComStatus(status string) ([]modelApresentacao.ReqStatusProjeto, error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	return projetosRepo.ListarProjetosComStatus(status)
}

func DeletarProjeto(id string) (err error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	dados, err := projetosRepo.ListarProjeto(id)

	if err != nil {
		return fmt.Errorf("projeto não Encontrado")
	}

	if dados == nil {
		return fmt.Errorf("em Aberto")
	}
	err = projetosRepo.DeletarProjeto(id)
	return
}

func AtualizarProjeto(id string, req *modelApresentacao.ReqAtualizarProjeto) (res *modelApresentacao.ReqAtualizarProjeto, err error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)
	dados, err := projetosRepo.ListarProjeto(id)

	if err != nil {
		return res, fmt.Errorf("projeto não Encontrado")
	}

	if dados == nil {
		return res, fmt.Errorf("em Aberto")
	}

	res, err = projetosRepo.AtualizarProjeto(id, req)
	if err != nil {
		return nil, fmt.Errorf("não foi possivel atualizar: Projeto Não existe")
	}
	return
}

func AtualizarStatusProjeto(id string, req *modelApresentacao.ReqAtualizarProjeto) (res *modelApresentacao.ReqAtualizarProjeto, err error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)
	dados, err := projetosRepo.ListarProjeto(id)

	if err != nil {
		return res, fmt.Errorf("projeto não Encontrado")
	}

	if dados == nil {
		return res, fmt.Errorf("em Aberto")
	}

	res, err = projetosRepo.AtualizarStatusProjeto(id, req)
	if err != nil {
		return nil, fmt.Errorf("não foi possivel atualizar status: Projeto Não existe")
	}
	return
}
