// DeleteSkinController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/skins/application"
)

type DeleteSkinController struct {
	deleteUseCase *application.DeleteSkinUseCase
}

func NewDeleteSkinController(deleteUseCase *application.DeleteSkinUseCase) *DeleteSkinController {
	return &DeleteSkinController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteSkinController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar la skin",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Skin eliminada exitosamente",
	})
}