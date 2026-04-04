// CreateRetoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/retos/application"
	"github.com/vicpoo/ApiPolarpets/src/retos/domain/entities"
)

type CreateRetoController struct {
	createUseCase *application.CreateRetoUseCase
}

func NewCreateRetoController(createUseCase *application.CreateRetoUseCase) *CreateRetoController {
	return &CreateRetoController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateRetoController) Run(c *gin.Context) {
	var request struct {
		Titulo         string `json:"titulo"`
		Descripcion    string `json:"descripcion"`
		PuntosGenerados int32  `json:"puntos_generados"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	reto := entities.NewRetos(request.Titulo, request.Descripcion, request.PuntosGenerados)

	createdReto, err := ctrl.createUseCase.Run(reto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el reto",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdReto)
}