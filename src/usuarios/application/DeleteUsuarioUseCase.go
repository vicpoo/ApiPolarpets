// DeleteUsuarioUseCase.go
package application

import repositories "github.com/vicpoo/ApiPolarpets/src/usuarios/domain"

type DeleteUsuarioUseCase struct {
	repo repositories.IUsuario
}

func NewDeleteUsuarioUseCase(repo repositories.IUsuario) *DeleteUsuarioUseCase {
	return &DeleteUsuarioUseCase{repo: repo}
}

func (uc *DeleteUsuarioUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}