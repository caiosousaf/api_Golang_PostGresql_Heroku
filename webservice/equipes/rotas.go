package equipes

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/", novaEquipe)
	r.GET("/", listarEquipes)
	r.GET("/:id", buscarEquipe)
	r.GET("/:id/membros", buscarMembrosDeEquipe)
}