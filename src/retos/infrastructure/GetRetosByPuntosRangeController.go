// GetRetosByPuntosRangeController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/retos/application"
)

type GetRetosByPuntosRangeController struct {
	getByPuntosRangeUseCase *application.GetRetosByPuntosRangeUseCase
}

func NewGetRetosByPuntosRangeController(getByPuntosRangeUseCase *application.GetRetosByPuntosRangeUseCase) *GetRetosByPuntosRangeController {
	return &GetRetosByPuntosRangeController{
		getByPuntosRangeUseCase: getByPuntosRangeUseCase,
	}
}

func (ctrl *GetRetosByPuntosRangeController) Run(c *gin.Context) {
	minPuntosStr := c.Query("min_puntos")
	maxPuntosStr := c.Query("max_puntos")

	if minPuntosStr == "" || maxPuntosStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "min_puntos y max_puntos son requeridos",
		})
		return
	}

	minPuntos, err := strconv.Atoi(minPuntosStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "min_puntos inválido",
			"error":   err.Error(),
		})
		return
	}

	maxPuntos, err := strconv.Atoi(maxPuntosStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "max_puntos inválido",
			"error":   err.Error(),
		})
		return
	}

	retos, err := ctrl.getByPuntosRangeUseCase.Run(int32(minPuntos), int32(maxPuntos))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los retos por rango de puntos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, retos)
}