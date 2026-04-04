// inventario_usuario_repository.go
package domain

import (
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain/entities"
)

type IInventarioUsuario interface {
	// CRUD básico
	Save(inventario *entities.InventarioUsuario) error
	Update(inventario *entities.InventarioUsuario) error
	Delete(id int32) error
	DeleteByUserAndProduct(idUsuario, idProducto int32) error
	GetById(id int32) (*entities.InventarioUsuario, error)
	GetAll() ([]entities.InventarioUsuario, error)
	
	// Métodos adicionales útiles
	GetByUser(idUsuario int32) ([]entities.InventarioUsuario, error)
	GetByProducto(idProducto int32) ([]entities.InventarioUsuario, error)
	GetByUserAndProduct(idUsuario, idProducto int32) (*entities.InventarioUsuario, error)
	ExistsInInventory(idUsuario, idProducto int32) (bool, error)
	GetInventarioByUserWithDetails(idUsuario int32) ([]InventarioDetalles, error)
	GetInventarioByUserAndProductWithDetails(idUsuario, idProducto int32) (*InventarioDetalles, error)
	GetCantidadProductosByUser(idUsuario int32) (int32, error)
	GetProductosByTipoInInventory(idUsuario int32, tipo string) ([]entities.InventarioUsuario, error)
}

// InventarioDetalles - Estructura para obtener inventario con detalles del producto
type InventarioDetalles struct {
	IDInventario     int32   `json:"id_inventario"`
	IDUsuario        int32   `json:"id_usuario"`
	Username         string  `json:"username"`
	IDProducto       int32   `json:"id_producto"`
	NombreProducto   string  `json:"nombre_producto"`
	DescripcionProducto string `json:"descripcion_producto"`
	TipoProducto     string  `json:"tipo_producto"`
	PrecioProducto   float64 `json:"precio_producto"`
	IDSkin           *int32  `json:"id_skin,omitempty"`
	IDTipoMascota    *int32  `json:"id_tipo_mascota,omitempty"`
}