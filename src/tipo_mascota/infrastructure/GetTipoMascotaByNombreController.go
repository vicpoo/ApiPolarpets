// GetTipoMascotaByNombreController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/application"
)

type GetTipoMascotaByNombreController struct {
	getByNombreUseCase *application.GetTipoMascotaByNombreUseCase
}

func NewGetTipoMascotaByNombreController(getByNombreUseCase *application.GetTipoMascotaByNombreUseCase) *GetTipoMascotaByNombreController {
	return &GetTipoMascotaByNombreController{
		getByNombreUseCase: getByNombreUseCase,
	}
}

func (ctrl *GetTipoMascotaByNombreController) Run(c *gin.Context) {
	nombre := c.Query("nombre")
	if nombre == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nombre es requerido",
		})
		return
	}

	tipoMascota, err := ctrl.getByNombreUseCase.Run(nombre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el tipo de mascota",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tipoMascota)
}