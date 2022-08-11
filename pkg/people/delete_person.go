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
// @Failure 400 {array} models.Error400Delete
// @Failure 404 {array} models.Error404Delete
// @Tags People
// @Router /pessoas/{id} [delete]
func (h handler) DeletePerson(c *gin.Context) {
	id := c.Param("id")
	var IdExist	int
	var pessoa models.Pessoa

	// Search the database if the person with the selected id exists
	if result := h.DB.Raw("select count(*) from pessoas where id_pessoa = ?", id).Scan(&IdExist); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	// if it exists then delete it
	if IdExist == 1 {
		if result := h.DB.First(&pessoa, id); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
	
		h.DB.Delete(&pessoa)
		c.Status(http.StatusOK)
	} else {
		c.JSON(400, gin.H{
			"message": "Unable to delete. non-existent ID " ,
		})
		return
	}

}