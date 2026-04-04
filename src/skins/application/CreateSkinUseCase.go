// CreateSkinUseCase.go
package application

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	repositories "github.com/vicpoo/ApiPolarpets/src/skins/domain"
	"github.com/vicpoo/ApiPolarpets/src/skins/domain/entities"
)

type CreateSkinUseCase struct {
	repo repositories.ISkins
}

func NewCreateSkinUseCase(repo repositories.ISkins) *CreateSkinUseCase {
	return &CreateSkinUseCase{repo: repo}
}

func (uc *CreateSkinUseCase) Run(skin *entities.Skins, imageFile io.Reader, fileName string) (*entities.Skins, error) {
	// Guardar la imagen en la carpeta /skins
	if imageFile != nil && fileName != "" {
		// Crear la carpeta si no existe
		skinsDir := "./skins"
		if err := os.MkdirAll(skinsDir, 0755); err != nil {
			return nil, err
		}

		// Limpiar el nombre del archivo y generar nombre único
		ext := filepath.Ext(fileName)
		cleanName := strings.ToLower(strings.ReplaceAll(skin.GetNombre(), " ", "_"))
		uniqueFileName := cleanName + ext
		filePath := filepath.Join(skinsDir, uniqueFileName)

		// Crear el archivo
		dst, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		defer dst.Close()

		// Copiar el contenido de la imagen
		if _, err := io.Copy(dst, imageFile); err != nil {
			return nil, err
		}

		// Guardar la URL en la entidad
		imageURL := "/skins/" + uniqueFileName
		skin.SetImagenURL(imageURL)
	}

	err := uc.repo.Save(skin)
	if err != nil {
		return nil, err
	}
	return skin, nil
}