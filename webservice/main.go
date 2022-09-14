package main

import (
	"gerenciadorDeProjetos/webservice/equipes"
	"gerenciadorDeProjetos/webservice/pessoas"
	"gerenciadorDeProjetos/webservice/projetos"
	"gerenciadorDeProjetos/webservice/tasks"
	"gerenciadorDeProjetos/webservice/users"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eq := r.Group("equipes")
	pe := r.Group("pessoas")
	pr := r.Group("projetos")
	tk := r.Group("tasks")
	us := r.Group("users")

	equipes.Router(eq)
	pessoas.Router(pe)
	projetos.Router(pr)
	tasks.Router(tk)
	users.Router(us)

	r.Run()
}