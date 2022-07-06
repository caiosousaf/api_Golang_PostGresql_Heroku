package pessoas

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

	routes := r.Group("/pessoas")
	routes.POST("/", h.AddPerson)
	routes.POST("/:id/task", h.AddTaskPessoa)
	routes.GET("/", h.GetPeople)
	routes.GET("/:id", h.GetPerson)
	routes.GET("/:id/tasks", h.GetTaskPerson)
	routes.PUT("/:id", h.UpdatePerson)
	routes.DELETE("/:id", h.DeletePerson)
}