// GetTipoMascotaByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/application"
)

type GetTipoMascotaByIdController struct {
	getByIdUseCase *application.GetTipoMascotaByIdUseCase
}

func NewGetTipoMascotaByIdController(getByIdUseCase *application.GetTipoMascotaByIdUseCase) *GetTipoMascotaByIdController {
	return &GetTipoMascotaByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetTipoMascotaByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	tipoMascota, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el tipo de mascota",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tipoMascota)
}