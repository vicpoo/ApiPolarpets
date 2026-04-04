// GetAllNivelesController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/niveles/application"
)

type GetAllNivelesController struct {
	getAllUseCase *application.GetAllNivelesUseCase
}

func NewGetAllNivelesController(getAllUseCase *application.GetAllNivelesUseCase) *GetAllNivelesController {
	return &GetAllNivelesController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllNivelesController) Run(c *gin.Context) {
	niveles, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los niveles",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, niveles)
}