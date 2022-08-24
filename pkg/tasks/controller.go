package tasks

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

	routes := r.Group("/tasks", middlewares.Auth())
	routes.POST("/", h.AddTask)
	routes.GET("/status/:status", h.GetStatusTasks)
	routes.GET("/", h.GetTasks)
	routes.GET("/:id", h.GetTask)
	routes.GET("/list", h.GetStatusList)
	routes.PUT("/:id", h.UpdateTask)
	routes.PUT("/:id/status", h.UpdateStatusTask)
	routes.DELETE("/:id", h.DeleteTask)
}