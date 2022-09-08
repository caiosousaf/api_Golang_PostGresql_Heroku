package main

import (
	"gerenciadorDeProjetos/webservice/tasks"
	"gerenciadorDeProjetos/webservice/equipes"
	"gerenciadorDeProjetos/webservice/pessoas"
	"gerenciadorDeProjetos/webservice/projetos"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eq := r.Group("equipes")
	pe := r.Group("pessoas")
	pr := r.Group("projetos")
	tk := r.Group("tasks")

	equipes.Router(eq)
	pessoas.Router(pe)
	projetos.Router(pr)
	tasks.Router(tk)

	r.Run("localhost:3030")
}