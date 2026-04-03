// usuario_repository.go
package domain

import (
	"github.com/vicpoo/ApiPolarpets/src/usuarios/domain/entities"
)

type IUsuario interface {
	// Registro y Login
	Register(usuario *entities.Usuario) error
	Login(email, password string) (*entities.Usuario, error)
	
	// CRUD completo
	Save(usuario *entities.Usuario) error
	Update(usuario *entities.Usuario) error
	Delete(id int32) error
	GetById(id int32) (*entities.Usuario, error)
	GetAll() ([]entities.Usuario, error)
	
	// Métodos adicionales útiles
	GetByEmail(email string) (*entities.Usuario, error)
	UpdateMascotaActiva(idUsuario int32, idMascotaActiva *int32) error
}