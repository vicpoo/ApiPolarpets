// GetMascotasBySkinController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/application"
)

type GetMascotasBySkinController struct {
	getBySkinUseCase *application.GetMascotasBySkinUseCase
}

func NewGetMascotasBySkinController(getBySkinUseCase *application.GetMascotasBySkinUseCase) *GetMascotasBySkinController {
	return &GetMascotasBySkinController{
		getBySkinUseCase: getBySkinUseCase,
	}
}

func (ctrl *GetMascotasBySkinController) Run(c *gin.Context) {
	idParam := c.Param("id_skin")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de skin inválido",
			"error":   err.Error(),
		})
		return
	}

	mascotas, err := ctrl.getBySkinUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las mascotas por skin",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, mascotas)
}