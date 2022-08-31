package equipe

import (
	"github.com/Brun0Nasc/sys-projetos/config/database"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"
	"github.com/Brun0Nasc/sys-projetos/infra/equipes"
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