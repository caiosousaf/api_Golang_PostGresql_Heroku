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
		return nil, fmt.Errorf("project not found in database")
	}

	if dados == nil {
		return nil, fmt.Errorf("unrecognized error")
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
		return fmt.Errorf("project not found in database")
	}

	if dados == nil {
		return fmt.Errorf("unrecognized error")
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
		return res, fmt.Errorf("project not found in database")
	}

	if dados == nil {
		return res, fmt.Errorf("unrecognized error")
	}

	res, err = projetosRepo.AtualizarProjeto(id, req)
	if err != nil {
		return nil, fmt.Errorf("unable to update: Team does not exist")
	}
	return
}

func AtualizarStatusProjeto(id string, req *modelApresentacao.ReqAtualizarProjeto) (res *modelApresentacao.ReqAtualizarProjeto, err error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	res, err = projetosRepo.AtualizarStatusProjeto(id, req)
	if err != nil {
		return nil, fmt.Errorf("unable to update status: Project Does not exist")
	}
	return
}
