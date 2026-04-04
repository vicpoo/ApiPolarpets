// GetAllTiposMascotaController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/application"
)

type GetAllTiposMascotaController struct {
	getAllUseCase *application.GetAllTiposMascotaUseCase
}

func NewGetAllTiposMascotaController(getAllUseCase *application.GetAllTiposMascotaUseCase) *GetAllTiposMascotaController {
	return &GetAllTiposMascotaController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllTiposMascotaController) Run(c *gin.Context) {
	tiposMascota, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los tipos de mascota",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tiposMascota)
}