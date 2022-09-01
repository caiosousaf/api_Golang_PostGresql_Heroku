package pessoas

import (
	"gerenciadorDeProjetos/config/database"
	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"

	"github.com/gin-gonic/gin"
)

func novaPessoa(req *modelApresentacao.ReqPessoa, c *gin.Context) {
	db := database.Conectar()
	defer db.Close()
	
}