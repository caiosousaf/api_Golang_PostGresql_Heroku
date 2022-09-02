package pessoas

import (
	"database/sql"
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

func ListarPessoas(c *gin.Context) {
	fmt.Println("Tentando Listar todas as pessoas")
	if pessoas, err := pessoas.ListarPessoas(); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message":"Nenhum registro encontrado", "err":err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, pessoas)
	}
}