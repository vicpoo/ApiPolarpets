// GetMascotasByUserController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/application"
)

type GetMascotasByUserController struct {
	getByUserUseCase *application.GetMascotasByUserUseCase
}

func NewGetMascotasByUserController(getByUserUseCase *application.GetMascotasByUserUseCase) *GetMascotasByUserController {
	return &GetMascotasByUserController{
		getByUserUseCase: getByUserUseCase,
	}
}

func (ctrl *GetMascotasByUserController) Run(c *gin.Context) {
	idParam := c.Param("id_user")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	mascotas, err := ctrl.getByUserUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las mascotas del usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, mascotas)
}