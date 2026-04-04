// GetAllProductosConDetallesController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/productos/application"
)

type GetAllProductosConDetallesController struct {
	getAllProductosConDetallesUseCase *application.GetAllProductosConDetallesUseCase
}

func NewGetAllProductosConDetallesController(getAllProductosConDetallesUseCase *application.GetAllProductosConDetallesUseCase) *GetAllProductosConDetallesController {
	return &GetAllProductosConDetallesController{
		getAllProductosConDetallesUseCase: getAllProductosConDetallesUseCase,
	}
}

func (ctrl *GetAllProductosConDetallesController) Run(c *gin.Context) {
	detalles, err := ctrl.getAllProductosConDetallesUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los productos con detalles",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, detalles)
}