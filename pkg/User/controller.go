package user

import (
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/server/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/gin-contrib/cors"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	r.Use(cors.Default())

	//middlewares.Auth()
	routes := r.Group("/user")
	routes.POST("/", h.CreateUser, middlewares.Auth())
	routes.POST("/login", h.Login)
}