// UpdateHabitoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/habito/application"
	"github.com/vicpoo/ApiPolarpets/src/habito/domain/entities"
)

type UpdateHabitoController struct {
	updateUseCase *application.UpdateHabitoUseCase
}

func NewUpdateHabitoController(updateUseCase *application.UpdateHabitoUseCase) *UpdateHabitoController {
	return &UpdateHabitoController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateHabitoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var habitoRequest struct {
		IDUser      int32  `json:"id_user"`
		Titulo      string `json:"titulo"`
		Descripcion string `json:"descripcion"`
		Puntos      int32  `json:"puntos"`
	}

	if err := c.ShouldBindJSON(&habitoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	habito := entities.NewHabito(
		habitoRequest.IDUser,
		habitoRequest.Titulo,
		habitoRequest.Descripcion,
		habitoRequest.Puntos,
	)
	habito.SetIDHabito(int32(id))

	updatedHabito, err := ctrl.updateUseCase.Run(habito)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el hábito",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedHabito)
}