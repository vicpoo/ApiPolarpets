// GetAllUserRetosController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/application"
)

type GetAllUserRetosController struct {
	getAllUseCase *application.GetAllUserRetosUseCase
}

func NewGetAllUserRetosController(getAllUseCase *application.GetAllUserRetosUseCase) *GetAllUserRetosController {
	return &GetAllUserRetosController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllUserRetosController) Run(c *gin.Context) {
	userRetos, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los registros",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userRetos)
}