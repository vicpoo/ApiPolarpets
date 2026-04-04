// UpdatePagoEstadoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/pagos/application"
)

type UpdatePagoEstadoController struct {
	updateEstadoUseCase *application.UpdatePagoEstadoUseCase
}

func NewUpdatePagoEstadoController(updateEstadoUseCase *application.UpdatePagoEstadoUseCase) *UpdatePagoEstadoController {
	return &UpdatePagoEstadoController{
		updateEstadoUseCase: updateEstadoUseCase,
	}
}

func (ctrl *UpdatePagoEstadoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var request struct {
		Estado string `json:"estado"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	err = ctrl.updateEstadoUseCase.Run(int32(id), request.Estado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el estado del pago",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Estado del pago actualizado exitosamente",
	})
}