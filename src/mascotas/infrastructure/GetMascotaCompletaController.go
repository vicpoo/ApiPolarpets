// GetMascotaCompletaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/application"
)

type GetMascotaCompletaController struct {
	getMascotaCompletaUseCase *application.GetMascotaCompletaUseCase
}

func NewGetMascotaCompletaController(getMascotaCompletaUseCase *application.GetMascotaCompletaUseCase) *GetMascotaCompletaController {
	return &GetMascotaCompletaController{
		getMascotaCompletaUseCase: getMascotaCompletaUseCase,
	}
}

func (ctrl *GetMascotaCompletaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	mascotaDetalles, err := ctrl.getMascotaCompletaUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la mascota con detalles",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, mascotaDetalles)
}