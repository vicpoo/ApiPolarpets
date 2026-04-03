// UpdateRolUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/roles/domain"
	"github.com/vicpoo/ApiPolarpets/src/roles/domain/entities"
)

type UpdateRolUseCase struct {
	repo repositories.IRol
}

func NewUpdateRolUseCase(repo repositories.IRol) *UpdateRolUseCase {
	return &UpdateRolUseCase{repo: repo}
}

func (uc *UpdateRolUseCase) Run(rol *entities.Rol) (*entities.Rol, error) {
	err := uc.repo.Update(rol)
	if err != nil {
		return nil, err
	}
	return rol, nil
}