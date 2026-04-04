// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiPolarpets/src/skins/application"
)

func InitSkinsDependencies() (
	*CreateSkinController,
	*UpdateSkinController,
	*DeleteSkinController,
	*GetSkinByIdController,
	*GetAllSkinsController,
	*GetSkinsByTipoMascotaController,
	*GetSkinByNombreController,
	*GetSkinByTipoMascotaAndNombreController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLSkinsRepository()

	// Casos de uso
	createUseCase := application.NewCreateSkinUseCase(repo)
	updateUseCase := application.NewUpdateSkinUseCase(repo)
	deleteUseCase := application.NewDeleteSkinUseCase(repo)
	getByIdUseCase := application.NewGetSkinByIdUseCase(repo)
	getAllUseCase := application.NewGetAllSkinsUseCase(repo)
	getByTipoMascotaUseCase := application.NewGetSkinsByTipoMascotaUseCase(repo)
	getByNombreUseCase := application.NewGetSkinByNombreUseCase(repo)
	getByTipoMascotaAndNombreUseCase := application.NewGetSkinByTipoMascotaAndNombreUseCase(repo)

	// Controladores
	createController := NewCreateSkinController(createUseCase)
	updateController := NewUpdateSkinController(updateUseCase)
	deleteController := NewDeleteSkinController(deleteUseCase)
	getByIdController := NewGetSkinByIdController(getByIdUseCase)
	getAllController := NewGetAllSkinsController(getAllUseCase)
	getByTipoMascotaController := NewGetSkinsByTipoMascotaController(getByTipoMascotaUseCase)
	getByNombreController := NewGetSkinByNombreController(getByNombreUseCase)
	getByTipoMascotaAndNombreController := NewGetSkinByTipoMascotaAndNombreController(getByTipoMascotaAndNombreUseCase)

	return createController,
		updateController,
		deleteController,
		getByIdController,
		getAllController,
		getByTipoMascotaController,
		getByNombreController,
		getByTipoMascotaAndNombreController
}