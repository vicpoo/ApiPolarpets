// GetInventarioByUserAndProductWithDetailsController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/application"
)

type GetInventarioByUserAndProductWithDetailsController struct {
	getWithDetailsByUserAndProductUseCase *application.GetInventarioByUserAndProductWithDetailsUseCase
}

func NewGetInventarioByUserAndProductWithDetailsController(getWithDetailsByUserAndProductUseCase *application.GetInventarioByUserAndProductWithDetailsUseCase) *GetInventarioByUserAndProductWithDetailsController {
	return &GetInventarioByUserAndProductWithDetailsController{
		getWithDetailsByUserAndProductUseCase: getWithDetailsByUserAndProductUseCase,
	}
}

func (ctrl *GetInventarioByUserAndProductWithDetailsController) Run(c *gin.Context) {
	idUsuarioParam := c.Query("id_usuario")
	idProductoParam := c.Query("id_producto")

	if idUsuarioParam == "" || idProductoParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id_usuario y id_producto son requeridos",
		})
		return
	}

	idUsuario, err := strconv.Atoi(idUsuarioParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	idProducto, err := strconv.Atoi(idProductoParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de producto inválido",
			"error":   err.Error(),
		})
		return
	}

	detalle, err := ctrl.getWithDetailsByUserAndProductUseCase.Run(int32(idUsuario), int32(idProducto))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el detalle del inventario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, detalle)
}