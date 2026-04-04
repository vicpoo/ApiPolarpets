// GetHabitosByUserController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/habito/application"
)

type GetHabitosByUserController struct {
	getByUserUseCase *application.GetHabitosByUserUseCase
}

func NewGetHabitosByUserController(getByUserUseCase *application.GetHabitosByUserUseCase) *GetHabitosByUserController {
	return &GetHabitosByUserController{
		getByUserUseCase: getByUserUseCase,
	}
}

func (ctrl *GetHabitosByUserController) Run(c *gin.Context) {
	idParam := c.Param("id_user")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	habitos, err := ctrl.getByUserUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los hábitos del usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, habitos)
}