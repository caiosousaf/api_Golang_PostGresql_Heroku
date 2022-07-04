package equipes

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

	routes := r.Group("/equipes")
	routes.GET("/", h.GetTeams)
	routes.POST("/", h.AddTeam)
	routes.GET("/:id", h.GetTeam)
	routes.GET("/:id/projetos", h.GetTeamProject)
	routes.GET("/:id/pessoas", h.GetTeamMembers)
	routes.GET("/membros", h.GetTeamsMembers)
	routes.PUT("/:id", h.UpdateTeam)
	routes.DELETE("/:id", h.DeleteTeam)
}