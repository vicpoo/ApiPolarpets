// ExistsRegistroHoyUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"

type ExistsRegistroHoyUseCase struct {
	repo repositories.IRegistroHabito
}

func NewExistsRegistroHoyUseCase(repo repositories.IRegistroHabito) *ExistsRegistroHoyUseCase {
	return &ExistsRegistroHoyUseCase{repo: repo}
}

func (uc *ExistsRegistroHoyUseCase) Run(idHabito int32) (bool, error) {
	return uc.repo.ExistsRegistroHoy(idHabito)
}