package projetos

import (
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

func ListarProjetosComStatus(status string) ([]modelApresentacao.ReqStatusProjeto, error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)
	
	return projetosRepo.ListarProjetosComStatus(status)
}

func DeletarProjeto(id string) error {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	return projetosRepo.DeletarProjeto(id)
}