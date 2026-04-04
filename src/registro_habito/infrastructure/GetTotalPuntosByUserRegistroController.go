// GetTotalPuntosByUserRegistroController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/application"
)

type GetTotalPuntosByUserRegistroController struct {
	getTotalPuntosByUserUseCase *application.GetTotalPuntosByUserRegistroUseCase
}

func NewGetTotalPuntosByUserRegistroController(getTotalPuntosByUserUseCase *application.GetTotalPuntosByUserRegistroUseCase) *GetTotalPuntosByUserRegistroController {
	return &GetTotalPuntosByUserRegistroController{
		getTotalPuntosByUserUseCase: getTotalPuntosByUserUseCase,
	}
}

func (ctrl *GetTotalPuntosByUserRegistroController) Run(c *gin.Context) {
	idParam := c.Param("id_user")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	totalPuntos, err := ctrl.getTotalPuntosByUserUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los puntos totales del usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id_user":      id,
		"total_puntos": totalPuntos,
	})
}