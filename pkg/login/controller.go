package login

import (
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

	routes := r.Group("/login")
	routes.POST("/", h.Login)
}

