package pessoas

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/", novaPessoa)
	r.GET("/", listarPessoas)
	r.GET("/:id", listarPessoa)
	r.GET("/:id/tasks", listarTarefasPessoa)
	r.PUT("/:id", atualizarPessoa)
	r.DELETE("/:id", deletarPessoa)
	r.GET("/filtros", listarPessoasFiltro)
}