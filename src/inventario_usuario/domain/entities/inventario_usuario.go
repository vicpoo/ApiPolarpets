// inventario_usuario.go
package entities

type InventarioUsuario struct {
	IDInventario int32 `json:"id_inventario" gorm:"column:id_inventario;primaryKey;autoIncrement"`
	IDUsuario    int32 `json:"id_usuario" gorm:"column:id_usuario"`
	IDProducto   int32 `json:"id_producto" gorm:"column:id_producto"`
}

// Setters
func (iu *InventarioUsuario) SetIDInventario(id int32) {
	iu.IDInventario = id
}

func (iu *InventarioUsuario) SetIDUsuario(idUsuario int32) {
	iu.IDUsuario = idUsuario
}

func (iu *InventarioUsuario) SetIDProducto(idProducto int32) {
	iu.IDProducto = idProducto
}

// Getters
func (iu *InventarioUsuario) GetIDInventario() int32 {
	return iu.IDInventario
}

func (iu *InventarioUsuario) GetIDUsuario() int32 {
	return iu.IDUsuario
}

func (iu *InventarioUsuario) GetIDProducto() int32 {
	return iu.IDProducto
}

// Constructor
func NewInventarioUsuario(idUsuario, idProducto int32) *InventarioUsuario {
	return &InventarioUsuario{
		IDUsuario:  idUsuario,
		IDProducto: idProducto,
	}
}