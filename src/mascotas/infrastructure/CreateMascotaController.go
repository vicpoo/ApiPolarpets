// CreateMascotaController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/application"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/domain/entities"
)

type CreateMascotaController struct {
	createUseCase *application.CreateMascotaUseCase
}

func NewCreateMascotaController(createUseCase *application.CreateMascotaUseCase) *CreateMascotaController {
	return &CreateMascotaController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateMascotaController) Run(c *gin.Context) {
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

	createdMascota, err := ctrl.createUseCase.Run(mascota)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear la mascota",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdMascota)
}