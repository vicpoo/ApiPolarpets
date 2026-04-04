// GetAllInventarioController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/application"
)

type GetAllInventarioController struct {
	getAllUseCase *application.GetAllInventarioUseCase
}

func NewGetAllInventarioController(getAllUseCase *application.GetAllInventarioUseCase) *GetAllInventarioController {
	return &GetAllInventarioController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllInventarioController) Run(c *gin.Context) {
	inventario, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el inventario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, inventario)
}