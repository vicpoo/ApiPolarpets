// GetRegistroHabitoByHabitoAndFechaController.go
package infrastructure

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/application"
)

type GetRegistroHabitoByHabitoAndFechaController struct {
	getByHabitoAndFechaUseCase *application.GetRegistroHabitoByHabitoAndFechaUseCase
}

func NewGetRegistroHabitoByHabitoAndFechaController(getByHabitoAndFechaUseCase *application.GetRegistroHabitoByHabitoAndFechaUseCase) *GetRegistroHabitoByHabitoAndFechaController {
	return &GetRegistroHabitoByHabitoAndFechaController{
		getByHabitoAndFechaUseCase: getByHabitoAndFechaUseCase,
	}
}

func (ctrl *GetRegistroHabitoByHabitoAndFechaController) Run(c *gin.Context) {
	idParam := c.Query("id_habito")
	fechaStr := c.Query("fecha")

	if idParam == "" || fechaStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id_habito y fecha son requeridos",
		})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de hábito inválido",
			"error":   err.Error(),
		})
		return
	}

	fecha, err := time.Parse("2006-01-02", fechaStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Formato de fecha inválido. Use YYYY-MM-DD",
			"error":   err.Error(),
		})
		return
	}

	registro, err := ctrl.getByHabitoAndFechaUseCase.Run(int32(id), fecha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el registro",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, registro)
}