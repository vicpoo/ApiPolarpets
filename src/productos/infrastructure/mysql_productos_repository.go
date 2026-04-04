// mysql_productos_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiPolarpets/src/core"
	repositories "github.com/vicpoo/ApiPolarpets/src/productos/domain"
	"github.com/vicpoo/ApiPolarpets/src/productos/domain/entities"
)

type MySQLProductosRepository struct {
	conn *sql.DB
}

func NewMySQLProductosRepository() repositories.IProductos {
	conn := core.GetBD()
	return &MySQLProductosRepository{conn: conn}
}

// ========== CRUD BÁSICO ==========

// Save - Guardar un nuevo producto
func (mysql *MySQLProductosRepository) Save(producto *entities.Productos) error {
	query := `
		INSERT INTO productos (nombre, descripcion, tipo, precio, id_skin, id_tipo_mascota)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := mysql.conn.Exec(query,
		producto.GetNombre(),
		producto.GetDescripcion(),
		producto.GetTipo(),
		producto.GetPrecio(),
		producto.GetIDSkin(),
		producto.GetIDTipoMascota(),
	)
	if err != nil {
		log.Println("Error al guardar el producto:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	producto.SetIDProducto(int32(id))

	return nil
}

// Update - Actualizar un producto existente
func (mysql *MySQLProductosRepository) Update(producto *entities.Productos) error {
	query := `
		UPDATE productos
		SET nombre = ?, descripcion = ?, tipo = ?, precio = ?, id_skin = ?, id_tipo_mascota = ?
		WHERE id_producto = ?
	`
	result, err := mysql.conn.Exec(query,
		producto.GetNombre(),
		producto.GetDescripcion(),
		producto.GetTipo(),
		producto.GetPrecio(),
		producto.GetIDSkin(),
		producto.GetIDTipoMascota(),
		producto.GetIDProducto(),
	)
	if err != nil {
		log.Println("Error al actualizar el producto:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("producto con ID %d no encontrado", producto.GetIDProducto())
	}

	return nil
}

// Delete - Eliminar un producto por ID
func (mysql *MySQLProductosRepository) Delete(id int32) error {
	query := "DELETE FROM productos WHERE id_producto = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el producto:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("producto con ID %d no encontrado", id)
	}

	return nil
}

// GetById - Obtener un producto por ID
func (mysql *MySQLProductosRepository) GetById(id int32) (*entities.Productos, error) {
	query := `
		SELECT id_producto, nombre, descripcion, tipo, precio, id_skin, id_tipo_mascota
		FROM productos
		WHERE id_producto = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var producto entities.Productos
	var idProducto int32
	var nombre string
	var descripcion string
	var tipo string
	var precio float64
	var idSkin sql.NullInt32
	var idTipoMascota sql.NullInt32

	err := row.Scan(&idProducto, &nombre, &descripcion, &tipo, &precio, &idSkin, &idTipoMascota)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("producto con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el producto por ID:", err)
		return nil, err
	}

	producto.SetIDProducto(idProducto)
	producto.SetNombre(nombre)
	producto.SetDescripcion(descripcion)
	producto.SetTipo(tipo)
	producto.SetPrecio(precio)

	if idSkin.Valid {
		producto.SetIDSkin(&idSkin.Int32)
	}
	if idTipoMascota.Valid {
		producto.SetIDTipoMascota(&idTipoMascota.Int32)
	}

	return &producto, nil
}

