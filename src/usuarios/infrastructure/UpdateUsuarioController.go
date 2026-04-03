// UpdateUsuarioController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/application"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/domain/entities"
)

type UpdateUsuarioController struct {
	updateUseCase *application.UpdateUsuarioUseCase
}

func NewUpdateUsuarioController(updateUseCase *application.UpdateUsuarioUseCase) *UpdateUsuarioController {
	return &UpdateUsuarioController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateUsuarioController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

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
	usuario.SetIDUsuario(int32(id))

	updatedUsuario, err := ctrl.updateUseCase.Run(usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedUsuario)
}