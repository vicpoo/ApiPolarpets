// CompleteRetoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/application"
)

type CompleteRetoController struct {
	completeRetoUseCase *application.CompleteRetoUseCase
}

func NewCompleteRetoController(completeRetoUseCase *application.CompleteRetoUseCase) *CompleteRetoController {
	return &CompleteRetoController{
		completeRetoUseCase: completeRetoUseCase,
	}
}

func (ctrl *CompleteRetoController) Run(c *gin.Context) {
	var request struct {
		IDUsuario int32 `json:"id_usuario"`
		IDReto    int32 `json:"id_reto"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	err := ctrl.completeRetoUseCase.Run(request.IDUsuario, request.IDReto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo completar el reto",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Reto completado exitosamente",
		"message": "¡Felicidades! Has ganado puntos",
	})
}