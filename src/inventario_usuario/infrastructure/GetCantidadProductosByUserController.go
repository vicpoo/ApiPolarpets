// GetCantidadProductosByUserController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/application"
)

type GetCantidadProductosByUserController struct {
	getCantidadUseCase *application.GetCantidadProductosByUserUseCase
}

func NewGetCantidadProductosByUserController(getCantidadUseCase *application.GetCantidadProductosByUserUseCase) *GetCantidadProductosByUserController {
	return &GetCantidadProductosByUserController{
		getCantidadUseCase: getCantidadUseCase,
	}
}

func (ctrl *GetCantidadProductosByUserController) Run(c *gin.Context) {
	idParam := c.Param("id_usuario")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	cantidad, err := ctrl.getCantidadUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la cantidad de productos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id_usuario": id,
		"cantidad":   cantidad,
	})
}