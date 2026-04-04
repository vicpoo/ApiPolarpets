// UpdateSkinUseCase.go
package application

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	repositories "github.com/vicpoo/ApiPolarpets/src/skins/domain"
	"github.com/vicpoo/ApiPolarpets/src/skins/domain/entities"
)

type UpdateSkinUseCase struct {
	repo repositories.ISkins
}

func NewUpdateSkinUseCase(repo repositories.ISkins) *UpdateSkinUseCase {
	return &UpdateSkinUseCase{repo: repo}
}

func (uc *UpdateSkinUseCase) Run(skin *entities.Skins, imageFile io.Reader, fileName string) (*entities.Skins, error) {
	// Obtener la skin actual para verificar si tiene imagen
	currentSkin, err := uc.repo.GetById(skin.GetIDSkins())
	if err != nil {
		return nil, err
	}

	// Si se sube una nueva imagen
	if imageFile != nil && fileName != "" {
		// Eliminar la imagen anterior si existe
		if currentSkin.GetImagenURL() != "" {
			oldImagePath := "." + currentSkin.GetImagenURL()
			if err := os.Remove(oldImagePath); err != nil {
				// Solo loguear el error, no fallar la actualización
				println("Error al eliminar imagen anterior:", err.Error())
			}
		}

		// Guardar la nueva imagen
		skinsDir := "./skins"
		if err := os.MkdirAll(skinsDir, 0755); err != nil {
			return nil, err
		}

		ext := filepath.Ext(fileName)
		cleanName := strings.ToLower(strings.ReplaceAll(skin.GetNombre(), " ", "_"))
		uniqueFileName := cleanName + ext
		filePath := filepath.Join(skinsDir, uniqueFileName)

		dst, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, imageFile); err != nil {
			return nil, err
		}

		imageURL := "/skins/" + uniqueFileName
		skin.SetImagenURL(imageURL)
	} else {
		// Mantener la imagen actual si no se sube una nueva
		skin.SetImagenURL(currentSkin.GetImagenURL())
	}

	err = uc.repo.Update(skin)
	if err != nil {
		return nil, err
	}
	return skin, nil
}