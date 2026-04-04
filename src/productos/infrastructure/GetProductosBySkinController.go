// GetProductosBySkinController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/productos/application"
)

type GetProductosBySkinController struct {
	getBySkinUseCase *application.GetProductosBySkinUseCase
}

func NewGetProductosBySkinController(getBySkinUseCase *application.GetProductosBySkinUseCase) *GetProductosBySkinController {
	return &GetProductosBySkinController{
		getBySkinUseCase: getBySkinUseCase,
	}
}

func (ctrl *GetProductosBySkinController) Run(c *gin.Context) {
	idParam := c.Param("id_skin")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de skin inválido",
			"error":   err.Error(),
		})
		return
	}

	productos, err := ctrl.getBySkinUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los productos por skin",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, productos)
}