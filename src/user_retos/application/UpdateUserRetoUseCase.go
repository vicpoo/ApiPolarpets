// UpdateUserRetoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/user_retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/domain/entities"
)

type UpdateUserRetoUseCase struct {
	repo repositories.IUserRetos
}

func NewUpdateUserRetoUseCase(repo repositories.IUserRetos) *UpdateUserRetoUseCase {
	return &UpdateUserRetoUseCase{repo: repo}
}

func (uc *UpdateUserRetoUseCase) Run(userReto *entities.UserRetos) (*entities.UserRetos, error) {
	err := uc.repo.Update(userReto)
	if err != nil {
		return nil, err
	}
	return userReto, nil
}