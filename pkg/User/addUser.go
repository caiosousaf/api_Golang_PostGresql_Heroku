package user

import (
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/services"
	"github.com/gin-gonic/gin"
)

func (h handler) CreateUser(c *gin.Context) {
	var p models.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	p.Password = services.SHAR256Encoder(p.Password)



	if result := h.DB.Create(&p); result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create user: " + err.Error(),
		})
		return
	}

	c.Status(204)
}