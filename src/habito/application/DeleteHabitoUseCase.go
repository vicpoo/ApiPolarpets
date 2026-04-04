// DeleteHabitoUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/habito/domain"

type DeleteHabitoUseCase struct {
	repo repositories.IHabito
}

func NewDeleteHabitoUseCase(repo repositories.IHabito) *DeleteHabitoUseCase {
	return &DeleteHabitoUseCase{repo: repo}
}

func (uc *DeleteHabitoUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}