package users

import (
	"fmt"
	"gerenciadorDeProjetos/domain/users"
	modelApresentacao "gerenciadorDeProjetos/domain/users/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security bearerAuth
// @Summary POST a new User
// @Description POST a new User. For the request to be met, the "nome", "email", "password", are required. 
// @Param		NewUser		body	string		true	"NewUser"
// @Accept json
// @Produce json
// @Success 201 {object} modelApresentacao.ReqUser "OK"
// @Failure 401,400 {array} errorstratment.ResError
// @Tags Users
// @Router /users [post]
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