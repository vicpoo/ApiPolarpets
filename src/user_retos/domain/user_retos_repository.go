// user_retos_repository.go
package domain

import (
	"github.com/vicpoo/ApiPolarpets/src/user_retos/domain/entities"
)

type IUserRetos interface {
	// CRUD básico
	Save(userReto *entities.UserRetos) error
	Update(userReto *entities.UserRetos) error
	Delete(id int32) error
	GetById(id int32) (*entities.UserRetos, error)
	GetAll() ([]entities.UserRetos, error)
	
	// Métodos adicionales útiles
	GetByUser(idUsuario int32) ([]entities.UserRetos, error)
	GetByReto(idReto int32) ([]entities.UserRetos, error)
	GetByUserAndReto(idUsuario, idReto int32) (*entities.UserRetos, error)
	GetCompletedByUser(idUsuario int32) ([]entities.UserRetos, error)
	GetPendingByUser(idUsuario int32) ([]entities.UserRetos, error)
	CompleteReto(idUsuario, idReto int32) error
	GetUserRetosConDetalles(idUsuario int32) ([]UserRetoDetalles, error)
}

// UserRetoDetalles - Estructura para obtener retos del usuario con detalles del reto
type UserRetoDetalles struct {
	IDUserRetos    int32  `json:"id_user_retos"`
	IDUsuario      int32  `json:"id_usuario"`
	Username       string `json:"username"`
	IDReto         int32  `json:"id_reto"`
	TituloReto     string `json:"titulo_reto"`
	DescripcionReto string `json:"descripcion_reto"`
	PuntosGenerados int32  `json:"puntos_generados"`
	Completo       bool   `json:"completo"`
}