// CreateHabitoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/habito/application"
	"github.com/vicpoo/ApiPolarpets/src/habito/domain/entities"
)

type CreateHabitoController struct {
	createUseCase *application.CreateHabitoUseCase
}

func NewCreateHabitoController(createUseCase *application.CreateHabitoUseCase) *CreateHabitoController {
	return &CreateHabitoController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateHabitoController) Run(c *gin.Context) {
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

	createdHabito, err := ctrl.createUseCase.Run(habito)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el hábito",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdHabito)
}