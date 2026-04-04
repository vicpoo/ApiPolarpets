// pagos.go
package entities

import "time"

type Pagos struct {
	IDPago           int32     `json:"id_pago" gorm:"column:id_pago;primaryKey;autoIncrement"`
	IDUsuario        int32     `json:"id_usuario" gorm:"column:id_usuario"`
	Monto            float64   `json:"monto" gorm:"column:monto"`
	MetodoPago       string    `json:"metodo_pago" gorm:"column:metodo_pago"`
	Estado           string    `json:"estado" gorm:"column:estado"`
	Fecha            time.Time `json:"fecha" gorm:"column:fecha"`
	ReferenciaExterna string   `json:"referencia_externa" gorm:"column:referencia_externa"`
}

// Setters
func (p *Pagos) SetIDPago(id int32) {
	p.IDPago = id
}

func (p *Pagos) SetIDUsuario(idUsuario int32) {
	p.IDUsuario = idUsuario
}

func (p *Pagos) SetMonto(monto float64) {
	p.Monto = monto
}

func (p *Pagos) SetMetodoPago(metodoPago string) {
	p.MetodoPago = metodoPago
}

func (p *Pagos) SetEstado(estado string) {
	p.Estado = estado
}

func (p *Pagos) SetFecha(fecha time.Time) {
	p.Fecha = fecha
}

func (p *Pagos) SetReferenciaExterna(referenciaExterna string) {
	p.ReferenciaExterna = referenciaExterna
}

// Getters
func (p *Pagos) GetIDPago() int32 {
	return p.IDPago
}

func (p *Pagos) GetIDUsuario() int32 {
	return p.IDUsuario
}

func (p *Pagos) GetMonto() float64 {
	return p.Monto
}

func (p *Pagos) GetMetodoPago() string {
	return p.MetodoPago
}

func (p *Pagos) GetEstado() string {
	return p.Estado
}

func (p *Pagos) GetFecha() time.Time {
	return p.Fecha
}

func (p *Pagos) GetReferenciaExterna() string {
	return p.ReferenciaExterna
}

// Constructor
func NewPagos(idUsuario int32, monto float64, metodoPago, estado, referenciaExterna string) *Pagos {
	return &Pagos{
		IDUsuario:        idUsuario,
		Monto:            monto,
		MetodoPago:       metodoPago,
		Estado:           estado,
		Fecha:            time.Now(),
		ReferenciaExterna: referenciaExterna,
	}
}