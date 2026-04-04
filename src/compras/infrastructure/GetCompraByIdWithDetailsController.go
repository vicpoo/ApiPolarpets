// GetCompraByIdWithDetailsController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/compras/application"
)

type GetCompraByIdWithDetailsController struct {
	getByIdWithDetailsUseCase *application.GetCompraByIdWithDetailsUseCase
}

func NewGetCompraByIdWithDetailsController(getByIdWithDetailsUseCase *application.GetCompraByIdWithDetailsUseCase) *GetCompraByIdWithDetailsController {
	return &GetCompraByIdWithDetailsController{
		getByIdWithDetailsUseCase: getByIdWithDetailsUseCase,
	}
}

func (ctrl *GetCompraByIdWithDetailsController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	detalle, err := ctrl.getByIdWithDetailsUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el detalle de la compra",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, detalle)
}