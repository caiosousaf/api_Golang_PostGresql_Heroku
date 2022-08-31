package login

import (
	"net/http"

	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"github.com/Brun0Nasc/sys-projetos/pkg/common/services"
	"github.com/gin-gonic/gin"
)

type MakeLogin struct {
	Email	string	`json:"email"`
	Senha	string	`json:"senha"`
}

func (h handler) Login(c *gin.Context) {
	var p MakeLogin

	if err := c.ShouldBindJSON(&p); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var login models.Login
	login.Email = p.Email
	login.Senha = p.Senha

	var user models.User
	dbError := h.DB.Where("email = ?", login.Email).First(&user).Error
	if dbError != nil {
		c.JSON(400, gin.H{"error":"cannot find user"})
		return
	}

	if user.Senha != services.SHA256Encoder(login.Senha){
		c.JSON(400, gin.H{"error":"invalid cedentials",})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID_User)

	if err != nil{
		c.JSON(500, gin.H{"error":err.Error(),})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}