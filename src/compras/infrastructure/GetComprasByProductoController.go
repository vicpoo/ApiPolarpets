// GetComprasByProductoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/compras/application"
)

type GetComprasByProductoController struct {
	getByProductoUseCase *application.GetComprasByProductoUseCase
}

func NewGetComprasByProductoController(getByProductoUseCase *application.GetComprasByProductoUseCase) *GetComprasByProductoController {
	return &GetComprasByProductoController{
		getByProductoUseCase: getByProductoUseCase,
	}
}

func (ctrl *GetComprasByProductoController) Run(c *gin.Context) {
	idParam := c.Param("id_producto")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de producto inválido",
			"error":   err.Error(),
		})
		return
	}

	compras, err := ctrl.getByProductoUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las compras del producto",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, compras)
}