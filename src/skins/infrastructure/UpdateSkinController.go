// UpdateSkinController.go
package infrastructure

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/skins/application"
	"github.com/vicpoo/ApiPolarpets/src/skins/domain/entities"
)

type UpdateSkinController struct {
	updateUseCase *application.UpdateSkinUseCase
}

func NewUpdateSkinController(updateUseCase *application.UpdateSkinUseCase) *UpdateSkinController {
	return &UpdateSkinController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateSkinController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	// Obtener datos del formulario
	nombre := c.PostForm("nombre")
	idTipoMascotaStr := c.PostForm("id_tipo_mascota")

	// Validar campos requeridos
	if nombre == "" || idTipoMascotaStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Los campos nombre y id_tipo_mascota son requeridos",
		})
		return
	}

	// Convertir id_tipo_mascota a int32
	var idTipoMascota int32
	_, err = fmt.Sscan(idTipoMascotaStr, &idTipoMascota)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de tipo mascota inválido",
			"error":   err.Error(),
		})
		return
	}

	// Crear la entidad skin
	skin := entities.NewSkins(idTipoMascota, nombre, "")
	skin.SetIDSkins(int32(id))

	// Verificar si se subió una nueva imagen
	var imageFile io.Reader
	var imageFileName string
	
	file, imageHeader, err := c.Request.FormFile("imagen")
	if err == nil {
		defer file.Close()
		imageFile = file
		imageFileName = imageHeader.Filename
	}

	// Ejecutar caso de uso
	updatedSkin, err := ctrl.updateUseCase.Run(skin, imageFile, imageFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la skin",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedSkin)
}