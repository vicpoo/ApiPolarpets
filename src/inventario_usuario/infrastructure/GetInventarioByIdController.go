// GetInventarioByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/application"
)

type GetInventarioByIdController struct {
	getByIdUseCase *application.GetInventarioByIdUseCase
}

func NewGetInventarioByIdController(getByIdUseCase *application.GetInventarioByIdUseCase) *GetInventarioByIdController {
	return &GetInventarioByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetInventarioByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	inventario, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el registro del inventario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, inventario)
}