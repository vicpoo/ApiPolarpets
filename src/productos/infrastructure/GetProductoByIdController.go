// GetProductoByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/productos/application"
)

type GetProductoByIdController struct {
	getByIdUseCase *application.GetProductoByIdUseCase
}

func NewGetProductoByIdController(getByIdUseCase *application.GetProductoByIdUseCase) *GetProductoByIdController {
	return &GetProductoByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetProductoByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	producto, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el producto",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, producto)
}