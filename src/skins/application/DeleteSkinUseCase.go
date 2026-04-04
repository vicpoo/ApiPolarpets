// DeleteSkinUseCase.go
package application

import (
	"os"

	repositories "github.com/vicpoo/ApiPolarpets/src/skins/domain"
)

type DeleteSkinUseCase struct {
	repo repositories.ISkins
}

func NewDeleteSkinUseCase(repo repositories.ISkins) *DeleteSkinUseCase {
	return &DeleteSkinUseCase{repo: repo}
}

func (uc *DeleteSkinUseCase) Run(id int32) error {
	// Obtener la skin para eliminar también su imagen
	skin, err := uc.repo.GetById(id)
	if err != nil {
		return err
	}

	// Eliminar la imagen asociada si existe
	if skin.GetImagenURL() != "" {
		imagePath := "." + skin.GetImagenURL()
		if err := os.Remove(imagePath); err != nil {
			// Solo loguear el error, no fallar la eliminación
			println("Error al eliminar imagen:", err.Error())
		}
	}

	return uc.repo.Delete(id)
}