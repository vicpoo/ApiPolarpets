// mysql_compras_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/vicpoo/ApiPolarpets/src/core"
	repositories "github.com/vicpoo/ApiPolarpets/src/compras/domain"
	"github.com/vicpoo/ApiPolarpets/src/compras/domain/entities"
)

type MySQLComprasRepository struct {
	conn *sql.DB
}

func NewMySQLComprasRepository() repositories.ICompras {
	conn := core.GetBD()
	return &MySQLComprasRepository{conn: conn}
}

// ========== CRUD BÁSICO ==========

// Save - Registrar una nueva compra
func (mysql *MySQLComprasRepository) Save(compra *entities.Compras) error {
	query := `
		INSERT INTO compras (id_usuario, id_producto, id_pago, fecha)
		VALUES (?, ?, ?, ?)
	`
	result, err := mysql.conn.Exec(query,
		compra.GetIDUsuario(),
		compra.GetIDProducto(),
		compra.GetIDPago(),
		compra.GetFecha(),
	)
	if err != nil {
		log.Println("Error al registrar la compra:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	compra.SetIDCompra(int32(id))

	return nil
}

// Update - Actualizar una compra existente
func (mysql *MySQLComprasRepository) Update(compra *entities.Compras) error {
	query := `
		UPDATE compras
		SET id_usuario = ?, id_producto = ?, id_pago = ?, fecha = ?
		WHERE id_compra = ?
	`
	result, err := mysql.conn.Exec(query,
		compra.GetIDUsuario(),
		compra.GetIDProducto(),
		compra.GetIDPago(),
		compra.GetFecha(),
		compra.GetIDCompra(),
	)
	if err != nil {
		log.Println("Error al actualizar la compra:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("compra con ID %d no encontrada", compra.GetIDCompra())
	}

	return nil
}

// Delete - Eliminar una compra por ID
func (mysql *MySQLComprasRepository) Delete(id int32) error {
	query := "DELETE FROM compras WHERE id_compra = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar la compra:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("compra con ID %d no encontrada", id)
	}

	return nil
}

// GetById - Obtener una compra por ID
func (mysql *MySQLComprasRepository) GetById(id int32) (*entities.Compras, error) {
	query := `
		SELECT id_compra, id_usuario, id_producto, id_pago, fecha
		FROM compras
		WHERE id_compra = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var compra entities.Compras
	var idCompra int32
	var idUsuario int32
	var idProducto int32
	var idPago int32
	var fecha time.Time

	err := row.Scan(&idCompra, &idUsuario, &idProducto, &idPago, &fecha)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("compra con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la compra por ID:", err)
		return nil, err
	}

	compra.SetIDCompra(idCompra)
	compra.SetIDUsuario(idUsuario)
	compra.SetIDProducto(idProducto)
	compra.SetIDPago(idPago)
	compra.SetFecha(fecha)

	return &compra, nil
}

// GetAll - Obtener todas las compras
func (mysql *MySQLComprasRepository) GetAll() ([]entities.Compras, error) {
	query := `
		SELECT id_compra, id_usuario, id_producto, id_pago, fecha
		FROM compras
		ORDER BY fecha DESC
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todas las compras:", err)
		return nil, err
	}
	defer rows.Close()

	var compras []entities.Compras
	for rows.Next() {
		var compra entities.Compras
		var idCompra int32
		var idUsuario int32
		var idProducto int32
		var idPago int32
		var fecha time.Time

		err := rows.Scan(&idCompra, &idUsuario, &idProducto, &idPago, &fecha)
		if err != nil {
			log.Println("Error al escanear la compra:", err)
			return nil, err
		}

		compra.SetIDCompra(idCompra)
		compra.SetIDUsuario(idUsuario)
		compra.SetIDProducto(idProducto)
		compra.SetIDPago(idPago)
		compra.SetFecha(fecha)

		compras = append(compras, compra)
	}

	return compras, nil
}

// ========== MÉTODOS ADICIONALES ==========

// GetByUser - Obtener todas las compras de un usuario
func (mysql *MySQLComprasRepository) GetByUser(idUsuario int32) ([]entities.Compras, error) {
	query := `
		SELECT id_compra, id_usuario, id_producto, id_pago, fecha
		FROM compras
		WHERE id_usuario = ?
		ORDER BY fecha DESC
	`
	rows, err := mysql.conn.Query(query, idUsuario)
	if err != nil {
		log.Println("Error al obtener compras del usuario:", err)
		return nil, err
	}
	defer rows.Close()

	var compras []entities.Compras
	for rows.Next() {
		var compra entities.Compras
		var idCompra int32
		var idUsuarioValue int32
		var idProducto int32
		var idPago int32
		var fecha time.Time

		err := rows.Scan(&idCompra, &idUsuarioValue, &idProducto, &idPago, &fecha)
		if err != nil {
			log.Println("Error al escanear la compra:", err)
			return nil, err
		}

		compra.SetIDCompra(idCompra)
		compra.SetIDUsuario(idUsuarioValue)
		compra.SetIDProducto(idProducto)
		compra.SetIDPago(idPago)
		compra.SetFecha(fecha)

		compras = append(compras, compra)
	}

	return compras, nil
}

// GetByProducto - Obtener compras de un producto específico
func (mysql *MySQLComprasRepository) GetByProducto(idProducto int32) ([]entities.Compras, error) {
	query := `
		SELECT id_compra, id_usuario, id_producto, id_pago, fecha
		FROM compras
		WHERE id_producto = ?
		ORDER BY fecha DESC
	`
	rows, err := mysql.conn.Query(query, idProducto)
	if err != nil {
		log.Println("Error al obtener compras del producto:", err)
		return nil, err
	}
	defer rows.Close()

	var compras []entities.Compras
	for rows.Next() {
		var compra entities.Compras
		var idCompra int32
		var idUsuario int32
		var idProductoValue int32
		var idPago int32
		var fecha time.Time

		err := rows.Scan(&idCompra, &idUsuario, &idProductoValue, &idPago, &fecha)
		if err != nil {
			log.Println("Error al escanear la compra:", err)
			return nil, err
		}

		compra.SetIDCompra(idCompra)
		compra.SetIDUsuario(idUsuario)
		compra.SetIDProducto(idProductoValue)
		compra.SetIDPago(idPago)
		compra.SetFecha(fecha)

		compras = append(compras, compra)
	}

	return compras, nil
}

// GetByPago - Obtener compra asociada a un pago
func (mysql *MySQLComprasRepository) GetByPago(idPago int32) (*entities.Compras, error) {
	query := `
		SELECT id_compra, id_usuario, id_producto, id_pago, fecha
		FROM compras
		WHERE id_pago = ?
	`
	row := mysql.conn.QueryRow(query, idPago)

	var compra entities.Compras
	var idCompra int32
	var idUsuario int32
	var idProducto int32
	var idPagoValue int32
	var fecha time.Time

	err := row.Scan(&idCompra, &idUsuario, &idProducto, &idPagoValue, &fecha)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("compra con pago ID %d no encontrada", idPago)
		}
		log.Println("Error al buscar compra por pago:", err)
		return nil, err
	}

	compra.SetIDCompra(idCompra)
	compra.SetIDUsuario(idUsuario)
	compra.SetIDProducto(idProducto)
	compra.SetIDPago(idPagoValue)
	compra.SetFecha(fecha)

	return &compra, nil
}

