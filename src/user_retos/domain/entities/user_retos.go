// user_retos.go
package entities

type UserRetos struct {
	IDUserRetos int32 `json:"id_user_retos" gorm:"column:id_user_retos;primaryKey;autoIncrement"`
	IDUsuario   int32 `json:"id_usuario" gorm:"column:id_usuario"`
	IDReto      int32 `json:"id_reto" gorm:"column:id_reto"`
	Completo    bool  `json:"completo" gorm:"column:completo"`
}

// Setters
func (ur *UserRetos) SetIDUserRetos(id int32) {
	ur.IDUserRetos = id
}

func (ur *UserRetos) SetIDUsuario(idUsuario int32) {
	ur.IDUsuario = idUsuario
}

func (ur *UserRetos) SetIDReto(idReto int32) {
	ur.IDReto = idReto
}

func (ur *UserRetos) SetCompleto(completo bool) {
	ur.Completo = completo
}

// Getters
func (ur *UserRetos) GetIDUserRetos() int32 {
	return ur.IDUserRetos
}

func (ur *UserRetos) GetIDUsuario() int32 {
	return ur.IDUsuario
}

func (ur *UserRetos) GetIDReto() int32 {
	return ur.IDReto
}

func (ur *UserRetos) GetCompleto() bool {
	return ur.Completo
}

// Constructor
func NewUserRetos(idUsuario, idReto int32, completo bool) *UserRetos {
	return &UserRetos{
		IDUsuario: idUsuario,
		IDReto:    idReto,
		Completo:  completo,
	}
}