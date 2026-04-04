// productos_repository.go
package domain

import (
	"github.com/vicpoo/ApiPolarpets/src/productos/domain/entities"
)

type IProductos interface {
	// CRUD básico
	Save(producto *entities.Productos) error
	Update(producto *entities.Productos) error
	Delete(id int32) error
	GetById(id int32) (*entities.Productos, error)
	GetAll() ([]entities.Productos, error)
	
	// Métodos adicionales útiles
	GetByTipo(tipo string) ([]entities.Productos, error)
	GetByPrecioRange(minPrecio, maxPrecio float64) ([]entities.Productos, error)
	GetBySkin(idSkin int32) ([]entities.Productos, error)
	GetByTipoMascota(idTipoMascota int32) ([]entities.Productos, error)
	GetByNombre(nombre string) (*entities.Productos, error)
	GetProductosConDetalles(id int32) (*ProductoDetalles, error)
	GetAllProductosConDetalles() ([]ProductoDetalles, error)
}

// ProductoDetalles - Estructura para obtener producto con detalles de skin y tipo mascota
type ProductoDetalles struct {
	IDProducto        int32   `json:"id_producto"`
	Nombre            string  `json:"nombre"`
	Descripcion       string  `json:"descripcion"`
	Tipo              string  `json:"tipo"`
	Precio            float64 `json:"precio"`
	IDSkin            *int32  `json:"id_skin,omitempty"`
	NombreSkin        *string `json:"nombre_skin,omitempty"`
	ImagenURL         *string `json:"imagen_url,omitempty"`
	IDTipoMascota     *int32  `json:"id_tipo_mascota,omitempty"`
	NombreTipoMascota *string `json:"nombre_tipo_mascota,omitempty"`
}