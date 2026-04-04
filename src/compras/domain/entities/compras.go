// compras.go
package entities

import "time"

type Compras struct {
	IDCompra   int32     `json:"id_compra" gorm:"column:id_compra;primaryKey;autoIncrement"`
	IDUsuario  int32     `json:"id_usuario" gorm:"column:id_usuario"`
	IDProducto int32     `json:"id_producto" gorm:"column:id_producto"`
	IDPago     int32     `json:"id_pago" gorm:"column:id_pago"`
	Fecha      time.Time `json:"fecha" gorm:"column:fecha"`
}

// Setters
func (c *Compras) SetIDCompra(id int32) {
	c.IDCompra = id
}

func (c *Compras) SetIDUsuario(idUsuario int32) {
	c.IDUsuario = idUsuario
}

func (c *Compras) SetIDProducto(idProducto int32) {
	c.IDProducto = idProducto
}

func (c *Compras) SetIDPago(idPago int32) {
	c.IDPago = idPago
}

func (c *Compras) SetFecha(fecha time.Time) {
	c.Fecha = fecha
}

// Getters
func (c *Compras) GetIDCompra() int32 {
	return c.IDCompra
}

func (c *Compras) GetIDUsuario() int32 {
	return c.IDUsuario
}

func (c *Compras) GetIDProducto() int32 {
	return c.IDProducto
}

func (c *Compras) GetIDPago() int32 {
	return c.IDPago
}

func (c *Compras) GetFecha() time.Time {
	return c.Fecha
}

// Constructor
func NewCompras(idUsuario, idProducto, idPago int32) *Compras {
	return &Compras{
		IDUsuario:  idUsuario,
		IDProducto: idProducto,
		IDPago:     idPago,
		Fecha:      time.Now(),
	}
}