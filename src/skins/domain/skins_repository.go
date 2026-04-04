// skins_repository.go
package domain

import (
	"github.com/vicpoo/ApiPolarpets/src/skins/domain/entities"
)

type ISkins interface {
	// CRUD básico
	Save(skin *entities.Skins) error
	Update(skin *entities.Skins) error
	Delete(id int32) error
	GetById(id int32) (*entities.Skins, error)
	GetAll() ([]entities.Skins, error)
	
	// Métodos adicionales útiles
	GetByTipoMascota(idTipoMascota int32) ([]entities.Skins, error)
	GetByNombre(nombre string) (*entities.Skins, error)
	GetByTipoMascotaAndNombre(idTipoMascota int32, nombre string) (*entities.Skins, error)
}