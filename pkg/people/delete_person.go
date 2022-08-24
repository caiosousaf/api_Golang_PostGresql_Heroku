package pessoas

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// @Security bearerAuth
// @Summary DELETE a Person
// @Description DELETE a person
// @Param		id		path	int		true		"Pessoa_ID"
// @Accept json
// @Produce json
// @Success 200 {array} models.Pessoa
// @Failure 400 {array} models.Error400Delete
// @Failure 404 {array} models.Error404Delete
// @Tags People
// @Router /pessoas/{id} [delete]
func (h handler) DeletePerson(c *gin.Context) {
	id := c.Param("id")

	var pessoa models.Pessoa

	// Search the database if the person with the selected id exists

	// if it exists then delete it

	if result := h.DB.First(&pessoa, id); result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Unable to delete. non-existent ID " + result.Error.Error(),
		})
		return
	}

	h.DB.Delete(&pessoa)
	c.Status(http.StatusOK)

}
