// GetMascotasByTipoMascotaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/application"
)

type GetMascotasByTipoMascotaController struct {
	getByTipoMascotaUseCase *application.GetMascotasByTipoMascotaUseCase
}

func NewGetMascotasByTipoMascotaController(getByTipoMascotaUseCase *application.GetMascotasByTipoMascotaUseCase) *GetMascotasByTipoMascotaController {
	return &GetMascotasByTipoMascotaController{
		getByTipoMascotaUseCase: getByTipoMascotaUseCase,
	}
}

func (ctrl *GetMascotasByTipoMascotaController) Run(c *gin.Context) {
	idParam := c.Param("id_tipo_mascota")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de tipo mascota inválido",
			"error":   err.Error(),
		})
		return
	}

	mascotas, err := ctrl.getByTipoMascotaUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las mascotas por tipo",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, mascotas)
}