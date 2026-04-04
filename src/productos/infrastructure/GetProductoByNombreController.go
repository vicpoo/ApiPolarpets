// GetProductoByNombreController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/productos/application"
)

type GetProductoByNombreController struct {
	getByNombreUseCase *application.GetProductoByNombreUseCase
}

func NewGetProductoByNombreController(getByNombreUseCase *application.GetProductoByNombreUseCase) *GetProductoByNombreController {
	return &GetProductoByNombreController{
		getByNombreUseCase: getByNombreUseCase,
	}
}

func (ctrl *GetProductoByNombreController) Run(c *gin.Context) {
	nombre := c.Query("nombre")
	if nombre == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nombre es requerido",
		})
		return
	}

	producto, err := ctrl.getByNombreUseCase.Run(nombre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el producto",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, producto)
}