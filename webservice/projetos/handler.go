package projetos

import (
	"fmt"


	"gerenciadorDeProjetos/domain/projetos"
	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"

	"github.com/gin-gonic/gin"
)

func NovoProjeto(c *gin.Context) {
	fmt.Println("Tentando cadastrar um novo projeto")
	req := modelApresentacao.ReqProjeto{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly " , "error": err.Error(),
		})
		return
	}

	projetos.NovoProjeto(&req, c)
	c.JSON(201, gin.H{"OK": "Projeto cadastrado com sucesso"})
}