// GetUserRetosByUserController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/application"
)

type GetUserRetosByUserController struct {
	getByUserUseCase *application.GetUserRetosByUserUseCase
}

func NewGetUserRetosByUserController(getByUserUseCase *application.GetUserRetosByUserUseCase) *GetUserRetosByUserController {
	return &GetUserRetosByUserController{
		getByUserUseCase: getByUserUseCase,
	}
}

func (ctrl *GetUserRetosByUserController) Run(c *gin.Context) {
	idParam := c.Param("id_usuario")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	userRetos, err := ctrl.getByUserUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los retos del usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userRetos)
}