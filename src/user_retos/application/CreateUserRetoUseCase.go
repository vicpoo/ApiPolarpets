// CreateUserRetoUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/user_retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/domain/entities"
)

type CreateUserRetoUseCase struct {
	repo repositories.IUserRetos
}

func NewCreateUserRetoUseCase(repo repositories.IUserRetos) *CreateUserRetoUseCase {
	return &CreateUserRetoUseCase{repo: repo}
}

func (uc *CreateUserRetoUseCase) Run(userReto *entities.UserRetos) (*entities.UserRetos, error) {
	err := uc.repo.Save(userReto)
	if err != nil {
		return nil, err
	}
	return userReto, nil
}