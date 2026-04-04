// GetUserRetosByRetoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/application"
)

type GetUserRetosByRetoController struct {
	getByRetoUseCase *application.GetUserRetosByRetoUseCase
}

func NewGetUserRetosByRetoController(getByRetoUseCase *application.GetUserRetosByRetoUseCase) *GetUserRetosByRetoController {
	return &GetUserRetosByRetoController{
		getByRetoUseCase: getByRetoUseCase,
	}
}

func (ctrl *GetUserRetosByRetoController) Run(c *gin.Context) {
	idParam := c.Param("id_reto")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de reto inválido",
			"error":   err.Error(),
		})
		return
	}

	userRetos, err := ctrl.getByRetoUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los usuarios con este reto",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userRetos)
}