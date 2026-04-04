// GetInventarioByProductoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/application"
)

type GetInventarioByProductoController struct {
	getByProductoUseCase *application.GetInventarioByProductoUseCase
}

func NewGetInventarioByProductoController(getByProductoUseCase *application.GetInventarioByProductoUseCase) *GetInventarioByProductoController {
	return &GetInventarioByProductoController{
		getByProductoUseCase: getByProductoUseCase,
	}
}

func (ctrl *GetInventarioByProductoController) Run(c *gin.Context) {
	idParam := c.Param("id_producto")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de producto inválido",
			"error":   err.Error(),
		})
		return
	}

	inventario, err := ctrl.getByProductoUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el inventario del producto",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, inventario)
}