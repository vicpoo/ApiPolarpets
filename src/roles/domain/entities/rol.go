// rol.go
package entities

type Rol struct {
	ID     int32  `json:"id_rol" gorm:"column:id_rol;primaryKey;autoIncrement"`
	Nombre string `json:"nombre" gorm:"column:nombre;not null"`
}

// Setters
func (r *Rol) SetID(id int32) {
	r.ID = id
}

func (r *Rol) SetNombre(nombre string) {
	r.Nombre = nombre
}

// Getters
func (r *Rol) GetID() int32 {
	return r.ID
}

func (r *Rol) GetNombre() string {
	return r.Nombre
}

// Constructor
func NewRol(nombre string) *Rol {
	return &Rol{
		Nombre: nombre,
	}
}