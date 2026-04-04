// GetTotalPagadoByUserController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/pagos/application"
)

type GetTotalPagadoByUserController struct {
	getTotalPagadoUseCase *application.GetTotalPagadoByUserUseCase
}

func NewGetTotalPagadoByUserController(getTotalPagadoUseCase *application.GetTotalPagadoByUserUseCase) *GetTotalPagadoByUserController {
	return &GetTotalPagadoByUserController{
		getTotalPagadoUseCase: getTotalPagadoUseCase,
	}
}

func (ctrl *GetTotalPagadoByUserController) Run(c *gin.Context) {
	idParam := c.Param("id_usuario")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	total, err := ctrl.getTotalPagadoUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el total pagado",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id_usuario": id,
		"total":      total,
	})
}