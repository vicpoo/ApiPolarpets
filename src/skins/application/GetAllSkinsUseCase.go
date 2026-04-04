// GetAllSkinsUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/skins/domain"
	"github.com/vicpoo/ApiPolarpets/src/skins/domain/entities"
)

type GetAllSkinsUseCase struct {
	repo repositories.ISkins
}

func NewGetAllSkinsUseCase(repo repositories.ISkins) *GetAllSkinsUseCase {
	return &GetAllSkinsUseCase{repo: repo}
}

func (uc *GetAllSkinsUseCase) Run() ([]entities.Skins, error) {
	return uc.repo.GetAll()
}