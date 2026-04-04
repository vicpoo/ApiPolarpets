// DeleteCompraController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/compras/application"
)

type DeleteCompraController struct {
	deleteUseCase *application.DeleteCompraUseCase
}

func NewDeleteCompraController(deleteUseCase *application.DeleteCompraUseCase) *DeleteCompraController {
	return &DeleteCompraController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteCompraController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar la compra",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Compra eliminada exitosamente",
	})
}