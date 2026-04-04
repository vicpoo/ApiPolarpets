// GetProductoConDetallesController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/productos/application"
)

type GetProductoConDetallesController struct {
	getProductoConDetallesUseCase *application.GetProductoConDetallesUseCase
}

func NewGetProductoConDetallesController(getProductoConDetallesUseCase *application.GetProductoConDetallesUseCase) *GetProductoConDetallesController {
	return &GetProductoConDetallesController{
		getProductoConDetallesUseCase: getProductoConDetallesUseCase,
	}
}

func (ctrl *GetProductoConDetallesController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	detalles, err := ctrl.getProductoConDetallesUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el producto con detalles",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, detalles)
}