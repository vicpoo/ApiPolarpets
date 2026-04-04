// GetPagoByReferenciaExternaController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/pagos/application"
)

type GetPagoByReferenciaExternaController struct {
	getByReferenciaExternaUseCase *application.GetPagoByReferenciaExternaUseCase
}

func NewGetPagoByReferenciaExternaController(getByReferenciaExternaUseCase *application.GetPagoByReferenciaExternaUseCase) *GetPagoByReferenciaExternaController {
	return &GetPagoByReferenciaExternaController{
		getByReferenciaExternaUseCase: getByReferenciaExternaUseCase,
	}
}

func (ctrl *GetPagoByReferenciaExternaController) Run(c *gin.Context) {
	referencia := c.Query("referencia")
	if referencia == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Referencia externa es requerida",
		})
		return
	}

	pago, err := ctrl.getByReferenciaExternaUseCase.Run(referencia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el pago",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pago)
}