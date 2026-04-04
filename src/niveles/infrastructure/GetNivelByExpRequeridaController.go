// GetNivelByExpRequeridaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/niveles/application"
)

type GetNivelByExpRequeridaController struct {
	getByExpRequeridaUseCase *application.GetNivelByExpRequeridaUseCase
}

func NewGetNivelByExpRequeridaController(getByExpRequeridaUseCase *application.GetNivelByExpRequeridaUseCase) *GetNivelByExpRequeridaController {
	return &GetNivelByExpRequeridaController{
		getByExpRequeridaUseCase: getByExpRequeridaUseCase,
	}
}

func (ctrl *GetNivelByExpRequeridaController) Run(c *gin.Context) {
	expParam := c.Param("exp")
	exp, err := strconv.Atoi(expParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Experiencia inválida",
			"error":   err.Error(),
		})
		return
	}

	nivel, err := ctrl.getByExpRequeridaUseCase.Run(int32(exp))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el nivel",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nivel)
}