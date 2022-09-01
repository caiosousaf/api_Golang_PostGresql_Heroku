package main

import (
	"github.com/gin-gonic/gin"
	"gerenciadorDeProjetos/webservice/equipes"
)

func main() {
	r := gin.Default()

	eq := r.Group("equipes")

	equipes.Router(eq)

	r.Run()
}