// GetByFechaRange - Obtener compras en un rango de fechas
func (mysql *MySQLComprasRepository) GetByFechaRange(fechaInicio, fechaFin time.Time) ([]entities.Compras, error) {
	query := `
		SELECT id_compra, id_usuario, id_producto, id_pago, fecha
		FROM compras
		WHERE DATE(fecha) BETWEEN DATE(?) AND DATE(?)
		ORDER BY fecha DESC
	`
	rows, err := mysql.conn.Query(query, fechaInicio, fechaFin)
	if err != nil {
		log.Println("Error al obtener compras por rango de fechas:", err)
		return nil, err
	}
	defer rows.Close()

	var compras []entities.Compras
	for rows.Next() {
		var compra entities.Compras
		var idCompra int32
		var idUsuario int32
		var idProducto int32
		var idPago int32
		var fecha time.Time

		err := rows.Scan(&idCompra, &idUsuario, &idProducto, &idPago, &fecha)
		if err != nil {
			log.Println("Error al escanear la compra:", err)
			return nil, err
		}

		compra.SetIDCompra(idCompra)
		compra.SetIDUsuario(idUsuario)
		compra.SetIDProducto(idProducto)
		compra.SetIDPago(idPago)
		compra.SetFecha(fecha)

		compras = append(compras, compra)
	}

	return compras, nil
}

// GetComprasByUserWithDetails - Obtener compras de usuario con detalles
func (mysql *MySQLComprasRepository) GetComprasByUserWithDetails(idUsuario int32) ([]repositories.CompraDetalles, error) {
	query := `
		SELECT 
			c.id_compra,
			c.id_usuario,
			u.username,
			u.email,
			c.id_producto,
			p.nombre as nombre_producto,
			p.descripcion as descripcion_producto,
			p.tipo as tipo_producto,
			p.precio as precio_producto,
			c.id_pago,
			pg.monto as monto_pago,
			pg.metodo_pago,
			pg.estado as estado_pago,
			pg.referencia_externa,
			c.fecha as fecha_compra
		FROM compras c
		INNER JOIN usuarios u ON c.id_usuario = u.id_usuario
		INNER JOIN productos p ON c.id_producto = p.id_producto
		INNER JOIN pagos pg ON c.id_pago = pg.id_pago
		WHERE c.id_usuario = ?
		ORDER BY c.fecha DESC
	`
	rows, err := mysql.conn.Query(query, idUsuario)
	if err != nil {
		log.Println("Error al obtener compras con detalles:", err)
		return nil, err
	}
	defer rows.Close()

	var detalles []repositories.CompraDetalles
	for rows.Next() {
		var detalle repositories.CompraDetalles
		var referenciaExterna sql.NullString

		err := rows.Scan(
			&detalle.IDCompra,
			&detalle.IDUsuario,
			&detalle.Username,
			&detalle.Email,
			&detalle.IDProducto,
			&detalle.NombreProducto,
			&detalle.DescripcionProducto,
			&detalle.TipoProducto,
			&detalle.PrecioProducto,
			&detalle.IDPago,
			&detalle.MontoPago,
			&detalle.MetodoPago,
			&detalle.EstadoPago,
			&referenciaExterna,
			&detalle.FechaCompra,
		)
		if err != nil {
			log.Println("Error al escanear detalle de compra:", err)
			return nil, err
		}

		if referenciaExterna.Valid {
			detalle.ReferenciaExterna = referenciaExterna.String
		}

		detalles = append(detalles, detalle)
	}

	return detalles, nil
}

