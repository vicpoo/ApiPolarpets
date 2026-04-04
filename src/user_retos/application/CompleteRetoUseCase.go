// CompleteRetoUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/user_retos/domain"

type CompleteRetoUseCase struct {
	repo repositories.IUserRetos
}

func NewCompleteRetoUseCase(repo repositories.IUserRetos) *CompleteRetoUseCase {
	return &CompleteRetoUseCase{repo: repo}
}

func (uc *CompleteRetoUseCase) Run(idUsuario, idReto int32) error {
	return uc.repo.CompleteReto(idUsuario, idReto)
}