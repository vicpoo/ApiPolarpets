// CreateTipoMascotaController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/application"
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain/entities"
)

type CreateTipoMascotaController struct {
	createUseCase *application.CreateTipoMascotaUseCase
}

func NewCreateTipoMascotaController(createUseCase *application.CreateTipoMascotaUseCase) *CreateTipoMascotaController {
	return &CreateTipoMascotaController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateTipoMascotaController) Run(c *gin.Context) {
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

	createdTipoMascota, err := ctrl.createUseCase.Run(tipoMascota)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el tipo de mascota",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdTipoMascota)
}