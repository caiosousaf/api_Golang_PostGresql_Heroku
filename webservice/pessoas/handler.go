package pessoas

import (
	"fmt"
	"gerenciadorDeProjetos/domain/pessoas"
	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NovaPessoa(c *gin.Context) {
	fmt.Println("Tentando cadastrar uma nova pessoa")
	req := modelApresentacao.ReqPessoa{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly " , "error": err.Error(),
		})
		return
	}

	pessoas.NovaPessoa(&req, c)
	c.JSON(http.StatusCreated, gin.H{"OK": "Pessoa Cadastrada com sucesso"})
}