package main

import (
	"gerenciadorDeProjetos/config/server/middlewares"
	"gerenciadorDeProjetos/webservice/equipes"
	"gerenciadorDeProjetos/webservice/login"
	"gerenciadorDeProjetos/webservice/pessoas"
	"gerenciadorDeProjetos/webservice/projetos"
	"gerenciadorDeProjetos/webservice/tasks"
	"gerenciadorDeProjetos/webservice/users"
	"gerenciadorDeProjetos/docs"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @contact.name   Caio Sousa
// @contact.url    http://www.swagger.io/support
// @contact.email  caiosousafernandesferreira@hotmail.com

// @license.name  Mozilla Public License 2.0
// @license.url   https://www.mozilla.org/en-US/MPL/2.0/
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {

	docs.SwaggerInfo.Title = "Gerenciador de Projetos"
	docs.SwaggerInfo.Description = "REST API Desafio"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "golang-posgre-brisanet.herokuapp.com"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"https"}
	
	port := os.Getenv("PORT")
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	equipes.Router(r.Group("equipes", middlewares.Auth()))
	pessoas.Router(r.Group("pessoas", middlewares.Auth()) )
	projetos.Router(r.Group("projetos", middlewares.Auth()))
	tasks.Router(r.Group("tasks", middlewares.Auth()))
	users.Router(r.Group("users", middlewares.Auth()))
	login.Router(r.Group("login"))

	r.Run(":" + port)
	//export PATH=$(go env GOPATH)/bin:$PATH
	//r.Run()
}