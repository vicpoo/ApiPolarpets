// tipo_mascota.go
package entities

type TipoMascota struct {
	IDTipoMascota int32  `json:"id_tipo_mascota" gorm:"column:id_tipo_mascota;primaryKey;autoIncrement"`
	Nombre        string `json:"nombre" gorm:"column:nombre"`
	Descripcion   string `json:"descripcion" gorm:"column:descripcion"`
}

// Setters
func (tm *TipoMascota) SetIDTipoMascota(id int32) {
	tm.IDTipoMascota = id
}

func (tm *TipoMascota) SetNombre(nombre string) {
	tm.Nombre = nombre
}

func (tm *TipoMascota) SetDescripcion(descripcion string) {
	tm.Descripcion = descripcion
}

// Getters
func (tm *TipoMascota) GetIDTipoMascota() int32 {
	return tm.IDTipoMascota
}

func (tm *TipoMascota) GetNombre() string {
	return tm.Nombre
}

func (tm *TipoMascota) GetDescripcion() string {
	return tm.Descripcion
}

// Constructor
func NewTipoMascota(nombre, descripcion string) *TipoMascota {
	return &TipoMascota{
		Nombre:      nombre,
		Descripcion: descripcion,
	}
}