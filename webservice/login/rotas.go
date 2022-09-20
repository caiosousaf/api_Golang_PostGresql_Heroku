package login

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/", Login)
}