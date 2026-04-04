// CreateSkinController.go
package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPolarpets/src/skins/application"
	"github.com/vicpoo/ApiPolarpets/src/skins/domain/entities"
)

type CreateSkinController struct {
	createUseCase *application.CreateSkinUseCase
}

func NewCreateSkinController(createUseCase *application.CreateSkinUseCase) *CreateSkinController {
	return &CreateSkinController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateSkinController) Run(c *gin.Context) {
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
	_, err := fmt.Sscan(idTipoMascotaStr, &idTipoMascota)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de tipo mascota inválido",
			"error":   err.Error(),
		})
		return
	}

	// Obtener el archivo de imagen
	imageFile, imageHeader, err := c.Request.FormFile("imagen")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "La imagen es requerida",
			"error":   err.Error(),
		})
		return
	}
	defer imageFile.Close()

	// Crear la entidad skin
	skin := entities.NewSkins(idTipoMascota, nombre, "")

	// Ejecutar caso de uso
	createdSkin, err := ctrl.createUseCase.Run(skin, imageFile, imageHeader.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear la skin",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdSkin)
}