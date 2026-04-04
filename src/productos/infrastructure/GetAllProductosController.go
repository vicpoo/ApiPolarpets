// GetAllProductosController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/productos/application"
)

type GetAllProductosController struct {
	getAllUseCase *application.GetAllProductosUseCase
}

func NewGetAllProductosController(getAllUseCase *application.GetAllProductosUseCase) *GetAllProductosController {
	return &GetAllProductosController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllProductosController) Run(c *gin.Context) {
	productos, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los productos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, productos)
}