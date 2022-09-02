package main

import (
	"gerenciadorDeProjetos/webservice/equipes"
	"gerenciadorDeProjetos/webservice/pessoas"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eq := r.Group("equipes")
	pe := r.Group("pessoas")

	equipes.Router(eq)
	pessoas.Router(pe)

	r.Run()
}