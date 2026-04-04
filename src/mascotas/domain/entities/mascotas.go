// mascotas.go
package entities

type Mascota struct {
	IDMascota        int32 `json:"id_mascota" gorm:"column:id_mascota;primaryKey;autoIncrement"`
	IDUser           int32 `json:"id_user" gorm:"column:id_user"`
	IDNiveles        int32 `json:"id_niveles" gorm:"column:id_niveles"`
	IDSkin           int32 `json:"id_skin" gorm:"column:id_skin"`
	IDTipoMascota    int32 `json:"id_tipo_mascota" gorm:"column:id_tipo_mascota"`
	ExperienciaActual int32 `json:"experiencia_actual" gorm:"column:experiencia_actual"`
}

// Setters
func (m *Mascota) SetIDMascota(id int32) {
	m.IDMascota = id
}

func (m *Mascota) SetIDUser(idUser int32) {
	m.IDUser = idUser
}

func (m *Mascota) SetIDNiveles(idNiveles int32) {
	m.IDNiveles = idNiveles
}

func (m *Mascota) SetIDSkin(idSkin int32) {
	m.IDSkin = idSkin
}

func (m *Mascota) SetIDTipoMascota(idTipoMascota int32) {
	m.IDTipoMascota = idTipoMascota
}

func (m *Mascota) SetExperienciaActual(experienciaActual int32) {
	m.ExperienciaActual = experienciaActual
}

// Getters
func (m *Mascota) GetIDMascota() int32 {
	return m.IDMascota
}

func (m *Mascota) GetIDUser() int32 {
	return m.IDUser
}

func (m *Mascota) GetIDNiveles() int32 {
	return m.IDNiveles
}

func (m *Mascota) GetIDSkin() int32 {
	return m.IDSkin
}

func (m *Mascota) GetIDTipoMascota() int32 {
	return m.IDTipoMascota
}

func (m *Mascota) GetExperienciaActual() int32 {
	return m.ExperienciaActual
}

// Constructor
func NewMascota(idUser, idNiveles, idSkin, idTipoMascota, experienciaActual int32) *Mascota {
	return &Mascota{
		IDUser:           idUser,
		IDNiveles:        idNiveles,
		IDSkin:           idSkin,
		IDTipoMascota:    idTipoMascota,
		ExperienciaActual: experienciaActual,
	}
}