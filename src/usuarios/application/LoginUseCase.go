// LoginUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiPolarpets/src/usuarios/domain"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/domain/entities"
	"github.com/vicpoo/ApiPolarpets/src/core"
)

type LoginResponse struct {
	Token   string            `json:"token"`
	Usuario *entities.Usuario `json:"usuario"`
}

type LoginUseCase struct {
	repo repositories.IUsuario
}

func NewLoginUseCase(repo repositories.IUsuario) *LoginUseCase {
	return &LoginUseCase{repo: repo}
}

func (uc *LoginUseCase) Run(email, password string) (*LoginResponse, error) {
	// Autenticar usuario
	usuario, err := uc.repo.Login(email, password)
	if err != nil {
		return nil, err
	}

	// Generar token JWT
	token, err := core.GenerarToken(
		usuario.GetIDUsuario(),
		usuario.GetEmail(),
		usuario.GetUsername(),
		usuario.GetIDRol(),
	)
	if err != nil {
		return nil, err
	}

	// Limpiar contraseña
	usuario.SetPassword("")

	return &LoginResponse{
		Token:   token,
		Usuario: usuario,
	}, nil
}