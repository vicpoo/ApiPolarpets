// GetInventarioByUserController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/application"
)

type GetInventarioByUserController struct {
	getByUserUseCase *application.GetInventarioByUserUseCase
}

func NewGetInventarioByUserController(getByUserUseCase *application.GetInventarioByUserUseCase) *GetInventarioByUserController {
	return &GetInventarioByUserController{
		getByUserUseCase: getByUserUseCase,
	}
}

func (ctrl *GetInventarioByUserController) Run(c *gin.Context) {
	idParam := c.Param("id_usuario")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	inventario, err := ctrl.getByUserUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el inventario del usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, inventario)
}