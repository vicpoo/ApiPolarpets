// GetPagosByMetodoPagoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/pagos/application"
)

type GetPagosByMetodoPagoController struct {
	getByMetodoPagoUseCase *application.GetPagosByMetodoPagoUseCase
}

func NewGetPagosByMetodoPagoController(getByMetodoPagoUseCase *application.GetPagosByMetodoPagoUseCase) *GetPagosByMetodoPagoController {
	return &GetPagosByMetodoPagoController{
		getByMetodoPagoUseCase: getByMetodoPagoUseCase,
	}
}

func (ctrl *GetPagosByMetodoPagoController) Run(c *gin.Context) {
	metodoPago := c.Query("metodo_pago")
	if metodoPago == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Método de pago es requerido",
		})
		return
	}

	pagos, err := ctrl.getByMetodoPagoUseCase.Run(metodoPago)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los pagos por método de pago",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pagos)
}