// GetCompletedRetosByUserController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/application"
)

type GetCompletedRetosByUserController struct {
	getCompletedRetosUseCase *application.GetCompletedRetosByUserUseCase
}

func NewGetCompletedRetosByUserController(getCompletedRetosUseCase *application.GetCompletedRetosByUserUseCase) *GetCompletedRetosByUserController {
	return &GetCompletedRetosByUserController{
		getCompletedRetosUseCase: getCompletedRetosUseCase,
	}
}

func (ctrl *GetCompletedRetosByUserController) Run(c *gin.Context) {
	idParam := c.Param("id_usuario")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	completedRetos, err := ctrl.getCompletedRetosUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los retos completados",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, completedRetos)
}