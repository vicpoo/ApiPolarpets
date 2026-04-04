// UpdateTipoMascotaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/application"
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain/entities"
)

type UpdateTipoMascotaController struct {
	updateUseCase *application.UpdateTipoMascotaUseCase
}

func NewUpdateTipoMascotaController(updateUseCase *application.UpdateTipoMascotaUseCase) *UpdateTipoMascotaController {
	return &UpdateTipoMascotaController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateTipoMascotaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var tipoMascotaRequest struct {
		Nombre      string `json:"nombre"`
		Descripcion string `json:"descripcion"`
	}

	if err := c.ShouldBindJSON(&tipoMascotaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	tipoMascota := entities.NewTipoMascota(tipoMascotaRequest.Nombre, tipoMascotaRequest.Descripcion)
	tipoMascota.SetIDTipoMascota(int32(id))

	updatedTipoMascota, err := ctrl.updateUseCase.Run(tipoMascota)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el tipo de mascota",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedTipoMascota)
}