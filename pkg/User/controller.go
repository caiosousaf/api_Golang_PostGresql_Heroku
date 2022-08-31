package user

import (
	"github.com/Brun0Nasc/sys-projetos/pkg/common/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.Use(middlewares.CORSMiddleware())
	
	h := &handler{
		DB: db,
	}
	
	routes := r.Group("/user", middlewares.Auth())
	routes.POST("/", h.AddUser)
}

