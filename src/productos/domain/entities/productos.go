// productos.go
package entities

type Productos struct {
	IDProducto     int32   `json:"id_producto" gorm:"column:id_producto;primaryKey;autoIncrement"`
	Nombre         string  `json:"nombre" gorm:"column:nombre"`
	Descripcion    string  `json:"descripcion" gorm:"column:descripcion"`
	Tipo           string  `json:"tipo" gorm:"column:tipo"`
	Precio         float64 `json:"precio" gorm:"column:precio"`
	IDSkin         *int32  `json:"id_skin,omitempty" gorm:"column:id_skin"`
	IDTipoMascota  *int32  `json:"id_tipo_mascota,omitempty" gorm:"column:id_tipo_mascota"`
}

// Setters
func (p *Productos) SetIDProducto(id int32) {
	p.IDProducto = id
}

func (p *Productos) SetNombre(nombre string) {
	p.Nombre = nombre
}

func (p *Productos) SetDescripcion(descripcion string) {
	p.Descripcion = descripcion
}

func (p *Productos) SetTipo(tipo string) {
	p.Tipo = tipo
}

func (p *Productos) SetPrecio(precio float64) {
	p.Precio = precio
}

func (p *Productos) SetIDSkin(idSkin *int32) {
	p.IDSkin = idSkin
}

func (p *Productos) SetIDTipoMascota(idTipoMascota *int32) {
	p.IDTipoMascota = idTipoMascota
}

// Getters
func (p *Productos) GetIDProducto() int32 {
	return p.IDProducto
}

func (p *Productos) GetNombre() string {
	return p.Nombre
}

func (p *Productos) GetDescripcion() string {
	return p.Descripcion
}

func (p *Productos) GetTipo() string {
	return p.Tipo
}

func (p *Productos) GetPrecio() float64 {
	return p.Precio
}

func (p *Productos) GetIDSkin() *int32 {
	return p.IDSkin
}

func (p *Productos) GetIDTipoMascota() *int32 {
	return p.IDTipoMascota
}

// Constructor
func NewProductos(nombre, descripcion, tipo string, precio float64, idSkin *int32, idTipoMascota *int32) *Productos {
	return &Productos{
		Nombre:        nombre,
		Descripcion:   descripcion,
		Tipo:          tipo,
		Precio:        precio,
		IDSkin:        idSkin,
		IDTipoMascota: idTipoMascota,
	}
}