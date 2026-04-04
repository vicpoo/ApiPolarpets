// GetRetoByTituloController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/retos/application"
)

type GetRetoByTituloController struct {
	getByTituloUseCase *application.GetRetoByTituloUseCase
}

func NewGetRetoByTituloController(getByTituloUseCase *application.GetRetoByTituloUseCase) *GetRetoByTituloController {
	return &GetRetoByTituloController{
		getByTituloUseCase: getByTituloUseCase,
	}
}

func (ctrl *GetRetoByTituloController) Run(c *gin.Context) {
	titulo := c.Query("titulo")
	if titulo == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Título es requerido",
		})
		return
	}

	reto, err := ctrl.getByTituloUseCase.Run(titulo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el reto",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reto)
}