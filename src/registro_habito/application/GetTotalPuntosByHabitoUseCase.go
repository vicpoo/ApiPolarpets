// GetTotalPuntosByHabitoUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"

type GetTotalPuntosByHabitoUseCase struct {
	repo repositories.IRegistroHabito
}

func NewGetTotalPuntosByHabitoUseCase(repo repositories.IRegistroHabito) *GetTotalPuntosByHabitoUseCase {
	return &GetTotalPuntosByHabitoUseCase{repo: repo}
}

func (uc *GetTotalPuntosByHabitoUseCase) Run(idHabito int32) (int32, error) {
	return uc.repo.GetTotalPuntosByHabito(idHabito)
}