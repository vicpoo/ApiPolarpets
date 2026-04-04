// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiPolarpets/src/mascotas/application"
)

func InitMascotasDependencies() (
	*CreateMascotaController,
	*UpdateMascotaController,
	*DeleteMascotaController,
	*GetMascotaByIdController,
	*GetAllMascotasController,
	*GetMascotasByUserController,
	*GetMascotasByTipoMascotaController,
	*GetMascotasBySkinController,
	*GetMascotasByNivelController,
	*UpdateExperienciaController,
	*GetMascotaCompletaController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLMascotaRepository()

	// Casos de uso
	createUseCase := application.NewCreateMascotaUseCase(repo)
	updateUseCase := application.NewUpdateMascotaUseCase(repo)
	deleteUseCase := application.NewDeleteMascotaUseCase(repo)
	getByIdUseCase := application.NewGetMascotaByIdUseCase(repo)
	getAllUseCase := application.NewGetAllMascotasUseCase(repo)
	getByUserUseCase := application.NewGetMascotasByUserUseCase(repo)
	getByTipoMascotaUseCase := application.NewGetMascotasByTipoMascotaUseCase(repo)
	getBySkinUseCase := application.NewGetMascotasBySkinUseCase(repo)
	getByNivelUseCase := application.NewGetMascotasByNivelUseCase(repo)
	updateExperienciaUseCase := application.NewUpdateExperienciaUseCase(repo)
	getMascotaCompletaUseCase := application.NewGetMascotaCompletaUseCase(repo)

	// Controladores
	createController := NewCreateMascotaController(createUseCase)
	updateController := NewUpdateMascotaController(updateUseCase)
	deleteController := NewDeleteMascotaController(deleteUseCase)
	getByIdController := NewGetMascotaByIdController(getByIdUseCase)
	getAllController := NewGetAllMascotasController(getAllUseCase)
	getByUserController := NewGetMascotasByUserController(getByUserUseCase)
	getByTipoMascotaController := NewGetMascotasByTipoMascotaController(getByTipoMascotaUseCase)
	getBySkinController := NewGetMascotasBySkinController(getBySkinUseCase)
	getByNivelController := NewGetMascotasByNivelController(getByNivelUseCase)
	updateExperienciaController := NewUpdateExperienciaController(updateExperienciaUseCase)
	getMascotaCompletaController := NewGetMascotaCompletaController(getMascotaCompletaUseCase)

	return createController,
		updateController,
		deleteController,
		getByIdController,
		getAllController,
		getByUserController,
		getByTipoMascotaController,
		getBySkinController,
		getByNivelController,
		updateExperienciaController,
		getMascotaCompletaController
}