// UpdatePagoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/pagos/application"
	"github.com/vicpoo/ApiPolarpets/src/pagos/domain/entities"
)

type UpdatePagoController struct {
	updateUseCase *application.UpdatePagoUseCase
}

func NewUpdatePagoController(updateUseCase *application.UpdatePagoUseCase) *UpdatePagoController {
	return &UpdatePagoController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdatePagoController) Run(c *gin.Context) {
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
		IDUsuario         int32   `json:"id_usuario"`
		Monto             float64 `json:"monto"`
		MetodoPago        string  `json:"metodo_pago"`
		Estado            string  `json:"estado"`
		ReferenciaExterna string  `json:"referencia_externa"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	pago := entities.NewPagos(
		request.IDUsuario,
		request.Monto,
		request.MetodoPago,
		request.Estado,
		request.ReferenciaExterna,
	)
	pago.SetIDPago(int32(id))

	updatedPago, err := ctrl.updateUseCase.Run(pago)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el pago",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedPago)
}