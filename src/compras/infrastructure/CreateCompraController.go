// CreateCompraController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/compras/application"
	"github.com/vicpoo/ApiPolarpets/src/compras/domain/entities"
)

type CreateCompraController struct {
	createUseCase *application.CreateCompraUseCase
}

func NewCreateCompraController(createUseCase *application.CreateCompraUseCase) *CreateCompraController {
	return &CreateCompraController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateCompraController) Run(c *gin.Context) {
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

	createdCompra, err := ctrl.createUseCase.Run(compra)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo registrar la compra",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdCompra)
}