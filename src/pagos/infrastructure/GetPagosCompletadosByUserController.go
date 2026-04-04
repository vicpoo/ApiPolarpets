// GetPagosCompletadosByUserController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/pagos/application"
)

type GetPagosCompletadosByUserController struct {
	getPagosCompletadosUseCase *application.GetPagosCompletadosByUserUseCase
}

func NewGetPagosCompletadosByUserController(getPagosCompletadosUseCase *application.GetPagosCompletadosByUserUseCase) *GetPagosCompletadosByUserController {
	return &GetPagosCompletadosByUserController{
		getPagosCompletadosUseCase: getPagosCompletadosUseCase,
	}
}

func (ctrl *GetPagosCompletadosByUserController) Run(c *gin.Context) {
	idParam := c.Param("id_usuario")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	pagos, err := ctrl.getPagosCompletadosUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los pagos completados",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pagos)
}