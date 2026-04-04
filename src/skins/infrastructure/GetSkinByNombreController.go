// GetSkinByNombreController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/skins/application"
)

type GetSkinByNombreController struct {
	getByNombreUseCase *application.GetSkinByNombreUseCase
}

func NewGetSkinByNombreController(getByNombreUseCase *application.GetSkinByNombreUseCase) *GetSkinByNombreController {
	return &GetSkinByNombreController{
		getByNombreUseCase: getByNombreUseCase,
	}
}

func (ctrl *GetSkinByNombreController) Run(c *gin.Context) {
	nombre := c.Query("nombre")
	if nombre == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nombre es requerido",
		})
		return
	}

	skin, err := ctrl.getByNombreUseCase.Run(nombre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la skin",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, skin)
}