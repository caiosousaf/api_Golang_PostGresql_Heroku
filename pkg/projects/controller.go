package projetos

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

	routes := r.Group("/projetos")
	routes.POST("/", h.AddProject)
	routes.GET("/", h.GetProjects)
	routes.GET("/:id", h.GetProject)
	routes.GET("/:status", h.GetStatusProjects)
	routes.GET("/:id/tasks", h.GetProjectTasks)
	routes.PUT("/:id", h.UpdateProject)
	routes.DELETE("/:id", h.DeleteProject)
	
}