// LoginController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/application"
)

type LoginController struct {
	loginUseCase *application.LoginUseCase
}

func NewLoginController(loginUseCase *application.LoginUseCase) *LoginController {
	return &LoginController{
		loginUseCase: loginUseCase,
	}
}

func (ctrl *LoginController) Run(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	usuario, err := ctrl.loginUseCase.Run(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Credenciales inválidas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, usuario)
}