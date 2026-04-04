// GetSkinsByTipoMascotaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/skins/application"
)

type GetSkinsByTipoMascotaController struct {
	getByTipoMascotaUseCase *application.GetSkinsByTipoMascotaUseCase
}

func NewGetSkinsByTipoMascotaController(getByTipoMascotaUseCase *application.GetSkinsByTipoMascotaUseCase) *GetSkinsByTipoMascotaController {
	return &GetSkinsByTipoMascotaController{
		getByTipoMascotaUseCase: getByTipoMascotaUseCase,
	}
}

func (ctrl *GetSkinsByTipoMascotaController) Run(c *gin.Context) {
	idParam := c.Param("id_tipo_mascota")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de tipo mascota inválido",
			"error":   err.Error(),
		})
		return
	}

	skins, err := ctrl.getByTipoMascotaUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las skins",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, skins)
}