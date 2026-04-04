// compras_repository.go
package domain

import (
	"time"

	"github.com/vicpoo/ApiPolarpets/src/compras/domain/entities"
)

type ICompras interface {
	// CRUD básico
	Save(compra *entities.Compras) error
	Update(compra *entities.Compras) error
	Delete(id int32) error
	GetById(id int32) (*entities.Compras, error)
	GetAll() ([]entities.Compras, error)
	
	// Métodos adicionales útiles
	GetByUser(idUsuario int32) ([]entities.Compras, error)
	GetByProducto(idProducto int32) ([]entities.Compras, error)
	GetByPago(idPago int32) (*entities.Compras, error)
	GetByFechaRange(fechaInicio, fechaFin time.Time) ([]entities.Compras, error)
	GetComprasByUserWithDetails(idUsuario int32) ([]CompraDetalles, error)
	GetCompraByIdWithDetails(idCompra int32) (*CompraDetalles, error)
	GetTotalGastadoByUser(idUsuario int32) (float64, error)
	GetComprasRecientesByUser(idUsuario int32, limit int) ([]entities.Compras, error)
}

// CompraDetalles - Estructura para obtener compra con detalles de usuario, producto y pago
type CompraDetalles struct {
	IDCompra         int32     `json:"id_compra"`
	IDUsuario        int32     `json:"id_usuario"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	IDProducto       int32     `json:"id_producto"`
	NombreProducto   string    `json:"nombre_producto"`
	DescripcionProducto string  `json:"descripcion_producto"`
	TipoProducto     string    `json:"tipo_producto"`
	PrecioProducto   float64   `json:"precio_producto"`
	IDPago           int32     `json:"id_pago"`
	MontoPago        float64   `json:"monto_pago"`
	MetodoPago       string    `json:"metodo_pago"`
	EstadoPago       string    `json:"estado_pago"`
	ReferenciaExterna string   `json:"referencia_externa,omitempty"`
	FechaCompra      time.Time `json:"fecha_compra"`
}