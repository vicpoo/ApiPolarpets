// DeleteRegistroHabitoUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"

type DeleteRegistroHabitoUseCase struct {
	repo repositories.IRegistroHabito
}

func NewDeleteRegistroHabitoUseCase(repo repositories.IRegistroHabito) *DeleteRegistroHabitoUseCase {
	return &DeleteRegistroHabitoUseCase{repo: repo}
}

func (uc *DeleteRegistroHabitoUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}