// UpdateInventarioController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/application"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain/entities"
)

type UpdateInventarioController struct {
	updateUseCase *application.UpdateInventarioUseCase
}

func NewUpdateInventarioController(updateUseCase *application.UpdateInventarioUseCase) *UpdateInventarioController {
	return &UpdateInventarioController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateInventarioController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var request struct {
		IDUsuario  int32 `json:"id_usuario"`
		IDProducto int32 `json:"id_producto"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	inventario := entities.NewInventarioUsuario(request.IDUsuario, request.IDProducto)
	inventario.SetIDInventario(int32(id))

	updatedInventario, err := ctrl.updateUseCase.Run(inventario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el inventario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedInventario)
}