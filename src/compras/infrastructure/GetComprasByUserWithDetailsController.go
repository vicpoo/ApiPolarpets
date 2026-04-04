// GetComprasByUserWithDetailsController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/compras/application"
)

type GetComprasByUserWithDetailsController struct {
	getWithDetailsUseCase *application.GetComprasByUserWithDetailsUseCase
}

func NewGetComprasByUserWithDetailsController(getWithDetailsUseCase *application.GetComprasByUserWithDetailsUseCase) *GetComprasByUserWithDetailsController {
	return &GetComprasByUserWithDetailsController{
		getWithDetailsUseCase: getWithDetailsUseCase,
	}
}

func (ctrl *GetComprasByUserWithDetailsController) Run(c *gin.Context) {
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
			"message": "No se pudieron obtener los detalles de las compras",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, detalles)
}