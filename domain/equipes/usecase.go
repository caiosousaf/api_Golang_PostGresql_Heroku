package equipe

import (
	"fmt"
	"gerenciadorDeProjetos/config/database"
	modelApresentacao "gerenciadorDeProjetos/domain/equipes/model"
	modelPessoa "gerenciadorDeProjetos/domain/pessoas/model"
	"gerenciadorDeProjetos/infra/equipes"
)

func NovaEquipe(req *modelApresentacao.ReqEquipe)(*modelApresentacao.ReqEquipe, error) {
	db := database.Conectar()
	defer db.Close()
	equipesRepo := equipes.NovoRepo(db)

	str := *req.Nome_Equipe

	req.Nome_Equipe = &str

	return equipesRepo.NovaEquipe(req)
}

func ListarEquipes() ([]modelApresentacao.ReqEquipe, error){
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)
	return equipesRepo.ListarEquipes()
}

func BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error) {
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)
	return equipesRepo.BuscarEquipe(id)
}

func BuscarMembrosDeEquipe(id string) (res []modelPessoa.ReqMembros,err error) {
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)

	dados,err := equipesRepo.BuscarEquipe(id)
	//if len(dados) == 0 {
	if err != nil {
		return res, fmt.Errorf("equipe não Encontrada")
	}

	if dados == nil {
		return res, fmt.Errorf("em Aberto")
	}
	res, err = equipesRepo.BuscarMembrosDeEquipe(id)
	if err != nil {
		return nil, fmt.Errorf("não foi possivel buscar os membros")
	}
	return 
}

func BuscarProjetosDeEquipe(id string) (res []modelApresentacao.ReqEquipeProjetos, err error) {
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)

	dados,err := equipesRepo.BuscarEquipe(id)
	//if len(dados) == 0 {
	if err != nil {
		return res, fmt.Errorf("equipe não Encontrada")
	}

	if dados == nil {
		return res, fmt.Errorf("em Aberto")
	}
	res, err = equipesRepo.BuscarProjetosDeEquipe(id)
	if err != nil {
		return nil, fmt.Errorf("não foi possivel buscar projetos de uma equipe")
	}
	return
}

func BuscarTasksDeEquipe(id string) (res []modelApresentacao.ReqTasksbyTeam, err error) {
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)

	dados,err := equipesRepo.BuscarEquipe(id)
	//if len(dados) == 0 {
	if err != nil {
		return res, fmt.Errorf("equipe não Encontrada")
	}

	if dados == nil {
		return res, fmt.Errorf("em Aberto")
	}
	res, err = equipesRepo.BuscarTasksDeEquipe(id)
	if err != nil {
		return nil, fmt.Errorf("não foi possivel buscar tasks de uma equipe")
	}
	return
}

func DeletarEquipe(id string) (err error){
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)

	dados,err := equipesRepo.BuscarEquipe(id)
	//if len(dados) == 0 {
	if err != nil {
		return fmt.Errorf("equipe não Encontrada")
	}

	if dados == nil {
		return fmt.Errorf("em Aberto")
	}
	err = equipesRepo.DeletarEquipe(id)
	return
}

func AtualizarEquipe(id string, req *modelApresentacao.ReqEquipe) (res *modelApresentacao.ReqEquipe, err error){
	db := database.Conectar()
	defer db.Close()
	equipesRepo := equipes.NovoRepo(db)

	dados,err := equipesRepo.BuscarEquipe(id)
	//if len(dados) == 0 {
	if err != nil {
		return nil, fmt.Errorf("equipe não Encontrada")
	}

	if dados == nil {
		return nil, fmt.Errorf("em Aberto")
	}
	res, err = equipesRepo.AtualizarEquipe(id, req)
	if err != nil {
		return nil, fmt.Errorf("não foi possivel atualizar equipe")
	}
	return
}

