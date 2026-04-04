// GetHabitoByTituloController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/habito/application"
)

type GetHabitoByTituloController struct {
	getByTituloUseCase *application.GetHabitoByTituloUseCase
}

func NewGetHabitoByTituloController(getByTituloUseCase *application.GetHabitoByTituloUseCase) *GetHabitoByTituloController {
	return &GetHabitoByTituloController{
		getByTituloUseCase: getByTituloUseCase,
	}
}

func (ctrl *GetHabitoByTituloController) Run(c *gin.Context) {
	titulo := c.Query("titulo")
	if titulo == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Título es requerido",
		})
		return
	}

	habito, err := ctrl.getByTituloUseCase.Run(titulo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el hábito",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, habito)
}