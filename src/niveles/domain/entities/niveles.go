// niveles.go
package entities

type Niveles struct {
	IDNiveles     int32 `json:"id_niveles" gorm:"column:id_niveles;primaryKey;autoIncrement"`
	Nivel         int32 `json:"nivel" gorm:"column:nivel"`
	ExpRequerida  int32 `json:"exp_requerida" gorm:"column:exp_requerida"`
}

// Setters
func (n *Niveles) SetIDNiveles(id int32) {
	n.IDNiveles = id
}

func (n *Niveles) SetNivel(nivel int32) {
	n.Nivel = nivel
}

func (n *Niveles) SetExpRequerida(expRequerida int32) {
	n.ExpRequerida = expRequerida
}

// Getters
func (n *Niveles) GetIDNiveles() int32 {
	return n.IDNiveles
}

func (n *Niveles) GetNivel() int32 {
	return n.Nivel
}

func (n *Niveles) GetExpRequerida() int32 {
	return n.ExpRequerida
}

// Constructor
func NewNiveles(nivel, expRequerida int32) *Niveles {
	return &Niveles{
		Nivel:        nivel,
		ExpRequerida: expRequerida,
	}
}