// UpdateRetoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/retos/application"
	"github.com/vicpoo/ApiPolarpets/src/retos/domain/entities"
)

type UpdateRetoController struct {
	updateUseCase *application.UpdateRetoUseCase
}

func NewUpdateRetoController(updateUseCase *application.UpdateRetoUseCase) *UpdateRetoController {
	return &UpdateRetoController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateRetoController) Run(c *gin.Context) {
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
		Titulo         string `json:"titulo"`
		Descripcion    string `json:"descripcion"`
		PuntosGenerados int32  `json:"puntos_generados"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	reto := entities.NewRetos(request.Titulo, request.Descripcion, request.PuntosGenerados)
	reto.SetIDRetos(int32(id))

	updatedReto, err := ctrl.updateUseCase.Run(reto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el reto",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedReto)
}