package equipe

import (
	"gerenciadorDeProjetos/config/database"
	modelApresentacao "gerenciadorDeProjetos/domain/equipes/model"
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

func BuscarMembrosDeEquipe(id string) ([]modelApresentacao.ReqEquipeMembros, error) {
	db := database.Conectar()
	defer db.Close()

	equipesRepo := equipes.NovoRepo(db)
	return equipesRepo.BuscarMembrosDeEquipe(id)
}
