// habito.go
package entities

type Habito struct {
	IDHabito    int32  `json:"id_habito" gorm:"column:id_habito;primaryKey;autoIncrement"`
	IDUser      int32  `json:"id_user" gorm:"column:id_user"`
	Titulo      string `json:"titulo" gorm:"column:titulo"`
	Descripcion string `json:"descripcion" gorm:"column:descripcion"`
	Puntos      int32  `json:"puntos" gorm:"column:puntos"`
}

// Setters
func (h *Habito) SetIDHabito(id int32) {
	h.IDHabito = id
}

func (h *Habito) SetIDUser(idUser int32) {
	h.IDUser = idUser
}

func (h *Habito) SetTitulo(titulo string) {
	h.Titulo = titulo
}

func (h *Habito) SetDescripcion(descripcion string) {
	h.Descripcion = descripcion
}

func (h *Habito) SetPuntos(puntos int32) {
	h.Puntos = puntos
}

// Getters
func (h *Habito) GetIDHabito() int32 {
	return h.IDHabito
}

func (h *Habito) GetIDUser() int32 {
	return h.IDUser
}

func (h *Habito) GetTitulo() string {
	return h.Titulo
}

func (h *Habito) GetDescripcion() string {
	return h.Descripcion
}

func (h *Habito) GetPuntos() int32 {
	return h.Puntos
}

// Constructor
func NewHabito(idUser int32, titulo, descripcion string, puntos int32) *Habito {
	return &Habito{
		IDUser:      idUser,
		Titulo:      titulo,
		Descripcion: descripcion,
		Puntos:      puntos,
	}
}