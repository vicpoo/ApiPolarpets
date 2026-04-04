// GetMascotaByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/application"
)

type GetMascotaByIdController struct {
	getByIdUseCase *application.GetMascotaByIdUseCase
}

func NewGetMascotaByIdController(getByIdUseCase *application.GetMascotaByIdUseCase) *GetMascotaByIdController {
	return &GetMascotaByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetMascotaByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	mascota, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la mascota",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, mascota)
}