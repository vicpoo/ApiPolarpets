// DeleteNivelController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/niveles/application"
)

type DeleteNivelController struct {
	deleteUseCase *application.DeleteNivelUseCase
}

func NewDeleteNivelController(deleteUseCase *application.DeleteNivelUseCase) *DeleteNivelController {
	return &DeleteNivelController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteNivelController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar el nivel",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Nivel eliminado exitosamente",
	})
}