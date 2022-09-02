package pessoas

import (
	"gerenciadorDeProjetos/config/database"
	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
	"gerenciadorDeProjetos/infra/pessoas"

	"github.com/gin-gonic/gin"
)

func NovaPessoa(req *modelApresentacao.ReqPessoa, c *gin.Context) {
	db := database.Conectar()
	defer db.Close()
	pessoasRepo := pessoas.NovoRepo(db)
	
	str := *req.Nome_Pessoa

	req.Nome_Pessoa = &str

	pessoasRepo.NovaPessoa(req, c)
}