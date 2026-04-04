// GetNivelByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/niveles/application"
)

type GetNivelByIdController struct {
	getByIdUseCase *application.GetNivelByIdUseCase
}

func NewGetNivelByIdController(getByIdUseCase *application.GetNivelByIdUseCase) *GetNivelByIdController {
	return &GetNivelByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetNivelByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	nivel, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el nivel",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nivel)
}