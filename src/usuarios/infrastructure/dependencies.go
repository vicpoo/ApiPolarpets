// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiPolarpets/src/usuarios/application"
)

func InitUsuarioDependencies() (
	*RegisterController,
	*LoginController,
	*CreateUsuarioController,
	*UpdateUsuarioController,
	*DeleteUsuarioController,
	*GetUsuarioByIdController,
	*GetAllUsuariosController,
	*GetUsuarioByEmailController,
	*UpdateMascotaActivaController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLUsuarioRepository()

	// Casos de uso
	registerUseCase := application.NewRegisterUseCase(repo)
	loginUseCase := application.NewLoginUseCase(repo)
	createUseCase := application.NewCreateUsuarioUseCase(repo)
	updateUseCase := application.NewUpdateUsuarioUseCase(repo)
	deleteUseCase := application.NewDeleteUsuarioUseCase(repo)
	getByIdUseCase := application.NewGetUsuarioByIdUseCase(repo)
	getAllUseCase := application.NewGetAllUsuariosUseCase(repo)
	getByEmailUseCase := application.NewGetUsuarioByEmailUseCase(repo)
	updateMascotaActivaUseCase := application.NewUpdateMascotaActivaUseCase(repo)

	// Controladores
	registerController := NewRegisterController(registerUseCase)
	loginController := NewLoginController(loginUseCase)
	createController := NewCreateUsuarioController(createUseCase)
	updateController := NewUpdateUsuarioController(updateUseCase)
	deleteController := NewDeleteUsuarioController(deleteUseCase)
	getByIdController := NewGetUsuarioByIdController(getByIdUseCase)
	getAllController := NewGetAllUsuariosController(getAllUseCase)
	getByEmailController := NewGetUsuarioByEmailController(getByEmailUseCase)
	updateMascotaActivaController := NewUpdateMascotaActivaController(updateMascotaActivaUseCase)

	return registerController,
		loginController,
		createController,
		updateController,
		deleteController,
		getByIdController,
		getAllController,
		getByEmailController,
		updateMascotaActivaController
}