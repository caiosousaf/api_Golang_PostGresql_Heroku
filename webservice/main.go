package main

import (
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

	equipes.Router(eq)
	pessoas.Router(pe)
	projetos.Router(pr)

	r.Run()
}