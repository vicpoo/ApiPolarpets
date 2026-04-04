// mascotas_repository.go
package domain

import (
	"github.com/vicpoo/ApiPolarpets/src/mascotas/domain/entities"
)

type IMascota interface {
	// CRUD básico
	Save(mascota *entities.Mascota) error
	Update(mascota *entities.Mascota) error
	Delete(id int32) error
	GetById(id int32) (*entities.Mascota, error)
	GetAll() ([]entities.Mascota, error)
	
	// Métodos adicionales útiles
	GetByUser(idUser int32) ([]entities.Mascota, error)
	GetByTipoMascota(idTipoMascota int32) ([]entities.Mascota, error)
	GetBySkin(idSkin int32) ([]entities.Mascota, error)
	GetByNivel(idNiveles int32) ([]entities.Mascota, error)
	UpdateExperiencia(idMascota int32, nuevaExperiencia int32) error
	GetMascotaCompleta(idMascota int32) (*MascotaDetalles, error)
}

// MascotaDetalles - Estructura para obtener mascota con datos relacionados
type MascotaDetalles struct {
	IDMascota         int32  `json:"id_mascota"`
	IDUser            int32  `json:"id_user"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	IDNiveles         int32  `json:"id_niveles"`
	Nivel             int32  `json:"nivel"`
	ExpRequerida      int32  `json:"exp_requerida"`
	IDSkin            int32  `json:"id_skin"`
	NombreSkin        string `json:"nombre_skin"`
	ImagenURL         string `json:"imagen_url"`
	IDTipoMascota     int32  `json:"id_tipo_mascota"`
	NombreTipoMascota string `json:"nombre_tipo_mascota"`
	DescripcionTipo   string `json:"descripcion_tipo"`
	ExperienciaActual int32  `json:"experiencia_actual"`
}