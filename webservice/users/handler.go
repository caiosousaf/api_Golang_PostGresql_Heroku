package users

import (
	"fmt"
	"gerenciadorDeProjetos/domain/users"
	modelApresentacao "gerenciadorDeProjetos/domain/users/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NovoUsuario(c *gin.Context) {
	fmt.Println("Tentando cadastrar usuário")
	req := modelApresentacao.ReqUser{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not create. Parameters were not passed correctly ", "error": err.Error(),
		})
		return
	}

	if res, err := users.NovoUsuario(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, res)
	}
}