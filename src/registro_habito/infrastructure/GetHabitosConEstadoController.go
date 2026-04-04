// GetHabitosConEstadoController.go
package infrastructure

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/application"
)

type GetHabitosConEstadoController struct {
	getHabitosConEstadoUseCase *application.GetHabitosConEstadoUseCase
}

func NewGetHabitosConEstadoController(getHabitosConEstadoUseCase *application.GetHabitosConEstadoUseCase) *GetHabitosConEstadoController {
	return &GetHabitosConEstadoController{
		getHabitosConEstadoUseCase: getHabitosConEstadoUseCase,
	}
}

func (ctrl *GetHabitosConEstadoController) Run(c *gin.Context) {
	idUserParam := c.Query("id_user")
	fechaStr := c.Query("fecha")

	if idUserParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id_user es requerido",
		})
		return
	}

	idUser, err := strconv.Atoi(idUserParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	var fecha time.Time
	if fechaStr == "" {
		fecha = time.Now()
	} else {
		fecha, err = time.Parse("2006-01-02", fechaStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Formato de fecha inválido. Use YYYY-MM-DD",
				"error":   err.Error(),
			})
			return
		}
	}

	habitos, err := ctrl.getHabitosConEstadoUseCase.Run(int32(idUser), fecha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los hábitos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, habitos)
}