// habito_repository.go
package domain

import (
	"github.com/vicpoo/ApiPolarpets/src/habito/domain/entities"
)

type IHabito interface {
	// CRUD básico
	Save(habito *entities.Habito) error
	Update(habito *entities.Habito) error
	Delete(id int32) error
	GetById(id int32) (*entities.Habito, error)
	GetAll() ([]entities.Habito, error)
	
	// Métodos adicionales útiles
	GetByUser(idUser int32) ([]entities.Habito, error)
	GetByTitulo(titulo string) (*entities.Habito, error)
	GetByUserAndTitulo(idUser int32, titulo string) (*entities.Habito, error)
	GetTotalPuntosByUser(idUser int32) (int32, error)
}