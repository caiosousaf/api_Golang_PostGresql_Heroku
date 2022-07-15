package user

import (
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/services"
	"github.com/gin-gonic/gin"
)

func (h handler) Login(c *gin.Context) {
	

	var p models.Login

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var user models.User
	
	if result := h.DB.Where("email = ?", p.Email).First(&user); result.Error != nil {
		c.JSON(401, gin.H{
			"error": "cannot find user: ",
		})
		return
	}

	if user.Password != services.SHAR256Encoder(p.Password) {
		c.JSON(400, gin.H{
			"error": "Invalid Credentials: ",
		})
		return
	}
	token, err := services.NewJWTService().GenerateToken(user.ID_Usuario)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}