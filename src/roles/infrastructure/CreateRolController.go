// CreateRolController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/roles/application"
	"github.com/vicpoo/ApiPolarpets/src/roles/domain/entities"
)

type CreateRolController struct {
	createUseCase *application.CreateRolUseCase
}

func NewCreateRolController(createUseCase *application.CreateRolUseCase) *CreateRolController {
	return &CreateRolController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateRolController) Run(c *gin.Context) {
	var rolRequest struct {
		Nombre string `json:"nombre"`
	}

	if err := c.ShouldBindJSON(&rolRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	rol := entities.NewRol(rolRequest.Nombre)

	createdRol, err := ctrl.createUseCase.Run(rol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el rol",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdRol)
}