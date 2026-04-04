// GetHabitoByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/habito/application"
)

type GetHabitoByIdController struct {
	getByIdUseCase *application.GetHabitoByIdUseCase
}

func NewGetHabitoByIdController(getByIdUseCase *application.GetHabitoByIdUseCase) *GetHabitoByIdController {
	return &GetHabitoByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetHabitoByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	habito, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el hábito",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, habito)
}