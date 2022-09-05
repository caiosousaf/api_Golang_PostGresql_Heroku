package pessoas

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/", NovaPessoa)
	r.GET("/", ListarPessoas)
	r.GET("/:id", ListarPessoa)
}