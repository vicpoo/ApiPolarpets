// UpdateMascotaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/application"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/domain/entities"
)

type UpdateMascotaController struct {
	updateUseCase *application.UpdateMascotaUseCase
}

func NewUpdateMascotaController(updateUseCase *application.UpdateMascotaUseCase) *UpdateMascotaController {
	return &UpdateMascotaController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateMascotaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var mascotaRequest struct {
		IDUser           int32 `json:"id_user"`
		IDNiveles        int32 `json:"id_niveles"`
		IDSkin           int32 `json:"id_skin"`
		IDTipoMascota    int32 `json:"id_tipo_mascota"`
		ExperienciaActual int32 `json:"experiencia_actual"`
	}

	if err := c.ShouldBindJSON(&mascotaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	mascota := entities.NewMascota(
		mascotaRequest.IDUser,
		mascotaRequest.IDNiveles,
		mascotaRequest.IDSkin,
		mascotaRequest.IDTipoMascota,
		mascotaRequest.ExperienciaActual,
	)
	mascota.SetIDMascota(int32(id))

	updatedMascota, err := ctrl.updateUseCase.Run(mascota)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la mascota",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedMascota)
}