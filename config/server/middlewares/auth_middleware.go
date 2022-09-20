package middlewares

import (
	"gerenciadorDeProjetos/config/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
// Function to validate the user so that he has access to the content
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		token := header[len(Bearer_schema):]

		if !services.NewJWTService().ValidateToken(token) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}