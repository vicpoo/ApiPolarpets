// mysql_inventario_usuario_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiPolarpets/src/core"
	repositories "github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain"
	"github.com/vicpoo/ApiPolarpets/src/inventario_usuario/domain/entities"
)

type MySQLInventarioUsuarioRepository struct {
	conn *sql.DB
}

func NewMySQLInventarioUsuarioRepository() repositories.IInventarioUsuario {
	conn := core.GetBD()
	return &MySQLInventarioUsuarioRepository{conn: conn}
}

// ========== CRUD BÁSICO ==========

// Save - Agregar un producto al inventario del usuario
func (mysql *MySQLInventarioUsuarioRepository) Save(inventario *entities.InventarioUsuario) error {
	query := `
		INSERT INTO inventario_usuario (id_usuario, id_producto)
		VALUES (?, ?)
	`
	result, err := mysql.conn.Exec(query,
		inventario.GetIDUsuario(),
		inventario.GetIDProducto(),
	)
	if err != nil {
		log.Println("Error al agregar producto al inventario:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	inventario.SetIDInventario(int32(id))

	return nil
}

// Update - Actualizar un registro de inventario
func (mysql *MySQLInventarioUsuarioRepository) Update(inventario *entities.InventarioUsuario) error {
	query := `
		UPDATE inventario_usuario
		SET id_usuario = ?, id_producto = ?
		WHERE id_inventario = ?
	`
	result, err := mysql.conn.Exec(query,
		inventario.GetIDUsuario(),
		inventario.GetIDProducto(),
		inventario.GetIDInventario(),
	)
	if err != nil {
		log.Println("Error al actualizar el inventario:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("registro de inventario con ID %d no encontrado", inventario.GetIDInventario())
	}

	return nil
}

// Delete - Eliminar un registro de inventario por ID
func (mysql *MySQLInventarioUsuarioRepository) Delete(id int32) error {
	query := "DELETE FROM inventario_usuario WHERE id_inventario = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el registro de inventario:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("registro de inventario con ID %d no encontrado", id)
	}

	return nil
}

// DeleteByUserAndProduct - Eliminar un producto específico del inventario de un usuario
func (mysql *MySQLInventarioUsuarioRepository) DeleteByUserAndProduct(idUsuario, idProducto int32) error {
	query := "DELETE FROM inventario_usuario WHERE id_usuario = ? AND id_producto = ?"
	result, err := mysql.conn.Exec(query, idUsuario, idProducto)
	if err != nil {
		log.Println("Error al eliminar producto del inventario del usuario:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("producto %d no encontrado en inventario del usuario %d", idProducto, idUsuario)
	}

	return nil
}

// GetById - Obtener un registro de inventario por ID
func (mysql *MySQLInventarioUsuarioRepository) GetById(id int32) (*entities.InventarioUsuario, error) {
	query := `
		SELECT id_inventario, id_usuario, id_producto
		FROM inventario_usuario
		WHERE id_inventario = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var inventario entities.InventarioUsuario
	var idInventario int32
	var idUsuario int32
	var idProducto int32

	err := row.Scan(&idInventario, &idUsuario, &idProducto)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("registro de inventario con ID %d no encontrado", id)
		}
		log.Println("Error al buscar registro de inventario por ID:", err)
		return nil, err
	}

	inventario.SetIDInventario(idInventario)
	inventario.SetIDUsuario(idUsuario)
	inventario.SetIDProducto(idProducto)

	return &inventario, nil
}

// GetAll - Obtener todos los registros de inventario
func (mysql *MySQLInventarioUsuarioRepository) GetAll() ([]entities.InventarioUsuario, error) {
	query := `
		SELECT id_inventario, id_usuario, id_producto
		FROM inventario_usuario
		ORDER BY id_inventario ASC
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todo el inventario:", err)
		return nil, err
	}
	defer rows.Close()

	var inventario []entities.InventarioUsuario
	for rows.Next() {
		var item entities.InventarioUsuario
		var idInventario int32
		var idUsuario int32
		var idProducto int32

		err := rows.Scan(&idInventario, &idUsuario, &idProducto)
		if err != nil {
			log.Println("Error al escanear registro de inventario:", err)
			return nil, err
		}

		item.SetIDInventario(idInventario)
		item.SetIDUsuario(idUsuario)
		item.SetIDProducto(idProducto)

		inventario = append(inventario, item)
	}

	return inventario, nil
}

// ========== MÉTODOS ADICIONALES ==========

// GetByUser - Obtener todos los productos en el inventario de un usuario
func (mysql *MySQLInventarioUsuarioRepository) GetByUser(idUsuario int32) ([]entities.InventarioUsuario, error) {
	query := `
		SELECT id_inventario, id_usuario, id_producto
		FROM inventario_usuario
		WHERE id_usuario = ?
		ORDER BY id_inventario ASC
	`
	rows, err := mysql.conn.Query(query, idUsuario)
	if err != nil {
		log.Println("Error al obtener inventario del usuario:", err)
		return nil, err
	}
	defer rows.Close()

	var inventario []entities.InventarioUsuario
	for rows.Next() {
		var item entities.InventarioUsuario
		var idInventario int32
		var idUsuarioValue int32
		var idProducto int32

		err := rows.Scan(&idInventario, &idUsuarioValue, &idProducto)
		if err != nil {
			log.Println("Error al escanear registro de inventario:", err)
			return nil, err
		}

		item.SetIDInventario(idInventario)
		item.SetIDUsuario(idUsuarioValue)
		item.SetIDProducto(idProducto)

		inventario = append(inventario, item)
	}

	return inventario, nil
}

// GetByProducto - Obtener qué usuarios tienen un producto específico
func (mysql *MySQLInventarioUsuarioRepository) GetByProducto(idProducto int32) ([]entities.InventarioUsuario, error) {
	query := `
		SELECT id_inventario, id_usuario, id_producto
		FROM inventario_usuario
		WHERE id_producto = ?
		ORDER BY id_inventario ASC
	`
	rows, err := mysql.conn.Query(query, idProducto)
	if err != nil {
		log.Println("Error al obtener usuarios con este producto:", err)
		return nil, err
	}
	defer rows.Close()

	var inventario []entities.InventarioUsuario
	for rows.Next() {
		var item entities.InventarioUsuario
		var idInventario int32
		var idUsuario int32
		var idProductoValue int32

		err := rows.Scan(&idInventario, &idUsuario, &idProductoValue)
		if err != nil {
			log.Println("Error al escanear registro de inventario:", err)
			return nil, err
		}

		item.SetIDInventario(idInventario)
		item.SetIDUsuario(idUsuario)
		item.SetIDProducto(idProductoValue)

		inventario = append(inventario, item)
	}

	return inventario, nil
}

// GetByUserAndProduct - Verificar si un usuario tiene un producto específico
func (mysql *MySQLInventarioUsuarioRepository) GetByUserAndProduct(idUsuario, idProducto int32) (*entities.InventarioUsuario, error) {
	query := `
		SELECT id_inventario, id_usuario, id_producto
		FROM inventario_usuario
		WHERE id_usuario = ? AND id_producto = ?
	`
	row := mysql.conn.QueryRow(query, idUsuario, idProducto)

	var inventario entities.InventarioUsuario
	var idInventario int32
	var idUsuarioValue int32
	var idProductoValue int32

	err := row.Scan(&idInventario, &idUsuarioValue, &idProductoValue)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Error al buscar registro de inventario:", err)
		return nil, err
	}

	inventario.SetIDInventario(idInventario)
	inventario.SetIDUsuario(idUsuarioValue)
	inventario.SetIDProducto(idProductoValue)

	return &inventario, nil
}

// ExistsInInventory - Verificar existencia de un producto en el inventario del usuario
func (mysql *MySQLInventarioUsuarioRepository) ExistsInInventory(idUsuario, idProducto int32) (bool, error) {
	query := `
		SELECT COUNT(*) > 0
		FROM inventario_usuario
		WHERE id_usuario = ? AND id_producto = ?
	`
	var exists bool
	err := mysql.conn.QueryRow(query, idUsuario, idProducto).Scan(&exists)
	if err != nil {
		log.Println("Error al verificar existencia en inventario:", err)
		return false, err
	}
	return exists, nil
}

// GetInventarioByUserWithDetails - Obtener inventario del usuario con detalles del producto
func (mysql *MySQLInventarioUsuarioRepository) GetInventarioByUserWithDetails(idUsuario int32) ([]repositories.InventarioDetalles, error) {
	query := `
		SELECT 
			i.id_inventario,
			i.id_usuario,
			u.username,
			i.id_producto,
			p.nombre as nombre_producto,
			p.descripcion as descripcion_producto,
			p.tipo as tipo_producto,
			p.precio as precio_producto,
			p.id_skin,
			p.id_tipo_mascota
		FROM inventario_usuario i
		INNER JOIN usuarios u ON i.id_usuario = u.id_usuario
		INNER JOIN productos p ON i.id_producto = p.id_producto
		WHERE i.id_usuario = ?
		ORDER BY i.id_inventario ASC
	`
	rows, err := mysql.conn.Query(query, idUsuario)
	if err != nil {
		log.Println("Error al obtener inventario con detalles:", err)
		return nil, err
	}
	defer rows.Close()

	var detalles []repositories.InventarioDetalles
	for rows.Next() {
		var detalle repositories.InventarioDetalles
		var idSkin sql.NullInt32
		var idTipoMascota sql.NullInt32

		err := rows.Scan(
			&detalle.IDInventario,
			&detalle.IDUsuario,
			&detalle.Username,
			&detalle.IDProducto,
			&detalle.NombreProducto,
			&detalle.DescripcionProducto,
			&detalle.TipoProducto,
			&detalle.PrecioProducto,
			&idSkin,
			&idTipoMascota,
		)
		if err != nil {
			log.Println("Error al escanear detalle de inventario:", err)
			return nil, err
		}

		if idSkin.Valid {
			detalle.IDSkin = &idSkin.Int32
		}
		if idTipoMascota.Valid {
			detalle.IDTipoMascota = &idTipoMascota.Int32
		}

		detalles = append(detalles, detalle)
	}

	return detalles, nil
}

// GetInventarioByUserAndProductWithDetails - Obtener producto específico del inventario con detalles
func (mysql *MySQLInventarioUsuarioRepository) GetInventarioByUserAndProductWithDetails(idUsuario, idProducto int32) (*repositories.InventarioDetalles, error) {
	query := `
		SELECT 
			i.id_inventario,
			i.id_usuario,
			u.username,
			i.id_producto,
			p.nombre as nombre_producto,
			p.descripcion as descripcion_producto,
			p.tipo as tipo_producto,
			p.precio as precio_producto,
			p.id_skin,
			p.id_tipo_mascota
		FROM inventario_usuario i
		INNER JOIN usuarios u ON i.id_usuario = u.id_usuario
		INNER JOIN productos p ON i.id_producto = p.id_producto
		WHERE i.id_usuario = ? AND i.id_producto = ?
	`
	row := mysql.conn.QueryRow(query, idUsuario, idProducto)

	var detalle repositories.InventarioDetalles
	var idSkin sql.NullInt32
	var idTipoMascota sql.NullInt32

	err := row.Scan(
		&detalle.IDInventario,
		&detalle.IDUsuario,
		&detalle.Username,
		&detalle.IDProducto,
		&detalle.NombreProducto,
		&detalle.DescripcionProducto,
		&detalle.TipoProducto,
		&detalle.PrecioProducto,
		&idSkin,
		&idTipoMascota,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("producto %d no encontrado en inventario del usuario %d", idProducto, idUsuario)
		}
		log.Println("Error al obtener detalle de inventario:", err)
		return nil, err
	}

	if idSkin.Valid {
		detalle.IDSkin = &idSkin.Int32
	}
	if idTipoMascota.Valid {
		detalle.IDTipoMascota = &idTipoMascota.Int32
	}

	return &detalle, nil
}

