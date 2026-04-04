// GetRegistroHabitoByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/application"
)

type GetRegistroHabitoByIdController struct {
	getByIdUseCase *application.GetRegistroHabitoByIdUseCase
}

func NewGetRegistroHabitoByIdController(getByIdUseCase *application.GetRegistroHabitoByIdUseCase) *GetRegistroHabitoByIdController {
	return &GetRegistroHabitoByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetRegistroHabitoByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	registro, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el registro de hábito",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, registro)
}