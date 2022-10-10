package equipe

import (
	"fmt"
	"gerenciadorDeProjetos/config/database"
	modelApresentacao "gerenciadorDeProjetos/domain/equipes/model"
	modelPessoa "gerenciadorDeProjetos/domain/pessoas/model"
	"gerenciadorDeProjetos/infra/equipes"
	utils "gerenciadorDeProjetos/utils/params"
)

func NovaEquipe(req *modelApresentacao.ReqEquipe) (*modelApresentacao.ReqEquipe, error) {
	db := database.Conectar()
	defer db.Close()
	equipesRepo := equipes.NovoRepo(db)

	str := *req.Nome_Equipe

	req.Nome_Equipe = &str

	return equipesRepo.NovaEquipe(req)
}

func ListarEquipes() (res []modelApresentacao.ReqEquipe, err error) {
	db := database.Conectar() //error conection
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)

	res, err = equipesRepo.ListarEquipes()

	for i := range res {

		id := fmt.Sprint(*res[i].ID_Equipe)
		pessoa, err := equipesRepo.BuscarMembrosDeEquipe(id)
		if err != nil {
			return nil, err
		}
		res[i].Pessoas = &pessoa
	}
	return
}

func BuscarEquipe(id string) (res *modelApresentacao.ReqEquipe, err error) {
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)

	res, err = equipesRepo.BuscarEquipe(id)
	if err != nil {
		return nil, fmt.Errorf("team does not exist")
	}

	pessoas, err := equipesRepo.BuscarMembrosDeEquipe(id)
	if err != nil {
		return nil, err
	}
	projetos, err := equipesRepo.BuscarProjetosDeEquipe(id)

	tasks, err := equipesRepo.BuscarTasksDeEquipe(id)

	if err != nil {
		return nil, err
	}

	res.Pessoas = &pessoas
	res.Projetos = &projetos
	res.Tarefas = &tasks
	return
}

func BuscarMembrosDeEquipe(id string) (res []modelPessoa.ReqMembros, err error) {
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)

	dados, err := equipesRepo.BuscarEquipe(id)
	//if len(dados) == 0 {
	if err != nil {
		return res, fmt.Errorf("team does not exist")
	}

	if dados == nil {
		return res, fmt.Errorf("unrecognized error")
	}
	res, err = equipesRepo.BuscarMembrosDeEquipe(id)
	if err != nil {
		return nil, fmt.Errorf("could not find members")
	}
	return
}

func BuscarProjetosDeEquipe(id string) (res []modelApresentacao.ReqEquipeProjetos, err error) {
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)

	dados, err := equipesRepo.BuscarEquipe(id)
	//if len(dados) == 0 {
	if err != nil {
		return res, fmt.Errorf("team does not exist")
	}

	if dados == nil {
		return res, fmt.Errorf("unrecognized error")
	}
	res, err = equipesRepo.BuscarProjetosDeEquipe(id)
	if err != nil {
		return nil, fmt.Errorf("could not find projects of team")
	}
	return
}

func BuscarTasksDeEquipe(id string) (res []modelApresentacao.ReqTasksbyTeam, err error) {
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)

	dados, err := equipesRepo.BuscarEquipe(id)
	//if len(dados) == 0 {
	if err != nil {
		return res, fmt.Errorf("team does not exist")
	}

	if dados == nil {
		return res, fmt.Errorf("unrecognized error")
	}
	res, err = equipesRepo.BuscarTasksDeEquipe(id)
	if err != nil {
		return nil, fmt.Errorf("could not find tasks of team")
	}
	return
}

func DeletarEquipe(id string) (err error) {
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)

	dados, err := equipesRepo.BuscarEquipe(id)
	//if len(dados) == 0 {
	if err != nil {
		return fmt.Errorf("team does not exist")
	}

	if dados == nil {
		return fmt.Errorf("unrecognized error")
	}
	err = equipesRepo.DeletarEquipe(id)
	return
}

func AtualizarEquipe(id string, req *modelApresentacao.ReqEquipe) (res *modelApresentacao.ReqEquipe, err error) {
	db := database.Conectar()
	defer db.Close()
	equipesRepo := equipes.NovoRepo(db)

	dados, err := equipesRepo.BuscarEquipe(id)
	//if len(dados) == 0 {
	if err != nil {
		return nil, fmt.Errorf("team does not exist")
	}

	if dados == nil {
		return nil, fmt.Errorf("unrecognized error")
	}
	res, err = equipesRepo.AtualizarEquipe(id, req)
	if err != nil {
		return nil, fmt.Errorf("unable to update team")
	}
	return
}

func ListarEquipesFiltro(params *utils.RequestParams) (res []modelApresentacao.ReqEquipe, err error) {
	db := database.Conectar()
	defer db.Close()
	equipesRepo := equipes.NovoRepo(db)

	res, err = equipesRepo.ListarEquipesFiltro(params)
	if err != nil {
		return nil, fmt.Errorf("unable to search teams // " + err.Error())
	}

	for i := range res {

		id := fmt.Sprint(*res[i].ID_Equipe)
		pessoa, err := equipesRepo.BuscarMembrosDeEquipe(id)
		if err != nil {
			return nil, err
		}
		res[i].Pessoas = &pessoa
	}
	return
}
