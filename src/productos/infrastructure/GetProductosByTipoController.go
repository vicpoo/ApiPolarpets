// GetProductosByTipoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/productos/application"
)

type GetProductosByTipoController struct {
	getByTipoUseCase *application.GetProductosByTipoUseCase
}

func NewGetProductosByTipoController(getByTipoUseCase *application.GetProductosByTipoUseCase) *GetProductosByTipoController {
	return &GetProductosByTipoController{
		getByTipoUseCase: getByTipoUseCase,
	}
}

func (ctrl *GetProductosByTipoController) Run(c *gin.Context) {
	tipo := c.Query("tipo")
	if tipo == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Tipo es requerido",
		})
		return
	}

	productos, err := ctrl.getByTipoUseCase.Run(tipo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los productos por tipo",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, productos)
}