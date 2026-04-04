// GetAllRegistroHabitosController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/application"
)

type GetAllRegistroHabitosController struct {
	getAllUseCase *application.GetAllRegistroHabitosUseCase
}

func NewGetAllRegistroHabitosController(getAllUseCase *application.GetAllRegistroHabitosUseCase) *GetAllRegistroHabitosController {
	return &GetAllRegistroHabitosController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllRegistroHabitosController) Run(c *gin.Context) {
	registros, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los registros de hábitos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, registros)
}