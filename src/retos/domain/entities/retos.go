// retos.go
package entities

type Retos struct {
	IDRetos        int32  `json:"id_retos" gorm:"column:id_retos;primaryKey;autoIncrement"`
	Titulo         string `json:"titulo" gorm:"column:titulo"`
	Descripcion    string `json:"descripcion" gorm:"column:descripcion"`
	PuntosGenerados int32  `json:"puntos_generados" gorm:"column:puntos_generados"`
}

// Setters
func (r *Retos) SetIDRetos(id int32) {
	r.IDRetos = id
}

func (r *Retos) SetTitulo(titulo string) {
	r.Titulo = titulo
}

func (r *Retos) SetDescripcion(descripcion string) {
	r.Descripcion = descripcion
}

func (r *Retos) SetPuntosGenerados(puntosGenerados int32) {
	r.PuntosGenerados = puntosGenerados
}

// Getters
func (r *Retos) GetIDRetos() int32 {
	return r.IDRetos
}

func (r *Retos) GetTitulo() string {
	return r.Titulo
}

func (r *Retos) GetDescripcion() string {
	return r.Descripcion
}

func (r *Retos) GetPuntosGenerados() int32 {
	return r.PuntosGenerados
}

// Constructor
func NewRetos(titulo, descripcion string, puntosGenerados int32) *Retos {
	return &Retos{
		Titulo:         titulo,
		Descripcion:    descripcion,
		PuntosGenerados: puntosGenerados,
	}
}