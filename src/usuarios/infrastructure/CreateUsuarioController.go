// CreateUsuarioController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/application"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/domain/entities"
)

type CreateUsuarioController struct {
	createUseCase *application.CreateUsuarioUseCase
}

func NewCreateUsuarioController(createUseCase *application.CreateUsuarioUseCase) *CreateUsuarioController {
	return &CreateUsuarioController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateUsuarioController) Run(c *gin.Context) {
	var usuarioRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		IDRol    int32  `json:"id_rol"`
	}

	if err := c.ShouldBindJSON(&usuarioRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	usuario := entities.NewUsuario(
		usuarioRequest.Username,
		usuarioRequest.Email,
		usuarioRequest.Password,
		usuarioRequest.IDRol,
	)

	createdUsuario, err := ctrl.createUseCase.Run(usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdUsuario)
}