// GetAll - Obtener todos los productos
func (mysql *MySQLProductosRepository) GetAll() ([]entities.Productos, error) {
	query := `
		SELECT id_producto, nombre, descripcion, tipo, precio, id_skin, id_tipo_mascota
		FROM productos
		ORDER BY id_producto ASC
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los productos:", err)
		return nil, err
	}
	defer rows.Close()

	var productos []entities.Productos
	for rows.Next() {
		var producto entities.Productos
		var idProducto int32
		var nombre string
		var descripcion string
		var tipo string
		var precio float64
		var idSkin sql.NullInt32
		var idTipoMascota sql.NullInt32

		err := rows.Scan(&idProducto, &nombre, &descripcion, &tipo, &precio, &idSkin, &idTipoMascota)
		if err != nil {
			log.Println("Error al escanear el producto:", err)
			return nil, err
		}

		producto.SetIDProducto(idProducto)
		producto.SetNombre(nombre)
		producto.SetDescripcion(descripcion)
		producto.SetTipo(tipo)
		producto.SetPrecio(precio)

		if idSkin.Valid {
			producto.SetIDSkin(&idSkin.Int32)
		}
		if idTipoMascota.Valid {
			producto.SetIDTipoMascota(&idTipoMascota.Int32)
		}

		productos = append(productos, producto)
	}

	return productos, nil
}

// ========== MÉTODOS ADICIONALES ==========

// GetByTipo - Obtener productos por tipo
func (mysql *MySQLProductosRepository) GetByTipo(tipo string) ([]entities.Productos, error) {
	query := `
		SELECT id_producto, nombre, descripcion, tipo, precio, id_skin, id_tipo_mascota
		FROM productos
		WHERE tipo = ?
		ORDER BY nombre ASC
	`
	rows, err := mysql.conn.Query(query, tipo)
	if err != nil {
		log.Println("Error al obtener productos por tipo:", err)
		return nil, err
	}
	defer rows.Close()

	var productos []entities.Productos
	for rows.Next() {
		var producto entities.Productos
		var idProducto int32
		var nombre string
		var descripcion string
		var tipoValue string
		var precio float64
		var idSkin sql.NullInt32
		var idTipoMascota sql.NullInt32

		err := rows.Scan(&idProducto, &nombre, &descripcion, &tipoValue, &precio, &idSkin, &idTipoMascota)
		if err != nil {
			log.Println("Error al escanear el producto:", err)
			return nil, err
		}

		producto.SetIDProducto(idProducto)
		producto.SetNombre(nombre)
		producto.SetDescripcion(descripcion)
		producto.SetTipo(tipoValue)
		producto.SetPrecio(precio)

		if idSkin.Valid {
			producto.SetIDSkin(&idSkin.Int32)
		}
		if idTipoMascota.Valid {
			producto.SetIDTipoMascota(&idTipoMascota.Int32)
		}

		productos = append(productos, producto)
	}

	return productos, nil
}

// GetByPrecioRange - Obtener productos por rango de precio
func (mysql *MySQLProductosRepository) GetByPrecioRange(minPrecio, maxPrecio float64) ([]entities.Productos, error) {
	query := `
		SELECT id_producto, nombre, descripcion, tipo, precio, id_skin, id_tipo_mascota
		FROM productos
		WHERE precio BETWEEN ? AND ?
		ORDER BY precio ASC
	`
	rows, err := mysql.conn.Query(query, minPrecio, maxPrecio)
	if err != nil {
		log.Println("Error al obtener productos por rango de precio:", err)
		return nil, err
	}
	defer rows.Close()

	var productos []entities.Productos
	for rows.Next() {
		var producto entities.Productos
		var idProducto int32
		var nombre string
		var descripcion string
		var tipo string
		var precio float64
		var idSkin sql.NullInt32
		var idTipoMascota sql.NullInt32

		err := rows.Scan(&idProducto, &nombre, &descripcion, &tipo, &precio, &idSkin, &idTipoMascota)
		if err != nil {
			log.Println("Error al escanear el producto:", err)
			return nil, err
		}

		producto.SetIDProducto(idProducto)
		producto.SetNombre(nombre)
		producto.SetDescripcion(descripcion)
		producto.SetTipo(tipo)
		producto.SetPrecio(precio)

		if idSkin.Valid {
			producto.SetIDSkin(&idSkin.Int32)
		}
		if idTipoMascota.Valid {
			producto.SetIDTipoMascota(&idTipoMascota.Int32)
		}

		productos = append(productos, producto)
	}

	return productos, nil
}

// GetBySkin - Obtener productos asociados a una skin
func (mysql *MySQLProductosRepository) GetBySkin(idSkin int32) ([]entities.Productos, error) {
	query := `
		SELECT id_producto, nombre, descripcion, tipo, precio, id_skin, id_tipo_mascota
		FROM productos
		WHERE id_skin = ?
		ORDER BY nombre ASC
	`
	rows, err := mysql.conn.Query(query, idSkin)
	if err != nil {
		log.Println("Error al obtener productos por skin:", err)
		return nil, err
	}
	defer rows.Close()

	var productos []entities.Productos
	for rows.Next() {
		var producto entities.Productos
		var idProducto int32
		var nombre string
		var descripcion string
		var tipo string
		var precio float64
		var idSkinValue sql.NullInt32
		var idTipoMascota sql.NullInt32

		err := rows.Scan(&idProducto, &nombre, &descripcion, &tipo, &precio, &idSkinValue, &idTipoMascota)
		if err != nil {
			log.Println("Error al escanear el producto:", err)
			return nil, err
		}

		producto.SetIDProducto(idProducto)
		producto.SetNombre(nombre)
		producto.SetDescripcion(descripcion)
		producto.SetTipo(tipo)
		producto.SetPrecio(precio)

		if idSkinValue.Valid {
			producto.SetIDSkin(&idSkinValue.Int32)
		}
		if idTipoMascota.Valid {
			producto.SetIDTipoMascota(&idTipoMascota.Int32)
		}

		productos = append(productos, producto)
	}

	return productos, nil
}

// GetByTipoMascota - Obtener productos para un tipo de mascota específico
func (mysql *MySQLProductosRepository) GetByTipoMascota(idTipoMascota int32) ([]entities.Productos, error) {
	query := `
		SELECT id_producto, nombre, descripcion, tipo, precio, id_skin, id_tipo_mascota
		FROM productos
		WHERE id_tipo_mascota = ?
		ORDER BY nombre ASC
	`
	rows, err := mysql.conn.Query(query, idTipoMascota)
	if err != nil {
		log.Println("Error al obtener productos por tipo de mascota:", err)
		return nil, err
	}
	defer rows.Close()

	var productos []entities.Productos
	for rows.Next() {
		var producto entities.Productos
		var idProducto int32
		var nombre string
		var descripcion string
		var tipo string
		var precio float64
		var idSkin sql.NullInt32
		var idTipoMascotaValue sql.NullInt32

		err := rows.Scan(&idProducto, &nombre, &descripcion, &tipo, &precio, &idSkin, &idTipoMascotaValue)
		if err != nil {
			log.Println("Error al escanear el producto:", err)
			return nil, err
		}

		producto.SetIDProducto(idProducto)
		producto.SetNombre(nombre)
		producto.SetDescripcion(descripcion)
		producto.SetTipo(tipo)
		producto.SetPrecio(precio)

		if idSkin.Valid {
			producto.SetIDSkin(&idSkin.Int32)
		}
		if idTipoMascotaValue.Valid {
			producto.SetIDTipoMascota(&idTipoMascotaValue.Int32)
		}

		productos = append(productos, producto)
	}

	return productos, nil
}

// GetByNombre - Obtener un producto por nombre
func (mysql *MySQLProductosRepository) GetByNombre(nombre string) (*entities.Productos, error) {
	query := `
		SELECT id_producto, nombre, descripcion, tipo, precio, id_skin, id_tipo_mascota
		FROM productos
		WHERE nombre = ?
	`
	row := mysql.conn.QueryRow(query, nombre)

	var producto entities.Productos
	var idProducto int32
	var nombreValue string
	var descripcion string
	var tipo string
	var precio float64
	var idSkin sql.NullInt32
	var idTipoMascota sql.NullInt32

	err := row.Scan(&idProducto, &nombreValue, &descripcion, &tipo, &precio, &idSkin, &idTipoMascota)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("producto con nombre '%s' no encontrado", nombre)
		}
		log.Println("Error al buscar el producto por nombre:", err)
		return nil, err
	}

	producto.SetIDProducto(idProducto)
	producto.SetNombre(nombreValue)
	producto.SetDescripcion(descripcion)
	producto.SetTipo(tipo)
	producto.SetPrecio(precio)

	if idSkin.Valid {
		producto.SetIDSkin(&idSkin.Int32)
	}
	if idTipoMascota.Valid {
		producto.SetIDTipoMascota(&idTipoMascota.Int32)
	}

	return &producto, nil
}

// GetProductosConDetalles - Obtener producto con detalles de skin y tipo mascota
func (mysql *MySQLProductosRepository) GetProductosConDetalles(id int32) (*repositories.ProductoDetalles, error) {
	query := `
		SELECT 
			p.id_producto,
			p.nombre,
			p.descripcion,
			p.tipo,
			p.precio,
			p.id_skin,
			s.nombre as nombre_skin,
			s.imagen_url,
			p.id_tipo_mascota,
			tm.nombre as nombre_tipo_mascota
		FROM productos p
		LEFT JOIN skins s ON p.id_skin = s.id_skins
		LEFT JOIN tipo_mascota tm ON p.id_tipo_mascota = tm.id_tipo_mascota
		WHERE p.id_producto = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var detalles repositories.ProductoDetalles
	var idSkin sql.NullInt32
	var nombreSkin sql.NullString
	var imagenURL sql.NullString
	var idTipoMascota sql.NullInt32
	var nombreTipoMascota sql.NullString

	err := row.Scan(
		&detalles.IDProducto,
		&detalles.Nombre,
		&detalles.Descripcion,
		&detalles.Tipo,
		&detalles.Precio,
		&idSkin,
		&nombreSkin,
		&imagenURL,
		&idTipoMascota,
		&nombreTipoMascota,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("producto con ID %d no encontrado", id)
		}
		log.Println("Error al obtener producto con detalles:", err)
		return nil, err
	}

	if idSkin.Valid {
		detalles.IDSkin = &idSkin.Int32
		detalles.NombreSkin = &nombreSkin.String
		detalles.ImagenURL = &imagenURL.String
	}
	if idTipoMascota.Valid {
		detalles.IDTipoMascota = &idTipoMascota.Int32
		detalles.NombreTipoMascota = &nombreTipoMascota.String
	}

	return &detalles, nil
}

