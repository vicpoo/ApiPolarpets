// CreateNivelController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/niveles/application"
	"github.com/vicpoo/ApiPolarpets/src/niveles/domain/entities"
)

type CreateNivelController struct {
	createUseCase *application.CreateNivelUseCase
}

func NewCreateNivelController(createUseCase *application.CreateNivelUseCase) *CreateNivelController {
	return &CreateNivelController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateNivelController) Run(c *gin.Context) {
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

	createdNivel, err := ctrl.createUseCase.Run(nivel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el nivel",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdNivel)
}