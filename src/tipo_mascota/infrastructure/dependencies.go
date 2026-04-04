// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/application"
)

func InitTipoMascotaDependencies() (
	*CreateTipoMascotaController,
	*UpdateTipoMascotaController,
	*DeleteTipoMascotaController,
	*GetTipoMascotaByIdController,
	*GetAllTiposMascotaController,
	*GetTipoMascotaByNombreController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLTipoMascotaRepository()

	// Casos de uso
	createUseCase := application.NewCreateTipoMascotaUseCase(repo)
	updateUseCase := application.NewUpdateTipoMascotaUseCase(repo)
	deleteUseCase := application.NewDeleteTipoMascotaUseCase(repo)
	getByIdUseCase := application.NewGetTipoMascotaByIdUseCase(repo)
	getAllUseCase := application.NewGetAllTiposMascotaUseCase(repo)
	getByNombreUseCase := application.NewGetTipoMascotaByNombreUseCase(repo)

	// Controladores
	createController := NewCreateTipoMascotaController(createUseCase)
	updateController := NewUpdateTipoMascotaController(updateUseCase)
	deleteController := NewDeleteTipoMascotaController(deleteUseCase)
	getByIdController := NewGetTipoMascotaByIdController(getByIdUseCase)
	getAllController := NewGetAllTiposMascotaController(getAllUseCase)
	getByNombreController := NewGetTipoMascotaByNombreController(getByNombreUseCase)

	return createController,
		updateController,
		deleteController,
		getByIdController,
		getAllController,
		getByNombreController
}