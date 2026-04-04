// DeleteRetoUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/retos/domain"

type DeleteRetoUseCase struct {
	repo repositories.IRetos
}

func NewDeleteRetoUseCase(repo repositories.IRetos) *DeleteRetoUseCase {
	return &DeleteRetoUseCase{repo: repo}
}

func (uc *DeleteRetoUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}