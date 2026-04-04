// GetRegistroHabitoByHabitoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/application"
)

type GetRegistroHabitoByHabitoController struct {
	getByHabitoUseCase *application.GetRegistroHabitoByHabitoUseCase
}

func NewGetRegistroHabitoByHabitoController(getByHabitoUseCase *application.GetRegistroHabitoByHabitoUseCase) *GetRegistroHabitoByHabitoController {
	return &GetRegistroHabitoByHabitoController{
		getByHabitoUseCase: getByHabitoUseCase,
	}
}

func (ctrl *GetRegistroHabitoByHabitoController) Run(c *gin.Context) {
	idParam := c.Param("id_habito")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de hábito inválido",
			"error":   err.Error(),
		})
		return
	}

	registros, err := ctrl.getByHabitoUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los registros del hábito",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, registros)
}