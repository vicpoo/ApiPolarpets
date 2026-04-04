// niveles_repository.go
package domain

import (
	"github.com/vicpoo/ApiPolarpets/src/niveles/domain/entities"
)

type INiveles interface {
	// CRUD básico
	Save(nivel *entities.Niveles) error
	Update(nivel *entities.Niveles) error
	Delete(id int32) error
	GetById(id int32) (*entities.Niveles, error)
	GetAll() ([]entities.Niveles, error)
	
	// Métodos adicionales útiles
	GetByNivel(nivel int32) (*entities.Niveles, error)
	GetByExpRequerida(expRequerida int32) (*entities.Niveles, error)
	GetNextLevel(currentExp int32) (*entities.Niveles, error)
}