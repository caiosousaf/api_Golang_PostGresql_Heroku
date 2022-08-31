package projetos

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

	routes := r.Group("/projetos", middlewares.Auth())
	routes.POST("/", h.AddProjeto)
	routes.GET("/", h.GetProjetos)
	routes.GET("/:id", h.GetProjeto)
	routes.PUT("/:id", h.UpdateProjeto)
	routes.PUT("/:id/status", h.UpdateStatus)
	routes.DELETE("/:id", h.DeleteProjeto)
}