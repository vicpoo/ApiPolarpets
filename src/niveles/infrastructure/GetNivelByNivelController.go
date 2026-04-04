// GetNivelByNivelController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/niveles/application"
)

type GetNivelByNivelController struct {
	getByNivelUseCase *application.GetNivelByNivelUseCase
}

func NewGetNivelByNivelController(getByNivelUseCase *application.GetNivelByNivelUseCase) *GetNivelByNivelController {
	return &GetNivelByNivelController{
		getByNivelUseCase: getByNivelUseCase,
	}
}

func (ctrl *GetNivelByNivelController) Run(c *gin.Context) {
	nivelParam := c.Param("nivel")
	nivel, err := strconv.Atoi(nivelParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nivel inválido",
			"error":   err.Error(),
		})
		return
	}

	nivelInfo, err := ctrl.getByNivelUseCase.Run(int32(nivel))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el nivel",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nivelInfo)
}