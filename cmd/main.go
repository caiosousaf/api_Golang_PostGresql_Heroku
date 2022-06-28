package main

import (
	"os"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/people"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/projects"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/tasks"
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/teams"
    
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/db"
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

    r.Use(cors.Default())
    

    pessoas.RegisterRoutes(r, h)
    projetos.RegisterRoutes(r, h)
    equipes.RegisterRoutes(r, h)
    tasks.RegisterRoutes(r, h)
    // register more routes here

    r.Run(":"+port)
}