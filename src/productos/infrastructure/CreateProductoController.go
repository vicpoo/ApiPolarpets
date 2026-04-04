// CreateProductoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/productos/application"
	"github.com/vicpoo/ApiPolarpets/src/productos/domain/entities"
)

type CreateProductoController struct {
	createUseCase *application.CreateProductoUseCase
}

func NewCreateProductoController(createUseCase *application.CreateProductoUseCase) *CreateProductoController {
	return &CreateProductoController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateProductoController) Run(c *gin.Context) {
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

	createdProducto, err := ctrl.createUseCase.Run(producto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el producto",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdProducto)
}