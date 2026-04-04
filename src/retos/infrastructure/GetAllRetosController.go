// GetAllRetosController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/retos/application"
)

type GetAllRetosController struct {
	getAllUseCase *application.GetAllRetosUseCase
}

func NewGetAllRetosController(getAllUseCase *application.GetAllRetosUseCase) *GetAllRetosController {
	return &GetAllRetosController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllRetosController) Run(c *gin.Context) {
	retos, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los retos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, retos)
}