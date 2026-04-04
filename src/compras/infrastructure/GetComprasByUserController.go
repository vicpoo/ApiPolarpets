// GetComprasByUserController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/compras/application"
)

type GetComprasByUserController struct {
	getByUserUseCase *application.GetComprasByUserUseCase
}

func NewGetComprasByUserController(getByUserUseCase *application.GetComprasByUserUseCase) *GetComprasByUserController {
	return &GetComprasByUserController{
		getByUserUseCase: getByUserUseCase,
	}
}

func (ctrl *GetComprasByUserController) Run(c *gin.Context) {
	idParam := c.Param("id_usuario")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	compras, err := ctrl.getByUserUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las compras del usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, compras)
}