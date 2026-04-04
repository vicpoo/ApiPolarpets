// GetRegistroCompletoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/application"
)

type GetRegistroCompletoController struct {
	getRegistroCompletoUseCase *application.GetRegistroCompletoUseCase
}

func NewGetRegistroCompletoController(getRegistroCompletoUseCase *application.GetRegistroCompletoUseCase) *GetRegistroCompletoController {
	return &GetRegistroCompletoController{
		getRegistroCompletoUseCase: getRegistroCompletoUseCase,
	}
}

func (ctrl *GetRegistroCompletoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	detalles, err := ctrl.getRegistroCompletoUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los detalles del registro",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, detalles)
}