// mysql_pagos_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/vicpoo/ApiPolarpets/src/core"
	repositories "github.com/vicpoo/ApiPolarpets/src/pagos/domain"
	"github.com/vicpoo/ApiPolarpets/src/pagos/domain/entities"
)

type MySQLPagosRepository struct {
	conn *sql.DB
}

func NewMySQLPagosRepository() repositories.IPagos {
	conn := core.GetBD()
	return &MySQLPagosRepository{conn: conn}
}

// ========== CRUD BÁSICO ==========

// Save - Registrar un nuevo pago
func (mysql *MySQLPagosRepository) Save(pago *entities.Pagos) error {
	query := `
		INSERT INTO pagos (id_usuario, monto, metodo_pago, estado, fecha, referencia_externa)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := mysql.conn.Exec(query,
		pago.GetIDUsuario(),
		pago.GetMonto(),
		pago.GetMetodoPago(),
		pago.GetEstado(),
		pago.GetFecha(),
		pago.GetReferenciaExterna(),
	)
	if err != nil {
		log.Println("Error al registrar el pago:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	pago.SetIDPago(int32(id))

	return nil
}

// Update - Actualizar un pago existente
func (mysql *MySQLPagosRepository) Update(pago *entities.Pagos) error {
	query := `
		UPDATE pagos
		SET id_usuario = ?, monto = ?, metodo_pago = ?, estado = ?, fecha = ?, referencia_externa = ?
		WHERE id_pago = ?
	`
	result, err := mysql.conn.Exec(query,
		pago.GetIDUsuario(),
		pago.GetMonto(),
		pago.GetMetodoPago(),
		pago.GetEstado(),
		pago.GetFecha(),
		pago.GetReferenciaExterna(),
		pago.GetIDPago(),
	)
	if err != nil {
		log.Println("Error al actualizar el pago:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("pago con ID %d no encontrado", pago.GetIDPago())
	}

	return nil
}

// Delete - Eliminar un pago por ID
func (mysql *MySQLPagosRepository) Delete(id int32) error {
	query := "DELETE FROM pagos WHERE id_pago = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el pago:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("pago con ID %d no encontrado", id)
	}

	return nil
}

// GetById - Obtener un pago por ID
func (mysql *MySQLPagosRepository) GetById(id int32) (*entities.Pagos, error) {
	query := `
		SELECT id_pago, id_usuario, monto, metodo_pago, estado, fecha, referencia_externa
		FROM pagos
		WHERE id_pago = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var pago entities.Pagos
	var idPago int32
	var idUsuario int32
	var monto float64
	var metodoPago string
	var estado string
	var fecha time.Time
	var referenciaExterna string

	err := row.Scan(&idPago, &idUsuario, &monto, &metodoPago, &estado, &fecha, &referenciaExterna)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("pago con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el pago por ID:", err)
		return nil, err
	}

	pago.SetIDPago(idPago)
	pago.SetIDUsuario(idUsuario)
	pago.SetMonto(monto)
	pago.SetMetodoPago(metodoPago)
	pago.SetEstado(estado)
	pago.SetFecha(fecha)
	pago.SetReferenciaExterna(referenciaExterna)

	return &pago, nil
}

