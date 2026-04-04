// GetHabitoByUserAndTituloController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/habito/application"
)

type GetHabitoByUserAndTituloController struct {
	getByUserAndTituloUseCase *application.GetHabitoByUserAndTituloUseCase
}

func NewGetHabitoByUserAndTituloController(getByUserAndTituloUseCase *application.GetHabitoByUserAndTituloUseCase) *GetHabitoByUserAndTituloController {
	return &GetHabitoByUserAndTituloController{
		getByUserAndTituloUseCase: getByUserAndTituloUseCase,
	}
}

func (ctrl *GetHabitoByUserAndTituloController) Run(c *gin.Context) {
	idUserParam := c.Query("id_user")
	titulo := c.Query("titulo")

	if idUserParam == "" || titulo == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id_user y título son requeridos",
		})
		return
	}

	idUser, err := strconv.Atoi(idUserParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	habito, err := ctrl.getByUserAndTituloUseCase.Run(int32(idUser), titulo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el hábito",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, habito)
}