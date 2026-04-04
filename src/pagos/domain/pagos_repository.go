// pagos_repository.go
package domain

import (
	"time"

	"github.com/vicpoo/ApiPolarpets/src/pagos/domain/entities"
)

type IPagos interface {
	// CRUD básico
	Save(pago *entities.Pagos) error
	Update(pago *entities.Pagos) error
	Delete(id int32) error
	GetById(id int32) (*entities.Pagos, error)
	GetAll() ([]entities.Pagos, error)
	
	// Métodos adicionales útiles
	GetByUser(idUsuario int32) ([]entities.Pagos, error)
	GetByEstado(estado string) ([]entities.Pagos, error)
	GetByMetodoPago(metodoPago string) ([]entities.Pagos, error)
	GetByFechaRange(fechaInicio, fechaFin time.Time) ([]entities.Pagos, error)
	GetByReferenciaExterna(referencia string) (*entities.Pagos, error)
	GetTotalPagadoByUser(idUsuario int32) (float64, error)
	GetPagosCompletadosByUser(idUsuario int32) ([]entities.Pagos, error)
	UpdateEstado(id int32, estado string) error
}