// GetAllComprasController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/compras/application"
)

type GetAllComprasController struct {
	getAllUseCase *application.GetAllComprasUseCase
}

func NewGetAllComprasController(getAllUseCase *application.GetAllComprasUseCase) *GetAllComprasController {
	return &GetAllComprasController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllComprasController) Run(c *gin.Context) {
	compras, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las compras",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, compras)
}