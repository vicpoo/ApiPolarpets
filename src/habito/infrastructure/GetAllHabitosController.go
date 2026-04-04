// GetAllHabitosController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/habito/application"
)

type GetAllHabitosController struct {
	getAllUseCase *application.GetAllHabitosUseCase
}

func NewGetAllHabitosController(getAllUseCase *application.GetAllHabitosUseCase) *GetAllHabitosController {
	return &GetAllHabitosController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllHabitosController) Run(c *gin.Context) {
	habitos, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los hábitos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, habitos)
}