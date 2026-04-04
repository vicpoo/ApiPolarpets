// DeleteProductoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/productos/application"
)

type DeleteProductoController struct {
	deleteUseCase *application.DeleteProductoUseCase
}

func NewDeleteProductoController(deleteUseCase *application.DeleteProductoUseCase) *DeleteProductoController {
	return &DeleteProductoController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteProductoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar el producto",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Producto eliminado exitosamente",
	})
}