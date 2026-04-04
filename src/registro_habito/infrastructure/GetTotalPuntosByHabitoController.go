// GetTotalPuntosByHabitoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/application"
)

type GetTotalPuntosByHabitoController struct {
	getTotalPuntosByHabitoUseCase *application.GetTotalPuntosByHabitoUseCase
}

func NewGetTotalPuntosByHabitoController(getTotalPuntosByHabitoUseCase *application.GetTotalPuntosByHabitoUseCase) *GetTotalPuntosByHabitoController {
	return &GetTotalPuntosByHabitoController{
		getTotalPuntosByHabitoUseCase: getTotalPuntosByHabitoUseCase,
	}
}

func (ctrl *GetTotalPuntosByHabitoController) Run(c *gin.Context) {
	idParam := c.Param("id_habito")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de hábito inválido",
			"error":   err.Error(),
		})
		return
	}

	totalPuntos, err := ctrl.getTotalPuntosByHabitoUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los puntos totales del hábito",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id_habito":     id,
		"total_puntos": totalPuntos,
	})
}