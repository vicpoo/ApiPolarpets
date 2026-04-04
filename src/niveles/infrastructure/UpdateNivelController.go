// UpdateNivelController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/niveles/application"
	"github.com/vicpoo/ApiPolarpets/src/niveles/domain/entities"
)

type UpdateNivelController struct {
	updateUseCase *application.UpdateNivelUseCase
}

func NewUpdateNivelController(updateUseCase *application.UpdateNivelUseCase) *UpdateNivelController {
	return &UpdateNivelController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateNivelController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var nivelRequest struct {
		Nivel        int32 `json:"nivel"`
		ExpRequerida int32 `json:"exp_requerida"`
	}

	if err := c.ShouldBindJSON(&nivelRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	nivel := entities.NewNiveles(nivelRequest.Nivel, nivelRequest.ExpRequerida)
	nivel.SetIDNiveles(int32(id))

	updatedNivel, err := ctrl.updateUseCase.Run(nivel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el nivel",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedNivel)
}