// DeleteNivelUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/niveles/domain"

type DeleteNivelUseCase struct {
	repo repositories.INiveles
}

func NewDeleteNivelUseCase(repo repositories.INiveles) *DeleteNivelUseCase {
	return &DeleteNivelUseCase{repo: repo}
}

func (uc *DeleteNivelUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}