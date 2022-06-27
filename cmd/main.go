package main

import (
	"os"
	"github.com/caiosousaf/api_desafio_BrisaNet/pkg/people"
	"github.com/caiosousaf/api_desafio_BrisaNet/pkg/projects"
	"github.com/caiosousaf/api_desafio_BrisaNet/pkg/tasks"
	"github.com/caiosousaf/api_desafio_BrisaNet/pkg/teams"
    
	"github.com/caiosousaf/api_desafio_BrisaNet/pkg/common/db"
	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
    "github.com/gin-contrib/cors"
)

func main() {
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()
    

    port := os.Getenv("PORT")
    dbUrl := viper.Get("DB_URL").(string)

    r := gin.Default()
    h := db.Init(dbUrl)

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"https://foo.com"},
        AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
          return origin == "https://localhost:3000"
        },

      }))

    pessoas.RegisterRoutes(r, h)
    projetos.RegisterRoutes(r, h)
    equipes.RegisterRoutes(r, h)
    tasks.RegisterRoutes(r, h)
    // register more routes here

    r.Run(":"+port)
}