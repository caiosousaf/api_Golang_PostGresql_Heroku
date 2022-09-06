package equipe

import (
	"gerenciadorDeProjetos/config/database"
	modelApresentacao "gerenciadorDeProjetos/domain/equipes/model"
	modelPessoa "gerenciadorDeProjetos/domain/pessoas/model"
	"gerenciadorDeProjetos/infra/equipes"
	"github.com/gin-gonic/gin"
)

func NovaEquipe(req *modelApresentacao.ReqEquipe, c *gin.Context) {
	db := database.Conectar()
	defer db.Close()
	equipesRepo := equipes.NovoRepo(db)

	str := *req.Nome_Equipe

	req.Nome_Equipe = &str

	equipesRepo.NovaEquipe(req, c)
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

func BuscarMembrosDeEquipe(id string) ([]modelPessoa.ReqMembros, error) {
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)
	return equipesRepo.BuscarMembrosDeEquipe(id)
}

func BuscarProjetosDeEquipe(id string) ([]modelApresentacao.ReqEquipeProjetos, error) {
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)
	return equipesRepo.BuscarProjetosDeEquipe(id)
}

func DeletarEquipe(id string) error{
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)
	return equipesRepo.DeletarEquipe(id)
}

func AtualizarEquipe(id string, req *modelApresentacao.ReqEquipe) (*modelApresentacao.ReqEquipe, error){
	db := database.Conectar()
	defer db.Close()
	equipesRepo := equipes.NovoRepo(db)

	str := *req.Nome_Equipe

	req.Nome_Equipe = &str

	return equipesRepo.AtualizarEquipe(id, req)
}

