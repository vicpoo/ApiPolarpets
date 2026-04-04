// DeletePagoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/pagos/application"
)

type DeletePagoController struct {
	deleteUseCase *application.DeletePagoUseCase
}

func NewDeletePagoController(deleteUseCase *application.DeletePagoUseCase) *DeletePagoController {
	return &DeletePagoController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeletePagoController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el pago",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Pago eliminado exitosamente",
	})
}