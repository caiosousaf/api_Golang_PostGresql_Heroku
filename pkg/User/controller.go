package user

import (
	//"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/server/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	//middlewares.Auth()
	routes := r.Group("/user")
	routes.POST("/", h.CreateUser)
	routes.POST("/login", h.Login)
}