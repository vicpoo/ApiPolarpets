// GetAllMascotasController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/application"
)

type GetAllMascotasController struct {
	getAllUseCase *application.GetAllMascotasUseCase
}

func NewGetAllMascotasController(getAllUseCase *application.GetAllMascotasUseCase) *GetAllMascotasController {
	return &GetAllMascotasController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllMascotasController) Run(c *gin.Context) {
	mascotas, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las mascotas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, mascotas)
}