// GetMascotasByNivelController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/application"
)

type GetMascotasByNivelController struct {
	getByNivelUseCase *application.GetMascotasByNivelUseCase
}

func NewGetMascotasByNivelController(getByNivelUseCase *application.GetMascotasByNivelUseCase) *GetMascotasByNivelController {
	return &GetMascotasByNivelController{
		getByNivelUseCase: getByNivelUseCase,
	}
}

func (ctrl *GetMascotasByNivelController) Run(c *gin.Context) {
	idParam := c.Param("id_nivel")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de nivel inválido",
			"error":   err.Error(),
		})
		return
	}

	mascotas, err := ctrl.getByNivelUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las mascotas por nivel",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, mascotas)
}