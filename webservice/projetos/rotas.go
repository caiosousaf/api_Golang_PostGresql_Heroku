package projetos

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/", NovoProjeto)
	r.GET("/", ListarProjetos)
	r.GET("/:id", ListarProjeto)
	r.GET("/:id/tasks", ListarTasksProjeto)
	r.GET("/status/:status", ListarProjetosComStatus)
	r.GET("/filtros", listarProjetosFiltro)
	r.DELETE("/:id", DeletarProjeto)
	r.PUT("/:id", AtualizarProjeto)
	r.PUT("/:id/status", AtualizarStatusProjeto)
}