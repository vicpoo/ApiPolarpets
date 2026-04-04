// skins.go
package entities

type Skins struct {
	IDSkins        int32  `json:"id_skins" gorm:"column:id_skins;primaryKey;autoIncrement"`
	IDTipoMascota  int32  `json:"id_tipo_mascota" gorm:"column:id_tipo_mascota"`
	Nombre         string `json:"nombre" gorm:"column:nombre"`
	ImagenURL      string `json:"imagen_url" gorm:"column:imagen_url"`
}

// Setters
func (s *Skins) SetIDSkins(id int32) {
	s.IDSkins = id
}

func (s *Skins) SetIDTipoMascota(idTipoMascota int32) {
	s.IDTipoMascota = idTipoMascota
}

func (s *Skins) SetNombre(nombre string) {
	s.Nombre = nombre
}

func (s *Skins) SetImagenURL(imagenURL string) {
	s.ImagenURL = imagenURL
}

// Getters
func (s *Skins) GetIDSkins() int32 {
	return s.IDSkins
}

func (s *Skins) GetIDTipoMascota() int32 {
	return s.IDTipoMascota
}

func (s *Skins) GetNombre() string {
	return s.Nombre
}

func (s *Skins) GetImagenURL() string {
	return s.ImagenURL
}

// Constructor
func NewSkins(idTipoMascota int32, nombre, imagenURL string) *Skins {
	return &Skins{
		IDTipoMascota: idTipoMascota,
		Nombre:        nombre,
		ImagenURL:     imagenURL,
	}
}