// GetCompraByIdWithDetails - Obtener una compra con todos sus detalles
func (mysql *MySQLComprasRepository) GetCompraByIdWithDetails(idCompra int32) (*repositories.CompraDetalles, error) {
	query := `
		SELECT 
			c.id_compra,
			c.id_usuario,
			u.username,
			u.email,
			c.id_producto,
			p.nombre as nombre_producto,
			p.descripcion as descripcion_producto,
			p.tipo as tipo_producto,
			p.precio as precio_producto,
			c.id_pago,
			pg.monto as monto_pago,
			pg.metodo_pago,
			pg.estado as estado_pago,
			pg.referencia_externa,
			c.fecha as fecha_compra
		FROM compras c
		INNER JOIN usuarios u ON c.id_usuario = u.id_usuario
		INNER JOIN productos p ON c.id_producto = p.id_producto
		INNER JOIN pagos pg ON c.id_pago = pg.id_pago
		WHERE c.id_compra = ?
	`
	row := mysql.conn.QueryRow(query, idCompra)

	var detalle repositories.CompraDetalles
	var referenciaExterna sql.NullString

	err := row.Scan(
		&detalle.IDCompra,
		&detalle.IDUsuario,
		&detalle.Username,
		&detalle.Email,
		&detalle.IDProducto,
		&detalle.NombreProducto,
		&detalle.DescripcionProducto,
		&detalle.TipoProducto,
		&detalle.PrecioProducto,
		&detalle.IDPago,
		&detalle.MontoPago,
		&detalle.MetodoPago,
		&detalle.EstadoPago,
		&referenciaExterna,
		&detalle.FechaCompra,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("compra con ID %d no encontrada", idCompra)
		}
		log.Println("Error al obtener detalle de compra:", err)
		return nil, err
	}

	if referenciaExterna.Valid {
		detalle.ReferenciaExterna = referenciaExterna.String
	}

	return &detalle, nil
}

// GetTotalGastadoByUser - Suma total gastada por un usuario
func (mysql *MySQLComprasRepository) GetTotalGastadoByUser(idUsuario int32) (float64, error) {
	query := `
		SELECT COALESCE(SUM(p.precio), 0)
		FROM compras c
		INNER JOIN productos p ON c.id_producto = p.id_producto
		WHERE c.id_usuario = ?
	`
	var total float64
	err := mysql.conn.QueryRow(query, idUsuario).Scan(&total)
	if err != nil {
		log.Println("Error al obtener total gastado por usuario:", err)
		return 0, err
	}
	return total, nil
}

// GetComprasRecientesByUser - Obtener las últimas N compras de un usuario
func (mysql *MySQLComprasRepository) GetComprasRecientesByUser(idUsuario int32, limit int) ([]entities.Compras, error) {
	query := `
		SELECT id_compra, id_usuario, id_producto, id_pago, fecha
		FROM compras
		WHERE id_usuario = ?
		ORDER BY fecha DESC
		LIMIT ?
	`
	rows, err := mysql.conn.Query(query, idUsuario, limit)
	if err != nil {
		log.Println("Error al obtener compras recientes del usuario:", err)
		return nil, err
	}
	defer rows.Close()

	var compras []entities.Compras
	for rows.Next() {
		var compra entities.Compras
		var idCompra int32
		var idUsuarioValue int32
		var idProducto int32
		var idPago int32
		var fecha time.Time

		err := rows.Scan(&idCompra, &idUsuarioValue, &idProducto, &idPago, &fecha)
		if err != nil {
			log.Println("Error al escanear la compra:", err)
			return nil, err
		}

		compra.SetIDCompra(idCompra)
		compra.SetIDUsuario(idUsuarioValue)
		compra.SetIDProducto(idProducto)
		compra.SetIDPago(idPago)
		compra.SetFecha(fecha)

		compras = append(compras, compra)
	}

	return compras, nil
}