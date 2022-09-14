package main

import (
	"gerenciadorDeProjetos/config/server/middlewares"
	"gerenciadorDeProjetos/webservice/equipes"
	"gerenciadorDeProjetos/webservice/login"
	"gerenciadorDeProjetos/webservice/pessoas"
	"gerenciadorDeProjetos/webservice/projetos"
	"gerenciadorDeProjetos/webservice/tasks"
	"gerenciadorDeProjetos/webservice/users"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.Use(cors.Default())

	eq := r.Group("equipes", middlewares.Auth())
	pe := r.Group("pessoas", middlewares.Auth())
	pr := r.Group("projetos", middlewares.Auth())
	tk := r.Group("tasks", middlewares.Auth())
	us := r.Group("users", middlewares.Auth())
	lo := r.Group("login")

	equipes.Router(eq)
	pessoas.Router(pe)
	projetos.Router(pr)
	tasks.Router(tk)
	users.Router(us)
	login.Router(lo)

	r.Run(":" + port)
}