// GetSkinByNombreUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/skins/domain"
	"github.com/vicpoo/ApiPolarpets/src/skins/domain/entities"
)

type GetSkinByNombreUseCase struct {
	repo repositories.ISkins
}

func NewGetSkinByNombreUseCase(repo repositories.ISkins) *GetSkinByNombreUseCase {
	return &GetSkinByNombreUseCase{repo: repo}
}

func (uc *GetSkinByNombreUseCase) Run(nombre string) (*entities.Skins, error) {
	return uc.repo.GetByNombre(nombre)
}