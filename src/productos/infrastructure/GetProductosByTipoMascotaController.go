// GetProductosByTipoMascotaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/productos/application"
)

type GetProductosByTipoMascotaController struct {
	getByTipoMascotaUseCase *application.GetProductosByTipoMascotaUseCase
}

func NewGetProductosByTipoMascotaController(getByTipoMascotaUseCase *application.GetProductosByTipoMascotaUseCase) *GetProductosByTipoMascotaController {
	return &GetProductosByTipoMascotaController{
		getByTipoMascotaUseCase: getByTipoMascotaUseCase,
	}
}

func (ctrl *GetProductosByTipoMascotaController) Run(c *gin.Context) {
	idParam := c.Param("id_tipo_mascota")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de tipo mascota inválido",
			"error":   err.Error(),
		})
		return
	}

	productos, err := ctrl.getByTipoMascotaUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los productos por tipo de mascota",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, productos)
}