package login

import (
	"gerenciadorDeProjetos/config/services"
	"gerenciadorDeProjetos/domain/login"
	modelApresentacao "gerenciadorDeProjetos/domain/login/model"
	"net/http"
	utils "gerenciadorDeProjetos/utils/errors-tratment"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	req := modelApresentacao.Login{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.KeyError(err.Error(), 
		"Could not login. Parameters were not passed correctly", 400))
		return
	}

	user, err := login.LoginUsuario(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if user == nil && err == nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Credentials: ",
		})
		return
	}

	// Checks if there is an error in this request
	token, err := services.NewJWTService().GenerateToken(user.ID_Usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// If everything is true the token is generated
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
