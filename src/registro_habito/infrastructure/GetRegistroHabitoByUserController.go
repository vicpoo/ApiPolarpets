// GetRegistroHabitoByUserController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/application"
)

type GetRegistroHabitoByUserController struct {
	getByUserUseCase *application.GetRegistroHabitoByUserUseCase
}

func NewGetRegistroHabitoByUserController(getByUserUseCase *application.GetRegistroHabitoByUserUseCase) *GetRegistroHabitoByUserController {
	return &GetRegistroHabitoByUserController{
		getByUserUseCase: getByUserUseCase,
	}
}

func (ctrl *GetRegistroHabitoByUserController) Run(c *gin.Context) {
	idParam := c.Param("id_user")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	registros, err := ctrl.getByUserUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los registros del usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, registros)
}