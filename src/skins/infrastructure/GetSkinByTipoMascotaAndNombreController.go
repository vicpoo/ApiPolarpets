// GetSkinByTipoMascotaAndNombreController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/skins/application"
)

type GetSkinByTipoMascotaAndNombreController struct {
	getByTipoMascotaAndNombreUseCase *application.GetSkinByTipoMascotaAndNombreUseCase
}

func NewGetSkinByTipoMascotaAndNombreController(getByTipoMascotaAndNombreUseCase *application.GetSkinByTipoMascotaAndNombreUseCase) *GetSkinByTipoMascotaAndNombreController {
	return &GetSkinByTipoMascotaAndNombreController{
		getByTipoMascotaAndNombreUseCase: getByTipoMascotaAndNombreUseCase,
	}
}

func (ctrl *GetSkinByTipoMascotaAndNombreController) Run(c *gin.Context) {
	idParam := c.Query("id_tipo_mascota")
	nombre := c.Query("nombre")

	if idParam == "" || nombre == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id_tipo_mascota y nombre son requeridos",
		})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de tipo mascota inválido",
			"error":   err.Error(),
		})
		return
	}

	skin, err := ctrl.getByTipoMascotaAndNombreUseCase.Run(int32(id), nombre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la skin",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, skin)
}