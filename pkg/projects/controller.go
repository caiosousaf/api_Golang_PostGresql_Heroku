package projetos

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/server/middlewares"
	
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r.Use(middlewares.CORSMiddleware())

	routes := r.Group("/projetos", middlewares.Auth())
	routes.POST("/", h.AddProject)
	routes.GET("/", h.GetProjects)
	routes.GET("/:id", h.GetProject)
	routes.GET("/status/:status", h.GetStatusProjects)
	routes.GET("/:id/tasks", h.GetProjectTasks)
	routes.PUT("/:id", h.UpdateProject)
	routes.PUT("/:id/status", h.UpdateStatusProject)
	routes.DELETE("/:id", h.DeleteProject)

}