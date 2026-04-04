// GetNextLevelController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/niveles/application"
)

type GetNextLevelController struct {
	getNextLevelUseCase *application.GetNextLevelUseCase
}

func NewGetNextLevelController(getNextLevelUseCase *application.GetNextLevelUseCase) *GetNextLevelController {
	return &GetNextLevelController{
		getNextLevelUseCase: getNextLevelUseCase,
	}
}

func (ctrl *GetNextLevelController) Run(c *gin.Context) {
	expParam := c.Query("exp_actual")
	if expParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Experiencia actual es requerida",
		})
		return
	}

	expActual, err := strconv.Atoi(expParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Experiencia inválida",
			"error":   err.Error(),
		})
		return
	}

	nextLevel, err := ctrl.getNextLevelUseCase.Run(int32(expActual))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el siguiente nivel",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nextLevel)
}