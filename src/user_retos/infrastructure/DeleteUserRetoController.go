// DeleteUserRetoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/application"
)

type DeleteUserRetoController struct {
	deleteUseCase *application.DeleteUserRetoUseCase
}

func NewDeleteUserRetoController(deleteUseCase *application.DeleteUserRetoUseCase) *DeleteUserRetoController {
	return &DeleteUserRetoController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteUserRetoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar el registro",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Registro eliminado exitosamente",
	})
}