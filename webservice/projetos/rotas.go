package projetos

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/", NovoProjeto)
	r.GET("/", ListarProjetos)
	r.GET("/:id", ListarProjeto)
	r.GET("/status/:status", ListarProjetosComStatus)
}