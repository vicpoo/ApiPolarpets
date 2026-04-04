// DeleteInventarioByUserAndProductController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/application"
)

type DeleteInventarioByUserAndProductController struct {
	deleteByUserAndProductUseCase *application.DeleteInventarioByUserAndProductUseCase
}

func NewDeleteInventarioByUserAndProductController(deleteByUserAndProductUseCase *application.DeleteInventarioByUserAndProductUseCase) *DeleteInventarioByUserAndProductController {
	return &DeleteInventarioByUserAndProductController{
		deleteByUserAndProductUseCase: deleteByUserAndProductUseCase,
	}
}

func (ctrl *DeleteInventarioByUserAndProductController) Run(c *gin.Context) {
	idUsuarioParam := c.Param("id_usuario")
	idProductoParam := c.Param("id_producto")

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

	errDelete := ctrl.deleteByUserAndProductUseCase.Run(int32(idUsuario), int32(idProducto))
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar el producto del inventario del usuario",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Producto eliminado del inventario del usuario exitosamente",
	})
}