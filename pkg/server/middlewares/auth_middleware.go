package middlewares

import (
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/services"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Accept-Version, Content-MD5, Date, X-api-Version")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

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
			c.AbortWithStatus(401)
		}

		token := header[len(Bearer_schema):]

		if !services.NewJWTService().ValidateToken(token) {
			c.AbortWithStatus(401)
		}
	}
}