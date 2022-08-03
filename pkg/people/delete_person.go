package pessoas

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// @Summary DELETE a Person
// @Description DELETE a person
// @Param		id		path	int		true		"Pessoa_ID"
// @Accept json
// @Produce json
// @Success 200 {array} models.Pessoa
// @Failure 400,404 {string} string "error"
// @Tags People
// @Router /pessoas/{id} [delete]
func (h handler) DeletePerson(c *gin.Context) {
	id := c.Param("id")

	var pessoa models.Pessoa

	if result := h.DB.First(&pessoa, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&pessoa)

	c.Status(http.StatusOK)
}