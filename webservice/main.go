package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Brun0Nasc/sys-projetos/webservice/equipes"
)

func main() {
	r := gin.Default()

	eq := r.Group("equipes")

	equipes.Router(eq)

	r.Run("localhost:3030")
}