// GetInventarioByUserWithDetailsController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/application"
)

type GetInventarioByUserWithDetailsController struct {
	getWithDetailsUseCase *application.GetInventarioByUserWithDetailsUseCase
}

func NewGetInventarioByUserWithDetailsController(getWithDetailsUseCase *application.GetInventarioByUserWithDetailsUseCase) *GetInventarioByUserWithDetailsController {
	return &GetInventarioByUserWithDetailsController{
		getWithDetailsUseCase: getWithDetailsUseCase,
	}
}

func (ctrl *GetInventarioByUserWithDetailsController) Run(c *gin.Context) {
	idParam := c.Param("id_usuario")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	detalles, err := ctrl.getWithDetailsUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el inventario con detalles",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, detalles)
}