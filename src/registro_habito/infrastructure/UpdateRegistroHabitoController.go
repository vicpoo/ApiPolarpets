// UpdateRegistroHabitoController.go
package infrastructure

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/application"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/domain/entities"
)

type UpdateRegistroHabitoController struct {
	updateUseCase *application.UpdateRegistroHabitoUseCase
}

func NewUpdateRegistroHabitoController(updateUseCase *application.UpdateRegistroHabitoUseCase) *UpdateRegistroHabitoController {
	return &UpdateRegistroHabitoController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateRegistroHabitoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var request struct {
		IDHabito        int32     `json:"id_habito"`
		FechaRealizada  time.Time `json:"fecha_realizada"`
		PuntosGenerados int32     `json:"puntos_generados"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	registro := entities.NewRegistroHabito(request.IDHabito, request.FechaRealizada, request.PuntosGenerados)
	registro.SetIDRegistroHabito(int32(id))

	updatedRegistro, err := ctrl.updateUseCase.Run(registro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el registro de hábito",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedRegistro)
}