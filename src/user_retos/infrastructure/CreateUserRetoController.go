// CreateUserRetoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/application"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/domain/entities"
)

type CreateUserRetoController struct {
	createUseCase *application.CreateUserRetoUseCase
}

func NewCreateUserRetoController(createUseCase *application.CreateUserRetoUseCase) *CreateUserRetoController {
	return &CreateUserRetoController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateUserRetoController) Run(c *gin.Context) {
	var request struct {
		IDUsuario int32 `json:"id_usuario"`
		IDReto    int32 `json:"id_reto"`
		Completo  bool  `json:"completo"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	userReto := entities.NewUserRetos(request.IDUsuario, request.IDReto, request.Completo)

	createdUserReto, err := ctrl.createUseCase.Run(userReto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo asignar el reto al usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdUserReto)
}