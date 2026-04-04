// DeleteUserRetoUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/user_retos/domain"

type DeleteUserRetoUseCase struct {
	repo repositories.IUserRetos
}

func NewDeleteUserRetoUseCase(repo repositories.IUserRetos) *DeleteUserRetoUseCase {
	return &DeleteUserRetoUseCase{repo: repo}
}

func (uc *DeleteUserRetoUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}