// GetAll - Obtener todos los pagos
func (mysql *MySQLPagosRepository) GetAll() ([]entities.Pagos, error) {
	query := `
		SELECT id_pago, id_usuario, monto, metodo_pago, estado, fecha, referencia_externa
		FROM pagos
		ORDER BY fecha DESC
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los pagos:", err)
		return nil, err
	}
	defer rows.Close()

	var pagos []entities.Pagos
	for rows.Next() {
		var pago entities.Pagos
		var idPago int32
		var idUsuario int32
		var monto float64
		var metodoPago string
		var estado string
		var fecha time.Time
		var referenciaExterna string

		err := rows.Scan(&idPago, &idUsuario, &monto, &metodoPago, &estado, &fecha, &referenciaExterna)
		if err != nil {
			log.Println("Error al escanear el pago:", err)
			return nil, err
		}

		pago.SetIDPago(idPago)
		pago.SetIDUsuario(idUsuario)
		pago.SetMonto(monto)
		pago.SetMetodoPago(metodoPago)
		pago.SetEstado(estado)
		pago.SetFecha(fecha)
		pago.SetReferenciaExterna(referenciaExterna)

		pagos = append(pagos, pago)
	}

	return pagos, nil
}

// ========== MÉTODOS ADICIONALES ==========

// GetByUser - Obtener todos los pagos de un usuario
func (mysql *MySQLPagosRepository) GetByUser(idUsuario int32) ([]entities.Pagos, error) {
	query := `
		SELECT id_pago, id_usuario, monto, metodo_pago, estado, fecha, referencia_externa
		FROM pagos
		WHERE id_usuario = ?
		ORDER BY fecha DESC
	`
	rows, err := mysql.conn.Query(query, idUsuario)
	if err != nil {
		log.Println("Error al obtener pagos del usuario:", err)
		return nil, err
	}
	defer rows.Close()

	var pagos []entities.Pagos
	for rows.Next() {
		var pago entities.Pagos
		var idPago int32
		var idUsuarioValue int32
		var monto float64
		var metodoPago string
		var estado string
		var fecha time.Time
		var referenciaExterna string

		err := rows.Scan(&idPago, &idUsuarioValue, &monto, &metodoPago, &estado, &fecha, &referenciaExterna)
		if err != nil {
			log.Println("Error al escanear el pago:", err)
			return nil, err
		}

		pago.SetIDPago(idPago)
		pago.SetIDUsuario(idUsuarioValue)
		pago.SetMonto(monto)
		pago.SetMetodoPago(metodoPago)
		pago.SetEstado(estado)
		pago.SetFecha(fecha)
		pago.SetReferenciaExterna(referenciaExterna)

		pagos = append(pagos, pago)
	}

	return pagos, nil
}

// GetByEstado - Obtener pagos por estado
func (mysql *MySQLPagosRepository) GetByEstado(estado string) ([]entities.Pagos, error) {
	query := `
		SELECT id_pago, id_usuario, monto, metodo_pago, estado, fecha, referencia_externa
		FROM pagos
		WHERE estado = ?
		ORDER BY fecha DESC
	`
	rows, err := mysql.conn.Query(query, estado)
	if err != nil {
		log.Println("Error al obtener pagos por estado:", err)
		return nil, err
	}
	defer rows.Close()

	var pagos []entities.Pagos
	for rows.Next() {
		var pago entities.Pagos
		var idPago int32
		var idUsuario int32
		var monto float64
		var metodoPago string
		var estadoValue string
		var fecha time.Time
		var referenciaExterna string

		err := rows.Scan(&idPago, &idUsuario, &monto, &metodoPago, &estadoValue, &fecha, &referenciaExterna)
		if err != nil {
			log.Println("Error al escanear el pago:", err)
			return nil, err
		}

		pago.SetIDPago(idPago)
		pago.SetIDUsuario(idUsuario)
		pago.SetMonto(monto)
		pago.SetMetodoPago(metodoPago)
		pago.SetEstado(estadoValue)
		pago.SetFecha(fecha)
		pago.SetReferenciaExterna(referenciaExterna)

		pagos = append(pagos, pago)
	}

	return pagos, nil
}

// GetByMetodoPago - Obtener pagos por método de pago
func (mysql *MySQLPagosRepository) GetByMetodoPago(metodoPago string) ([]entities.Pagos, error) {
	query := `
		SELECT id_pago, id_usuario, monto, metodo_pago, estado, fecha, referencia_externa
		FROM pagos
		WHERE metodo_pago = ?
		ORDER BY fecha DESC
	`
	rows, err := mysql.conn.Query(query, metodoPago)
	if err != nil {
		log.Println("Error al obtener pagos por método de pago:", err)
		return nil, err
	}
	defer rows.Close()

	var pagos []entities.Pagos
	for rows.Next() {
		var pago entities.Pagos
		var idPago int32
		var idUsuario int32
		var monto float64
		var metodoPagoValue string
		var estado string
		var fecha time.Time
		var referenciaExterna string

		err := rows.Scan(&idPago, &idUsuario, &monto, &metodoPagoValue, &estado, &fecha, &referenciaExterna)
		if err != nil {
			log.Println("Error al escanear el pago:", err)
			return nil, err
		}

		pago.SetIDPago(idPago)
		pago.SetIDUsuario(idUsuario)
		pago.SetMonto(monto)
		pago.SetMetodoPago(metodoPagoValue)
		pago.SetEstado(estado)
		pago.SetFecha(fecha)
		pago.SetReferenciaExterna(referenciaExterna)

		pagos = append(pagos, pago)
	}

	return pagos, nil
}

// GetByFechaRange - Obtener pagos en un rango de fechas
func (mysql *MySQLPagosRepository) GetByFechaRange(fechaInicio, fechaFin time.Time) ([]entities.Pagos, error) {
	query := `
		SELECT id_pago, id_usuario, monto, metodo_pago, estado, fecha, referencia_externa
		FROM pagos
		WHERE DATE(fecha) BETWEEN DATE(?) AND DATE(?)
		ORDER BY fecha DESC
	`
	rows, err := mysql.conn.Query(query, fechaInicio, fechaFin)
	if err != nil {
		log.Println("Error al obtener pagos por rango de fechas:", err)
		return nil, err
	}
	defer rows.Close()

	var pagos []entities.Pagos
	for rows.Next() {
		var pago entities.Pagos
		var idPago int32
		var idUsuario int32
		var monto float64
		var metodoPago string
		var estado string
		var fecha time.Time
		var referenciaExterna string

		err := rows.Scan(&idPago, &idUsuario, &monto, &metodoPago, &estado, &fecha, &referenciaExterna)
		if err != nil {
			log.Println("Error al escanear el pago:", err)
			return nil, err
		}

		pago.SetIDPago(idPago)
		pago.SetIDUsuario(idUsuario)
		pago.SetMonto(monto)
		pago.SetMetodoPago(metodoPago)
		pago.SetEstado(estado)
		pago.SetFecha(fecha)
		pago.SetReferenciaExterna(referenciaExterna)

		pagos = append(pagos, pago)
	}

	return pagos, nil
}

// GetByReferenciaExterna - Buscar pago por referencia externa
func (mysql *MySQLPagosRepository) GetByReferenciaExterna(referencia string) (*entities.Pagos, error) {
	query := `
		SELECT id_pago, id_usuario, monto, metodo_pago, estado, fecha, referencia_externa
		FROM pagos
		WHERE referencia_externa = ?
	`
	row := mysql.conn.QueryRow(query, referencia)

	var pago entities.Pagos
	var idPago int32
	var idUsuario int32
	var monto float64
	var metodoPago string
	var estado string
	var fecha time.Time
	var referenciaExterna string

	err := row.Scan(&idPago, &idUsuario, &monto, &metodoPago, &estado, &fecha, &referenciaExterna)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("pago con referencia '%s' no encontrado", referencia)
		}
		log.Println("Error al buscar pago por referencia externa:", err)
		return nil, err
	}

	pago.SetIDPago(idPago)
	pago.SetIDUsuario(idUsuario)
	pago.SetMonto(monto)
	pago.SetMetodoPago(metodoPago)
	pago.SetEstado(estado)
	pago.SetFecha(fecha)
	pago.SetReferenciaExterna(referenciaExterna)

	return &pago, nil
}

// GetTotalPagadoByUser - Suma total de pagos completados de un usuario
func (mysql *MySQLPagosRepository) GetTotalPagadoByUser(idUsuario int32) (float64, error) {
	query := `
		SELECT COALESCE(SUM(monto), 0)
		FROM pagos
		WHERE id_usuario = ? AND estado = 'completado'
	`
	var total float64
	err := mysql.conn.QueryRow(query, idUsuario).Scan(&total)
	if err != nil {
		log.Println("Error al obtener total pagado por usuario:", err)
		return 0, err
	}
	return total, nil
}

// GetPagosCompletadosByUser - Obtener solo pagos exitosos de un usuario
func (mysql *MySQLPagosRepository) GetPagosCompletadosByUser(idUsuario int32) ([]entities.Pagos, error) {
	query := `
		SELECT id_pago, id_usuario, monto, metodo_pago, estado, fecha, referencia_externa
		FROM pagos
		WHERE id_usuario = ? AND estado = 'completado'
		ORDER BY fecha DESC
	`
	rows, err := mysql.conn.Query(query, idUsuario)
	if err != nil {
		log.Println("Error al obtener pagos completados del usuario:", err)
		return nil, err
	}
	defer rows.Close()

	var pagos []entities.Pagos
	for rows.Next() {
		var pago entities.Pagos
		var idPago int32
		var idUsuarioValue int32
		var monto float64
		var metodoPago string
		var estado string
		var fecha time.Time
		var referenciaExterna string

		err := rows.Scan(&idPago, &idUsuarioValue, &monto, &metodoPago, &estado, &fecha, &referenciaExterna)
		if err != nil {
			log.Println("Error al escanear el pago:", err)
			return nil, err
		}

		pago.SetIDPago(idPago)
		pago.SetIDUsuario(idUsuarioValue)
		pago.SetMonto(monto)
		pago.SetMetodoPago(metodoPago)
		pago.SetEstado(estado)
		pago.SetFecha(fecha)
		pago.SetReferenciaExterna(referenciaExterna)

		pagos = append(pagos, pago)
	}

	return pagos, nil
}

// UpdateEstado - Actualizar solo el estado del pago
func (mysql *MySQLPagosRepository) UpdateEstado(id int32, estado string) error {
	query := `
		UPDATE pagos
		SET estado = ?
		WHERE id_pago = ?
	`
	result, err := mysql.conn.Exec(query, estado, id)
	if err != nil {
		log.Println("Error al actualizar el estado del pago:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("pago con ID %d no encontrado", id)
	}

	return nil
}