// GetAllProductosConDetalles - Obtener todos los productos con detalles
func (mysql *MySQLProductosRepository) GetAllProductosConDetalles() ([]repositories.ProductoDetalles, error) {
	query := `
		SELECT 
			p.id_producto,
			p.nombre,
			p.descripcion,
			p.tipo,
			p.precio,
			p.id_skin,
			s.nombre as nombre_skin,
			s.imagen_url,
			p.id_tipo_mascota,
			tm.nombre as nombre_tipo_mascota
		FROM productos p
		LEFT JOIN skins s ON p.id_skin = s.id_skins
		LEFT JOIN tipo_mascota tm ON p.id_tipo_mascota = tm.id_tipo_mascota
		ORDER BY p.id_producto ASC
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los productos con detalles:", err)
		return nil, err
	}
	defer rows.Close()

	var detalles []repositories.ProductoDetalles
	for rows.Next() {
		var detalle repositories.ProductoDetalles
		var idSkin sql.NullInt32
		var nombreSkin sql.NullString
		var imagenURL sql.NullString
		var idTipoMascota sql.NullInt32
		var nombreTipoMascota sql.NullString

		err := rows.Scan(
			&detalle.IDProducto,
			&detalle.Nombre,
			&detalle.Descripcion,
			&detalle.Tipo,
			&detalle.Precio,
			&idSkin,
			&nombreSkin,
			&imagenURL,
			&idTipoMascota,
			&nombreTipoMascota,
		)
		if err != nil {
			log.Println("Error al escanear producto con detalles:", err)
			return nil, err
		}

		if idSkin.Valid {
			detalle.IDSkin = &idSkin.Int32
			detalle.NombreSkin = &nombreSkin.String
			detalle.ImagenURL = &imagenURL.String
		}
		if idTipoMascota.Valid {
			detalle.IDTipoMascota = &idTipoMascota.Int32
			detalle.NombreTipoMascota = &nombreTipoMascota.String
		}

		detalles = append(detalles, detalle)
	}

	return detalles, nil
}