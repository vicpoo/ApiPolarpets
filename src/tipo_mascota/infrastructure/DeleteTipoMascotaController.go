// DeleteTipoMascotaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/application"
)

type DeleteTipoMascotaController struct {
	deleteUseCase *application.DeleteTipoMascotaUseCase
}

func NewDeleteTipoMascotaController(deleteUseCase *application.DeleteTipoMascotaUseCase) *DeleteTipoMascotaController {
	return &DeleteTipoMascotaController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteTipoMascotaController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el tipo de mascota",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Tipo de mascota eliminado exitosamente",
	})
}