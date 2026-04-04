// GetCompraByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/compras/application"
)

type GetCompraByIdController struct {
	getByIdUseCase *application.GetCompraByIdUseCase
}

func NewGetCompraByIdController(getByIdUseCase *application.GetCompraByIdUseCase) *GetCompraByIdController {
	return &GetCompraByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetCompraByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	compra, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la compra",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, compra)
}