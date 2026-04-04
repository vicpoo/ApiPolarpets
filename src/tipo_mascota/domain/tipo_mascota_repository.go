// tipo_mascota_repository.go
package domain

import (
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain/entities"
)

type ITipoMascota interface {
	// CRUD básico
	Save(tipoMascota *entities.TipoMascota) error
	Update(tipoMascota *entities.TipoMascota) error
	Delete(id int32) error
	GetById(id int32) (*entities.TipoMascota, error)
	GetAll() ([]entities.TipoMascota, error)
	
	// Métodos adicionales útiles
	GetByNombre(nombre string) (*entities.TipoMascota, error)
}