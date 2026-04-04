// GetUserRetoByUserAndRetoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/application"
)

type GetUserRetoByUserAndRetoController struct {
	getByUserAndRetoUseCase *application.GetUserRetoByUserAndRetoUseCase
}

func NewGetUserRetoByUserAndRetoController(getByUserAndRetoUseCase *application.GetUserRetoByUserAndRetoUseCase) *GetUserRetoByUserAndRetoController {
	return &GetUserRetoByUserAndRetoController{
		getByUserAndRetoUseCase: getByUserAndRetoUseCase,
	}
}

func (ctrl *GetUserRetoByUserAndRetoController) Run(c *gin.Context) {
	idUsuarioParam := c.Query("id_usuario")
	idRetoParam := c.Query("id_reto")

	if idUsuarioParam == "" || idRetoParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id_usuario y id_reto son requeridos",
		})
		return
	}

	idUsuario, err := strconv.Atoi(idUsuarioParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	idReto, err := strconv.Atoi(idRetoParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de reto inválido",
			"error":   err.Error(),
		})
		return
	}

	userReto, err := ctrl.getByUserAndRetoUseCase.Run(int32(idUsuario), int32(idReto))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el registro",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userReto)
}