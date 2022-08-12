package projetos

import (
	"net/http"

	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// @Summary Delete a specific Project
// @Description DELETE a specific project. For the request to be met, the "id_projeto" are required
// @Param        id   		path      	int  	true  	"Projeto ID"
// @Accept json
// @Produce json
// @Success 200 {array} models.Projeto
// @Failure 400 {array} models.Error400Delete
// @Failure 404 {array} models.Error404Delete
// @Tags Projects
// @Router /projetos/{id} [delete]
func (h handler) DeleteProject(c *gin.Context) {
	id := c.Param("id")
	var IdExist int
	var projeto models.Projeto

	// Search the database if the person with the selected id exists
	if result := h.DB.Raw("select count(*) from projetos where id_projeto = ?", id).Scan(&IdExist); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	// if it exists then delete it
	if IdExist == 1 {
		if result := h.DB.First(&projeto, id); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
	
		h.DB.Delete(&projeto)
	
		c.Status(http.StatusOK)
	} else {
		c.JSON(400, gin.H{
			"message": "Unable to delete. non-existent ID " ,
		})
		return
	}

}