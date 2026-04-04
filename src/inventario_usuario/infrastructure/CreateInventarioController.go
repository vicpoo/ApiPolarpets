// CreateInventarioController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/application"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain/entities"
)

type CreateInventarioController struct {
	createUseCase *application.CreateInventarioUseCase
}

func NewCreateInventarioController(createUseCase *application.CreateInventarioUseCase) *CreateInventarioController {
	return &CreateInventarioController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateInventarioController) Run(c *gin.Context) {
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

	createdInventario, err := ctrl.createUseCase.Run(inventario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo agregar el producto al inventario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdInventario)
}