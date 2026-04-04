// GetPagosByEstadoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/pagos/application"
)

type GetPagosByEstadoController struct {
	getByEstadoUseCase *application.GetPagosByEstadoUseCase
}

func NewGetPagosByEstadoController(getByEstadoUseCase *application.GetPagosByEstadoUseCase) *GetPagosByEstadoController {
	return &GetPagosByEstadoController{
		getByEstadoUseCase: getByEstadoUseCase,
	}
}

func (ctrl *GetPagosByEstadoController) Run(c *gin.Context) {
	estado := c.Query("estado")
	if estado == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Estado es requerido",
		})
		return
	}

	pagos, err := ctrl.getByEstadoUseCase.Run(estado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los pagos por estado",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pagos)
}