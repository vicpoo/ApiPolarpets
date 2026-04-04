// GetUserRetosConDetallesController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/application"
)

type GetUserRetosConDetallesController struct {
	getUserRetosConDetallesUseCase *application.GetUserRetosConDetallesUseCase
}

func NewGetUserRetosConDetallesController(getUserRetosConDetallesUseCase *application.GetUserRetosConDetallesUseCase) *GetUserRetosConDetallesController {
	return &GetUserRetosConDetallesController{
		getUserRetosConDetallesUseCase: getUserRetosConDetallesUseCase,
	}
}

func (ctrl *GetUserRetosConDetallesController) Run(c *gin.Context) {
	idParam := c.Param("id_usuario")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	detalles, err := ctrl.getUserRetosConDetallesUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los detalles de los retos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, detalles)
}