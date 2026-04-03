// rol_repository.go
package domain

import (
	"github.com/vicpoo/ApiPolarpets/src/roles/domain/entities"
)

type IRol interface {
	Save(rol *entities.Rol) error
	Update(rol *entities.Rol) error
	Delete(id int32) error
	GetById(id int32) (*entities.Rol, error)
	GetAll() ([]entities.Rol, error)
}