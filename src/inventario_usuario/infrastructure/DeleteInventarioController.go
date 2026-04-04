// DeleteInventarioController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/application"
)

type DeleteInventarioController struct {
	deleteUseCase *application.DeleteInventarioUseCase
}

func NewDeleteInventarioController(deleteUseCase *application.DeleteInventarioUseCase) *DeleteInventarioController {
	return &DeleteInventarioController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteInventarioController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el registro del inventario",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Producto eliminado del inventario exitosamente",
	})
}