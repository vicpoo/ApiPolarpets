// GetUserRetoByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/application"
)

type GetUserRetoByIdController struct {
	getByIdUseCase *application.GetUserRetoByIdUseCase
}

func NewGetUserRetoByIdController(getByIdUseCase *application.GetUserRetoByIdUseCase) *GetUserRetoByIdController {
	return &GetUserRetoByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetUserRetoByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	userReto, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el registro",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userReto)
}	