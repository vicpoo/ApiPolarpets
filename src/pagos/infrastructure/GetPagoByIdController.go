// GetPagoByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/pagos/application"
)

type GetPagoByIdController struct {
	getByIdUseCase *application.GetPagoByIdUseCase
}

func NewGetPagoByIdController(getByIdUseCase *application.GetPagoByIdUseCase) *GetPagoByIdController {
	return &GetPagoByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetPagoByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	pago, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el pago",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pago)
}