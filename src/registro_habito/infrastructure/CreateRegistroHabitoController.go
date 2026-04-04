// CreateRegistroHabitoController.go
package infrastructure

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/application"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/domain/entities"
)

type CreateRegistroHabitoController struct {
	createUseCase *application.CreateRegistroHabitoUseCase
}

func NewCreateRegistroHabitoController(createUseCase *application.CreateRegistroHabitoUseCase) *CreateRegistroHabitoController {
	return &CreateRegistroHabitoController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateRegistroHabitoController) Run(c *gin.Context) {
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

	createdRegistro, err := ctrl.createUseCase.Run(registro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el registro de hábito",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdRegistro)
}