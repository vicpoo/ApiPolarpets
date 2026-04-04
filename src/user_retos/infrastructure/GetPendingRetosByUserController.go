// GetPendingRetosByUserController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/application"
)

type GetPendingRetosByUserController struct {
	getPendingRetosUseCase *application.GetPendingRetosByUserUseCase
}

func NewGetPendingRetosByUserController(getPendingRetosUseCase *application.GetPendingRetosByUserUseCase) *GetPendingRetosByUserController {
	return &GetPendingRetosByUserController{
		getPendingRetosUseCase: getPendingRetosUseCase,
	}
}

func (ctrl *GetPendingRetosByUserController) Run(c *gin.Context) {
	idParam := c.Param("id_usuario")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	pendingRetos, err := ctrl.getPendingRetosUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los retos pendientes",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pendingRetos)
}