// GetCantidadProductosByUser - Obtener la cantidad total de productos que tiene un usuario
func (mysql *MySQLInventarioUsuarioRepository) GetCantidadProductosByUser(idUsuario int32) (int32, error) {
	query := `
		SELECT COUNT(*)
		FROM inventario_usuario
		WHERE id_usuario = ?
	`
	var cantidad int32
	err := mysql.conn.QueryRow(query, idUsuario).Scan(&cantidad)
	if err != nil {
		log.Println("Error al obtener cantidad de productos del usuario:", err)
		return 0, err
	}
	return cantidad, nil
}

// GetProductosByTipoInInventory - Filtrar productos del inventario por tipo
func (mysql *MySQLInventarioUsuarioRepository) GetProductosByTipoInInventory(idUsuario int32, tipo string) ([]entities.InventarioUsuario, error) {
	query := `
		SELECT i.id_inventario, i.id_usuario, i.id_producto
		FROM inventario_usuario i
		INNER JOIN productos p ON i.id_producto = p.id_producto
		WHERE i.id_usuario = ? AND p.tipo = ?
		ORDER BY i.id_inventario ASC
	`
	rows, err := mysql.conn.Query(query, idUsuario, tipo)
	if err != nil {
		log.Println("Error al obtener productos por tipo en inventario:", err)
		return nil, err
	}
	defer rows.Close()

	var inventario []entities.InventarioUsuario
	for rows.Next() {
		var item entities.InventarioUsuario
		var idInventario int32
		var idUsuarioValue int32
		var idProducto int32

		err := rows.Scan(&idInventario, &idUsuarioValue, &idProducto)
		if err != nil {
			log.Println("Error al escanear registro de inventario:", err)
			return nil, err
		}

		item.SetIDInventario(idInventario)
		item.SetIDUsuario(idUsuarioValue)
		item.SetIDProducto(idProducto)

		inventario = append(inventario, item)
	}

	return inventario, nil
}