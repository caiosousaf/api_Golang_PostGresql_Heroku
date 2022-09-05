package projetos

import (
	"gerenciadorDeProjetos/config/database"
	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"
	"gerenciadorDeProjetos/infra/projetos"
	"github.com/gin-gonic/gin"
)

func NovoProjeto(req *modelApresentacao.ReqProjeto, c *gin.Context) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	projetosRepo.NovoProjeto(req, c)
}