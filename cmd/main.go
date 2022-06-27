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

    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http:localhost:3000"}
    // config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
    // config.AllowAllOrigins = true

    r.Use(cors.New(config))

    pessoas.RegisterRoutes(r, h)
    projetos.RegisterRoutes(r, h)
    equipes.RegisterRoutes(r, h)
    tasks.RegisterRoutes(r, h)
    // register more routes here

    r.Run(":"+port)
}