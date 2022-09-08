package tasks

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) { 
	r.POST("/", NovaTask)
	r.GET("/", ListarTasks)
	r.GET("/:id", ListarTask)
	r.GET("/status/:status", ListarStatusTasks)
}