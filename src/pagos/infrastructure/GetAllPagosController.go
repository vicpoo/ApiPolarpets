// GetAllPagosController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/pagos/application"
)

type GetAllPagosController struct {
	getAllUseCase *application.GetAllPagosUseCase
}

func NewGetAllPagosController(getAllUseCase *application.GetAllPagosUseCase) *GetAllPagosController {
	return &GetAllPagosController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllPagosController) Run(c *gin.Context) {
	pagos, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los pagos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, pagos)
}