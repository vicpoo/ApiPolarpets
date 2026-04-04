// UpdateRetoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/retos/domain/entities"
)

type UpdateRetoUseCase struct {
	repo repositories.IRetos
}

func NewUpdateRetoUseCase(repo repositories.IRetos) *UpdateRetoUseCase {
	return &UpdateRetoUseCase{repo: repo}
}

func (uc *UpdateRetoUseCase) Run(reto *entities.Retos) (*entities.Retos, error) {
	err := uc.repo.Update(reto)
	if err != nil {
		return nil, err
	}
	return reto, nil
}