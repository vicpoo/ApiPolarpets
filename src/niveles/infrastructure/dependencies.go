// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiPolarpets/src/niveles/application"
)

func InitNivelesDependencies() (
	*CreateNivelController,
	*UpdateNivelController,
	*DeleteNivelController,
	*GetNivelByIdController,
	*GetAllNivelesController,
	*GetNivelByNivelController,
	*GetNivelByExpRequeridaController,
	*GetNextLevelController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLNivelesRepository()

	// Casos de uso
	createUseCase := application.NewCreateNivelUseCase(repo)
	updateUseCase := application.NewUpdateNivelUseCase(repo)
	deleteUseCase := application.NewDeleteNivelUseCase(repo)
	getByIdUseCase := application.NewGetNivelByIdUseCase(repo)
	getAllUseCase := application.NewGetAllNivelesUseCase(repo)
	getByNivelUseCase := application.NewGetNivelByNivelUseCase(repo)
	getByExpRequeridaUseCase := application.NewGetNivelByExpRequeridaUseCase(repo)
	getNextLevelUseCase := application.NewGetNextLevelUseCase(repo)

	// Controladores
	createController := NewCreateNivelController(createUseCase)
	updateController := NewUpdateNivelController(updateUseCase)
	deleteController := NewDeleteNivelController(deleteUseCase)
	getByIdController := NewGetNivelByIdController(getByIdUseCase)
	getAllController := NewGetAllNivelesController(getAllUseCase)
	getByNivelController := NewGetNivelByNivelController(getByNivelUseCase)
	getByExpRequeridaController := NewGetNivelByExpRequeridaController(getByExpRequeridaUseCase)
	getNextLevelController := NewGetNextLevelController(getNextLevelUseCase)

	return createController,
		updateController,
		deleteController,
		getByIdController,
		getAllController,
		getByNivelController,
		getByExpRequeridaController,
		getNextLevelController
}