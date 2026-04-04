// UpdateExperienciaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/application"
)

type UpdateExperienciaController struct {
	updateExperienciaUseCase *application.UpdateExperienciaUseCase
}

func NewUpdateExperienciaController(updateExperienciaUseCase *application.UpdateExperienciaUseCase) *UpdateExperienciaController {
	return &UpdateExperienciaController{
		updateExperienciaUseCase: updateExperienciaUseCase,
	}
}

func (ctrl *UpdateExperienciaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var experienciaRequest struct {
		ExperienciaActual int32 `json:"experiencia_actual"`
	}

	if err := c.ShouldBindJSON(&experienciaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	err = ctrl.updateExperienciaUseCase.Run(int32(id), experienciaRequest.ExperienciaActual)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la experiencia",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Experiencia actualizada exitosamente",
	})
}