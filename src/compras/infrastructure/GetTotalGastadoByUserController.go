// GetTotalGastadoByUserController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/compras/application"
)

type GetTotalGastadoByUserController struct {
	getTotalGastadoUseCase *application.GetTotalGastadoByUserUseCase
}

func NewGetTotalGastadoByUserController(getTotalGastadoUseCase *application.GetTotalGastadoByUserUseCase) *GetTotalGastadoByUserController {
	return &GetTotalGastadoByUserController{
		getTotalGastadoUseCase: getTotalGastadoUseCase,
	}
}

func (ctrl *GetTotalGastadoByUserController) Run(c *gin.Context) {
	idParam := c.Param("id_usuario")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	total, err := ctrl.getTotalGastadoUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el total gastado",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id_usuario":     id,
		"total_gastado": total,
	})
}