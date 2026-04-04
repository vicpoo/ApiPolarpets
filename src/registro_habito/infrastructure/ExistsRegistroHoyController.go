// ExistsRegistroHoyController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/application"
)

type ExistsRegistroHoyController struct {
	existsRegistroHoyUseCase *application.ExistsRegistroHoyUseCase
}

func NewExistsRegistroHoyController(existsRegistroHoyUseCase *application.ExistsRegistroHoyUseCase) *ExistsRegistroHoyController {
	return &ExistsRegistroHoyController{
		existsRegistroHoyUseCase: existsRegistroHoyUseCase,
	}
}

func (ctrl *ExistsRegistroHoyController) Run(c *gin.Context) {
	idParam := c.Query("id_habito")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id_habito es requerido",
		})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de hábito inválido",
			"error":   err.Error(),
		})
		return
	}

	exists, err := ctrl.existsRegistroHoyUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo verificar el registro",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id_habito": id,
		"completado_hoy": exists,
	})
}