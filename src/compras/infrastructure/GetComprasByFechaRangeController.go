// GetComprasByFechaRangeController.go
package infrastructure

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/compras/application"
)

type GetComprasByFechaRangeController struct {
	getByFechaRangeUseCase *application.GetComprasByFechaRangeUseCase
}

func NewGetComprasByFechaRangeController(getByFechaRangeUseCase *application.GetComprasByFechaRangeUseCase) *GetComprasByFechaRangeController {
	return &GetComprasByFechaRangeController{
		getByFechaRangeUseCase: getByFechaRangeUseCase,
	}
}

func (ctrl *GetComprasByFechaRangeController) Run(c *gin.Context) {
	fechaInicioStr := c.Query("fecha_inicio")
	fechaFinStr := c.Query("fecha_fin")

	if fechaInicioStr == "" || fechaFinStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "fecha_inicio y fecha_fin son requeridos",
		})
		return
	}

	fechaInicio, err := time.Parse("2006-01-02", fechaInicioStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Formato de fecha_inicio inválido. Use YYYY-MM-DD",
			"error":   err.Error(),
		})
		return
	}

	fechaFin, err := time.Parse("2006-01-02", fechaFinStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Formato de fecha_fin inválido. Use YYYY-MM-DD",
			"error":   err.Error(),
		})
		return
	}

	compras, err := ctrl.getByFechaRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las compras en el rango de fechas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, compras)
}