// GetCompraByPagoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/compras/application"
)

type GetCompraByPagoController struct {
	getByPagoUseCase *application.GetCompraByPagoUseCase
}

func NewGetCompraByPagoController(getByPagoUseCase *application.GetCompraByPagoUseCase) *GetCompraByPagoController {
	return &GetCompraByPagoController{
		getByPagoUseCase: getByPagoUseCase,
	}
}

func (ctrl *GetCompraByPagoController) Run(c *gin.Context) {
	idParam := c.Param("id_pago")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de pago inválido",
			"error":   err.Error(),
		})
		return
	}

	compra, err := ctrl.getByPagoUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la compra asociada al pago",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, compra)
}