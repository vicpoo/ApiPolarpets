// GetComprasRecientesByUserController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/compras/application"
)

type GetComprasRecientesByUserController struct {
	getRecientesUseCase *application.GetComprasRecientesByUserUseCase
}

func NewGetComprasRecientesByUserController(getRecientesUseCase *application.GetComprasRecientesByUserUseCase) *GetComprasRecientesByUserController {
	return &GetComprasRecientesByUserController{
		getRecientesUseCase: getRecientesUseCase,
	}
}

func (ctrl *GetComprasRecientesByUserController) Run(c *gin.Context) {
	idParam := c.Param("id_usuario")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	limitParam := c.Query("limit")
	limit := 10 // valor por defecto
	if limitParam != "" {
		limit, err = strconv.Atoi(limitParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Limit inválido",
				"error":   err.Error(),
			})
			return
		}
	}

	compras, err := ctrl.getRecientesUseCase.Run(int32(id), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las compras recientes",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, compras)
}