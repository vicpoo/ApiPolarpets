// CreateRetoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/retos/domain/entities"
)

type CreateRetoUseCase struct {
	repo repositories.IRetos
}

func NewCreateRetoUseCase(repo repositories.IRetos) *CreateRetoUseCase {
	return &CreateRetoUseCase{repo: repo}
}

func (uc *CreateRetoUseCase) Run(reto *entities.Retos) (*entities.Retos, error) {
	err := uc.repo.Save(reto)
	if err != nil {
		return nil, err
	}
	return reto, nil
}