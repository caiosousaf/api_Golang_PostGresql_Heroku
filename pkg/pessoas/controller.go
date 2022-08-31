package pessoas

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

	routes := r.Group("/pessoas", middlewares.Auth())
	routes.POST("/", h.AddPessoa)
	routes.GET("/", h.GetPessoas)
	routes.GET("/:id", h.GetPessoa)
	routes.POST("/:id/tasks", h.AddTaskPessoa)
	routes.PUT("/:id", h.UpdatePessoa)
	routes.PUT("/:id/favoritar", h.FavoritarPessoa)
	routes.DELETE("/:id", h.DeletePessoa)
}