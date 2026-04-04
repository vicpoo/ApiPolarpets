// retos_repository.go
package domain

import (
	"github.com/vicpoo/ApiPolarpets/src/retos/domain/entities"
)

type IRetos interface {
	// CRUD básico
	Save(reto *entities.Retos) error
	Update(reto *entities.Retos) error
	Delete(id int32) error
	GetById(id int32) (*entities.Retos, error)
	GetAll() ([]entities.Retos, error)
	
	// Métodos adicionales útiles
	GetByTitulo(titulo string) (*entities.Retos, error)
	GetByPuntosRange(minPuntos, maxPuntos int32) ([]entities.Retos, error)
}