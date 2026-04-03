// RegisterController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/application"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/domain/entities"
)

type RegisterController struct {
	registerUseCase *application.RegisterUseCase
}

func NewRegisterController(registerUseCase *application.RegisterUseCase) *RegisterController {
	return &RegisterController{
		registerUseCase: registerUseCase,
	}
}

func (ctrl *RegisterController) Run(c *gin.Context) {
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

	createdUsuario, err := ctrl.registerUseCase.Run(usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo registrar el usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdUsuario)
}