// GetSkinByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/skins/domain"
	"github.com/vicpoo/ApiPolarpets/src/skins/domain/entities"
)

type GetSkinByIdUseCase struct {
	repo repositories.ISkins
}

func NewGetSkinByIdUseCase(repo repositories.ISkins) *GetSkinByIdUseCase {
	return &GetSkinByIdUseCase{repo: repo}
}

func (uc *GetSkinByIdUseCase) Run(id int32) (*entities.Skins, error) {
	return uc.repo.GetById(id)
}