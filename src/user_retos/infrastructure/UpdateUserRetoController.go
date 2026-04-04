// UpdateUserRetoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/application"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/domain/entities"
)

type UpdateUserRetoController struct {
	updateUseCase *application.UpdateUserRetoUseCase
}

func NewUpdateUserRetoController(updateUseCase *application.UpdateUserRetoUseCase) *UpdateUserRetoController {
	return &UpdateUserRetoController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateUserRetoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var request struct {
		IDUsuario int32 `json:"id_usuario"`
		IDReto    int32 `json:"id_reto"`
		Completo  bool  `json:"completo"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	userReto := entities.NewUserRetos(request.IDUsuario, request.IDReto, request.Completo)
	userReto.SetIDUserRetos(int32(id))

	updatedUserReto, err := ctrl.updateUseCase.Run(userReto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el registro",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedUserReto)
}