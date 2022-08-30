package pessoas

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

	routes := r.Group("/pessoas", middlewares.Auth())
	routes.POST("/", h.AddPerson)
	routes.GET("/", h.GetPeople)
	routes.GET("/filtros/", h.GetPersonName)
	routes.GET("/:id", h.GetPerson)
	routes.GET("/:id/tasks", h.GetTaskPerson)
	routes.PUT("/:id", h.UpdatePerson)
	routes.DELETE("/:id", h.DeletePerson)
}