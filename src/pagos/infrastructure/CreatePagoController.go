// CreatePagoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/pagos/application"
	"github.com/vicpoo/ApiPolarpets/src/pagos/domain/entities"
)

type CreatePagoController struct {
	createUseCase *application.CreatePagoUseCase
}

func NewCreatePagoController(createUseCase *application.CreatePagoUseCase) *CreatePagoController {
	return &CreatePagoController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreatePagoController) Run(c *gin.Context) {
	var request struct {
		IDUsuario         int32   `json:"id_usuario"`
		Monto             float64 `json:"monto"`
		MetodoPago        string  `json:"metodo_pago"`
		Estado            string  `json:"estado"`
		ReferenciaExterna string  `json:"referencia_externa"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	pago := entities.NewPagos(
		request.IDUsuario,
		request.Monto,
		request.MetodoPago,
		request.Estado,
		request.ReferenciaExterna,
	)

	createdPago, err := ctrl.createUseCase.Run(pago)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo registrar el pago",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdPago)
}