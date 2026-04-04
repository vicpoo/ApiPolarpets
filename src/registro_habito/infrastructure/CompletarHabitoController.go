// CompletarHabitoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/application"
)

type CompletarHabitoController struct {
	completarHabitoUseCase *application.CompletarHabitoUseCase
}

func NewCompletarHabitoController(completarHabitoUseCase *application.CompletarHabitoUseCase) *CompletarHabitoController {
	return &CompletarHabitoController{
		completarHabitoUseCase: completarHabitoUseCase,
	}
}

func (ctrl *CompletarHabitoController) Run(c *gin.Context) {
	var request struct {
		IDHabito int32 `json:"id_habito"`
		IDUser   int32 `json:"id_user"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	err := ctrl.completarHabitoUseCase.Run(request.IDHabito, request.IDUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo completar el hábito",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Hábito completado exitosamente",
		"message": "¡Felicidades! Has ganado puntos",
	})
}