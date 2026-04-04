// GetSkinByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/skins/application"
)

type GetSkinByIdController struct {
	getByIdUseCase *application.GetSkinByIdUseCase
}

func NewGetSkinByIdController(getByIdUseCase *application.GetSkinByIdUseCase) *GetSkinByIdController {
	return &GetSkinByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetSkinByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	skin, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la skin",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, skin)
}