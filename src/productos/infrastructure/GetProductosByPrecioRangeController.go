// GetProductosByPrecioRangeController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/productos/application"
)

type GetProductosByPrecioRangeController struct {
	getByPrecioRangeUseCase *application.GetProductosByPrecioRangeUseCase
}

func NewGetProductosByPrecioRangeController(getByPrecioRangeUseCase *application.GetProductosByPrecioRangeUseCase) *GetProductosByPrecioRangeController {
	return &GetProductosByPrecioRangeController{
		getByPrecioRangeUseCase: getByPrecioRangeUseCase,
	}
}

func (ctrl *GetProductosByPrecioRangeController) Run(c *gin.Context) {
	minPrecioStr := c.Query("min_precio")
	maxPrecioStr := c.Query("max_precio")

	if minPrecioStr == "" || maxPrecioStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "min_precio y max_precio son requeridos",
		})
		return
	}

	minPrecio, err := strconv.ParseFloat(minPrecioStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "min_precio inválido",
			"error":   err.Error(),
		})
		return
	}

	maxPrecio, err := strconv.ParseFloat(maxPrecioStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "max_precio inválido",
			"error":   err.Error(),
		})
		return
	}

	productos, err := ctrl.getByPrecioRangeUseCase.Run(minPrecio, maxPrecio)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los productos por rango de precio",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, productos)
}