// GetProductosByTipoInInventoryController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/application"
)

type GetProductosByTipoInInventoryController struct {
	getByTipoUseCase *application.GetProductosByTipoInInventoryUseCase
}

func NewGetProductosByTipoInInventoryController(getByTipoUseCase *application.GetProductosByTipoInInventoryUseCase) *GetProductosByTipoInInventoryController {
	return &GetProductosByTipoInInventoryController{
		getByTipoUseCase: getByTipoUseCase,
	}
}

func (ctrl *GetProductosByTipoInInventoryController) Run(c *gin.Context) {
	idParam := c.Param("id_usuario")
	tipo := c.Query("tipo")

	if idParam == "" || tipo == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id_usuario y tipo son requeridos",
		})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	inventario, err := ctrl.getByTipoUseCase.Run(int32(id), tipo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el inventario por tipo",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, inventario)
}