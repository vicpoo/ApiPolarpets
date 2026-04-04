// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiPolarpets/src/user_retos/application"
)

func InitUserRetosDependencies() (
	*CreateUserRetoController,
	*UpdateUserRetoController,
	*DeleteUserRetoController,
	*GetUserRetoByIdController,
	*GetAllUserRetosController,
	*GetUserRetosByUserController,
	*GetUserRetosByRetoController,
	*GetUserRetoByUserAndRetoController,
	*GetCompletedRetosByUserController,
	*GetPendingRetosByUserController,
	*CompleteRetoController,
	*GetUserRetosConDetallesController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLUserRetosRepository()

	// Casos de uso
	createUseCase := application.NewCreateUserRetoUseCase(repo)
	updateUseCase := application.NewUpdateUserRetoUseCase(repo)
	deleteUseCase := application.NewDeleteUserRetoUseCase(repo)
	getByIdUseCase := application.NewGetUserRetoByIdUseCase(repo)
	getAllUseCase := application.NewGetAllUserRetosUseCase(repo)
	getByUserUseCase := application.NewGetUserRetosByUserUseCase(repo)
	getByRetoUseCase := application.NewGetUserRetosByRetoUseCase(repo)
	getByUserAndRetoUseCase := application.NewGetUserRetoByUserAndRetoUseCase(repo)
	getCompletedRetosUseCase := application.NewGetCompletedRetosByUserUseCase(repo)
	getPendingRetosUseCase := application.NewGetPendingRetosByUserUseCase(repo)
	completeRetoUseCase := application.NewCompleteRetoUseCase(repo)
	getUserRetosConDetallesUseCase := application.NewGetUserRetosConDetallesUseCase(repo)

	// Controladores
	createController := NewCreateUserRetoController(createUseCase)
	updateController := NewUpdateUserRetoController(updateUseCase)
	deleteController := NewDeleteUserRetoController(deleteUseCase)
	getByIdController := NewGetUserRetoByIdController(getByIdUseCase)
	getAllController := NewGetAllUserRetosController(getAllUseCase)
	getByUserController := NewGetUserRetosByUserController(getByUserUseCase)
	getByRetoController := NewGetUserRetosByRetoController(getByRetoUseCase)
	getByUserAndRetoController := NewGetUserRetoByUserAndRetoController(getByUserAndRetoUseCase)
	getCompletedRetosController := NewGetCompletedRetosByUserController(getCompletedRetosUseCase)
	getPendingRetosController := NewGetPendingRetosByUserController(getPendingRetosUseCase)
	completeRetoController := NewCompleteRetoController(completeRetoUseCase)
	getUserRetosConDetallesController := NewGetUserRetosConDetallesController(getUserRetosConDetallesUseCase)

	return createController,
		updateController,
		deleteController,
		getByIdController,
		getAllController,
		getByUserController,
		getByRetoController,
		getByUserAndRetoController,
		getCompletedRetosController,
		getPendingRetosController,
		completeRetoController,
		getUserRetosConDetallesController
}