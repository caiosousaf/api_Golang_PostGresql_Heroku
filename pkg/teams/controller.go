package equipes

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

	//routes := r.Group("/equipes", middlewares.Auth())
	routes := r.Group("/equipes", middlewares.Auth())
	routes.GET("/", h.GetTeams)
	routes.POST("/", h.AddTeam)
	routes.GET("/:id", h.GetTeam)
	routes.GET("/:id/projetos", h.GetTeamProject)
	routes.GET("/:id/pessoas", h.GetTeamMembers)
	routes.PUT("/:id", h.UpdateTeam)
	routes.DELETE("/:id", h.DeleteTeam)
}