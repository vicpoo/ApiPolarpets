// UpdateCompraController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/compras/application"
	"github.com/vicpoo/ApiPolarpets/src/compras/domain/entities"
)

type UpdateCompraController struct {
	updateUseCase *application.UpdateCompraUseCase
}

func NewUpdateCompraController(updateUseCase *application.UpdateCompraUseCase) *UpdateCompraController {
	return &UpdateCompraController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateCompraController) Run(c *gin.Context) {
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
		IDPago     int32 `json:"id_pago"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	compra := entities.NewCompras(request.IDUsuario, request.IDProducto, request.IDPago)
	compra.SetIDCompra(int32(id))

	updatedCompra, err := ctrl.updateUseCase.Run(compra)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la compra",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedCompra)
}