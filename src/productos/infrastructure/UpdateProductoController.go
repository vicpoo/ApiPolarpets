// UpdateProductoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/productos/application"
	"github.com/vicpoo/ApiPolarpets/src/productos/domain/entities"
)

type UpdateProductoController struct {
	updateUseCase *application.UpdateProductoUseCase
}

func NewUpdateProductoController(updateUseCase *application.UpdateProductoUseCase) *UpdateProductoController {
	return &UpdateProductoController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateProductoController) Run(c *gin.Context) {
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
		Nombre        string  `json:"nombre"`
		Descripcion   string  `json:"descripcion"`
		Tipo          string  `json:"tipo"`
		Precio        float64 `json:"precio"`
		IDSkin        *int32  `json:"id_skin,omitempty"`
		IDTipoMascota *int32  `json:"id_tipo_mascota,omitempty"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	producto := entities.NewProductos(
		request.Nombre,
		request.Descripcion,
		request.Tipo,
		request.Precio,
		request.IDSkin,
		request.IDTipoMascota,
	)
	producto.SetIDProducto(int32(id))

	updatedProducto, err := ctrl.updateUseCase.Run(producto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el producto",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedProducto)
}