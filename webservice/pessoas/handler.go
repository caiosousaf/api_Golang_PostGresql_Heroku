package pessoas

import (
	"gerenciadorDeProjetos/domain/pessoas"
	modelApresentacao "gerenciadorDeProjetos/domain/pessoas/model"
	errorstratment "gerenciadorDeProjetos/utils/errors-tratment"
	utils "gerenciadorDeProjetos/utils/params"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security bearerAuth
// @Summary POST a new Person
// @Description POST a new Person. For the request to be met, the "nome_pessoa", "funcao_pessoa", "equipe_id" are required.
// @Param		Person		body	string		true	"New Person"
// @Accept json
// @Produce json
// @Success 201 {object} modelApresentacao.ReqPessoa "OK"
// @Failure 400,401 {array} errorstratment.ResError
// @Tags People
// @Router /pessoas [post]
func novaPessoa(c *gin.Context) {

	req := modelApresentacao.ReqPessoa{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorstratment.KeyError(err.Error(),
			"Could not create. Parameters were not passed correctly", 400))
		return
	}
	if res, err := pessoas.NovaPessoa(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorstratment.KeyError(err.Error(), "Team does not exist", 400))
	} else {
		c.JSON(http.StatusCreated, res)
	}

}

// @Security bearerAuth
// Get People
// @Summary Get All People
// @Description Returns all registered people and the name of the team they are assigned to
// @Accept json
// @Produce json
// @Success 200 {array} modelApresentacao.ReqGetPessoa "OK"
// @Failure 404,401 {array} errorstratment.ResError
// @Tags People
// @Router /pessoas [get]
func listarPessoas(c *gin.Context) {

	if pessoas, err := pessoas.ListarPessoas(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, pessoas)
	}
}

// @Security bearerAuth
// @Summary GET a specific Person
// @Description GET a specific person
// @Param		id		path	int		true		"Pessoa_ID"
// @Accept json
// @Produce json
// @Success 200 {array} modelApresentacao.ReqGetPessoa "OK"
// @Failure 401,404 {array} errorstratment.ResError
// @Tags People
// @Router /pessoas/{id} [get]
func listarPessoa(c *gin.Context) {
	id := c.Param("id")

	pessoas, err := pessoas.ListarPessoa(id)
	if err != nil {
		c.JSON(http.StatusNotFound, errorstratment.KeyError(err.Error(), "Person does not exist", 404))
	} else {
		c.JSON(http.StatusOK, pessoas)
	}
}

// @Security bearerAuth
// @Summary GET All Tasks of a specific Person
// @Description GET the tasks registered and assigned to a specific person
// @Param		id		path	int		true		"Pessoa_ID"
// @Accept json
// @Produce json
// @Success 200 {array} modelApresentacao.ReqTarefaPessoa "OK"
// @Failure 401,404 {array} errorstratment.ResError
// @Tags People
// @Router /pessoas/{id}/tasks [get]
func listarTarefasPessoa(c *gin.Context) {
	id := c.Param("id")

	if pessoas, err := pessoas.ListarTarefasPessoa(id); err != nil {
		c.JSON(http.StatusNotFound, errorstratment.KeyError(err.Error(), "Person does not exist", 404))
	} else if len(pessoas) == 0 {
		c.JSON(http.StatusNoContent, pessoas)
	} else {
		c.JSON(http.StatusOK, pessoas)
	}
}

// @Security bearerAuth
// @Summary PUT Person with ID
// @Description PUT a specific person. For the request to be met, the "nome_pessoa" and "funcao_pessoa" and "equipe_id" are required
// @Param        id   				path      	int  	true  	"Pessoa ID"
// @Param		Pessoa				body		string 	true 	"Pessoa"
// @Accept json
// @Produce json
// @Success 200 {object} modelApresentacao.ReqAtualizarPessoa
// @Failure 401,400	{array} errorstratment.ResError
// @Tags People
// @Router /pessoas/{id} [put]
func atualizarPessoa(c *gin.Context) {
	id := c.Param("id")

	req := modelApresentacao.ReqAtualizarPessoa{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not update. Parameters were not passed correctly.", "err": err.Error(),
		})
		return
	}

	if res, err := pessoas.AtualizarPessoa(id, &req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// @Security bearerAuth
// @Summary DELETE a Person
// @Description DELETE a person
// @Param		id		path	int		true		"Pessoa_ID"
// @Accept json
// @Produce json
// @Success 200 {array} errorstratment.ResOk "OK"
// @Failure 401,400 {array} errorstratment.ResError
// @Tags People
// @Router /pessoas/{id} [delete]
func deletarPessoa(c *gin.Context) {
	id := c.Param("id")

	if err := pessoas.DeletarPessoa(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, errorstratment.KeyOk("Person deleted successfully", 200))
	}
}

// @Security bearerAuth
// @Summary GET all people with sort 
// @Description GET all people with sort orderBy & || order (desc, cresc)
// @Param		orderBy		query			string				false		"column" 		Enums(id_pessoa,nome_pessoa, funcao_pessoa, data_contratacao, equipe_id)
// @Param		order		query			string				false		"order desc"	
// @Accept json
// @Produce json
// @Success 200 {array} modelApresentacao.ReqGetPessoa "OK"
// @Failure 401,400 {array} errorstratment.ResError
// @Tags People
// @Router /pessoas/filtros [get]
func listarPessoasFiltro(c *gin.Context) {
	params := utils.ParseParams(c)

	res, err := pessoas.ListarPessoasFiltro(params)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}