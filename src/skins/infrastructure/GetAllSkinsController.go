// GetAllSkinsController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/skins/application"
)

type GetAllSkinsController struct {
	getAllUseCase *application.GetAllSkinsUseCase
}

func NewGetAllSkinsController(getAllUseCase *application.GetAllSkinsUseCase) *GetAllSkinsController {
	return &GetAllSkinsController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllSkinsController) Run(c *gin.Context) {
	skins, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las skins",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, skins)
}