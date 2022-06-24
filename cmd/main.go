package main

import (
    "os"
	"github.com/caiosousaf/api_desafio_BrisaNet/pkg/people"
    "github.com/caiosousaf/api_desafio_BrisaNet/pkg/teams"
    "github.com/caiosousaf/api_desafio_BrisaNet/pkg/projects"
    "github.com/caiosousaf/api_desafio_BrisaNet/pkg/tasks"
	"time"
	"github.com/caiosousaf/api_desafio_BrisaNet/pkg/common/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
    "github.com/itsjamie/gin-cors"
)

func main() {
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()

    port := os.Getenv("PORT")
    dbUrl := viper.Get("DB_URL").(string)

    r := gin.Default()
    h := db.Init(dbUrl)

    
    pessoas.RegisterRoutes(r, h)
    projetos.RegisterRoutes(r, h)
    equipes.RegisterRoutes(r, h)
    tasks.RegisterRoutes(r, h)
    // register more routes here

    router := gin.New()

// Apply the middleware to the router (works with groups too)
router.Use(cors.Middleware(cors.Config{
	Origins:        "*",
	Methods:        "GET, PUT, POST, DELETE",
	RequestHeaders: "Origin, Authorization, Content-Type",
	ExposedHeaders: "Access-Control-Allow-Origin",
	MaxAge: 50 * time.Second,
	Credentials: false,
	ValidateHeaders: false,
}))

    
    r.Run(":"+port)
}