// UpdateMascotaActivaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/application"
)

type UpdateMascotaActivaController struct {
	updateMascotaActivaUseCase *application.UpdateMascotaActivaUseCase
}

func NewUpdateMascotaActivaController(updateMascotaActivaUseCase *application.UpdateMascotaActivaUseCase) *UpdateMascotaActivaController {
	return &UpdateMascotaActivaController{
		updateMascotaActivaUseCase: updateMascotaActivaUseCase,
	}
}

func (ctrl *UpdateMascotaActivaController) Run(c *gin.Context) {
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
		IDMascotaActiva *int32 `json:"id_mascota_activa"`
	}

	if err := c.ShouldBindJSON(&mascotaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	err = ctrl.updateMascotaActivaUseCase.Run(int32(id), mascotaRequest.IDMascotaActiva)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la mascota activa",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Mascota activa actualizada exitosamente",
	})
}