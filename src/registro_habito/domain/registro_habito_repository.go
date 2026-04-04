// registro_habito_repository.go - ACTUALIZADO
package domain

import (
	"time"

	"github.com/vicpoo/ApiPolarpets/src/registro_habito/domain/entities"
)

type IRegistroHabito interface {
	// CRUD básico
	Save(registro *entities.RegistroHabito) error
	Update(registro *entities.RegistroHabito) error
	Delete(id int32) error
	GetById(id int32) (*entities.RegistroHabito, error)
	GetAll() ([]entities.RegistroHabito, error)
	
	// Métodos adicionales útiles
	GetByHabito(idHabito int32) ([]entities.RegistroHabito, error)
	GetByHabitoAndFecha(idHabito int32, fecha time.Time) (*entities.RegistroHabito, error)
	GetByFechaRange(fechaInicio, fechaFin time.Time) ([]entities.RegistroHabito, error)
	GetByUser(idUser int32) ([]entities.RegistroHabito, error)
	GetTotalPuntosByHabito(idHabito int32) (int32, error)
	GetTotalPuntosByUser(idUser int32) (int32, error)
	GetRegistroCompleto(idRegistro int32) (*RegistroHabitoDetalles, error)
	
	// NUEVOS MÉTODOS PARA EL FRONTEND
	GetHabitosConEstadoByFecha(idUser int32, fecha time.Time) ([]entities.HabitoConEstado, error)
	CompletarHabito(idHabito int32, idUser int32, puntos int32) error
	ExistsRegistroHoy(idHabito int32) (bool, error)
}

// RegistroHabitoDetalles - Estructura para obtener registro con datos relacionados
type RegistroHabitoDetalles struct {
	IDRegistroHabito  int32     `json:"id_registro_habito"`
	IDHabito          int32     `json:"id_habito"`
	TituloHabito      string    `json:"titulo_habito"`
	DescripcionHabito string    `json:"descripcion_habito"`
	IDUser            int32     `json:"id_user"`
	Username          string    `json:"username"`
	FechaRealizada    time.Time `json:"fecha_realizada"`
	PuntosGenerados   int32     `json:"puntos_generados"`
}