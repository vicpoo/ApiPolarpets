// mysql_user_retos_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiPolarpets/src/core"
	repositories "github.com/vicpoo/ApiPolarpets/src/user_retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/user_retos/domain/entities"
)

type MySQLUserRetosRepository struct {
	conn *sql.DB
}

func NewMySQLUserRetosRepository() repositories.IUserRetos {
	conn := core.GetBD()
	return &MySQLUserRetosRepository{conn: conn}
}

// ========== CRUD BÁSICO ==========

// Save - Asignar un reto a un usuario
func (mysql *MySQLUserRetosRepository) Save(userReto *entities.UserRetos) error {
	query := `
		INSERT INTO user_retos (id_usuario, id_reto, completo)
		VALUES (?, ?, ?)
	`
	result, err := mysql.conn.Exec(query,
		userReto.GetIDUsuario(),
		userReto.GetIDReto(),
		userReto.GetCompleto(),
	)
	if err != nil {
		log.Println("Error al asignar el reto al usuario:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	userReto.SetIDUserRetos(int32(id))

	return nil
}

// Update - Actualizar un registro usuario-reto
func (mysql *MySQLUserRetosRepository) Update(userReto *entities.UserRetos) error {
	query := `
		UPDATE user_retos
		SET id_usuario = ?, id_reto = ?, completo = ?
		WHERE id_user_retos = ?
	`
	result, err := mysql.conn.Exec(query,
		userReto.GetIDUsuario(),
		userReto.GetIDReto(),
		userReto.GetCompleto(),
		userReto.GetIDUserRetos(),
	)
	if err != nil {
		log.Println("Error al actualizar el registro:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("registro con ID %d no encontrado", userReto.GetIDUserRetos())
	}

	return nil
}

// Delete - Eliminar un registro por ID
func (mysql *MySQLUserRetosRepository) Delete(id int32) error {
	query := "DELETE FROM user_retos WHERE id_user_retos = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el registro:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("registro con ID %d no encontrado", id)
	}

	return nil
}

// GetById - Obtener un registro por ID
func (mysql *MySQLUserRetosRepository) GetById(id int32) (*entities.UserRetos, error) {
	query := `
		SELECT id_user_retos, id_usuario, id_reto, completo
		FROM user_retos
		WHERE id_user_retos = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var userReto entities.UserRetos
	var idUserRetos int32
	var idUsuario int32
	var idReto int32
	var completo bool

	err := row.Scan(&idUserRetos, &idUsuario, &idReto, &completo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("registro con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el registro por ID:", err)
		return nil, err
	}

	userReto.SetIDUserRetos(idUserRetos)
	userReto.SetIDUsuario(idUsuario)
	userReto.SetIDReto(idReto)
	userReto.SetCompleto(completo)

	return &userReto, nil
}

// GetAll - Obtener todos los registros
func (mysql *MySQLUserRetosRepository) GetAll() ([]entities.UserRetos, error) {
	query := `
		SELECT id_user_retos, id_usuario, id_reto, completo
		FROM user_retos
		ORDER BY id_user_retos ASC
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los registros:", err)
		return nil, err
	}
	defer rows.Close()

	var userRetos []entities.UserRetos
	for rows.Next() {
		var userReto entities.UserRetos
		var idUserRetos int32
		var idUsuario int32
		var idReto int32
		var completo bool

		err := rows.Scan(&idUserRetos, &idUsuario, &idReto, &completo)
		if err != nil {
			log.Println("Error al escanear el registro:", err)
			return nil, err
		}

		userReto.SetIDUserRetos(idUserRetos)
		userReto.SetIDUsuario(idUsuario)
		userReto.SetIDReto(idReto)
		userReto.SetCompleto(completo)

		userRetos = append(userRetos, userReto)
	}

	return userRetos, nil
}

// ========== MÉTODOS ADICIONALES ==========

// GetByUser - Obtener todos los retos de un usuario
func (mysql *MySQLUserRetosRepository) GetByUser(idUsuario int32) ([]entities.UserRetos, error) {
	query := `
		SELECT id_user_retos, id_usuario, id_reto, completo
		FROM user_retos
		WHERE id_usuario = ?
		ORDER BY id_user_retos ASC
	`
	rows, err := mysql.conn.Query(query, idUsuario)
	if err != nil {
		log.Println("Error al obtener retos del usuario:", err)
		return nil, err
	}
	defer rows.Close()

	var userRetos []entities.UserRetos
	for rows.Next() {
		var userReto entities.UserRetos
		var idUserRetos int32
		var idUsuarioValue int32
		var idReto int32
		var completo bool

		err := rows.Scan(&idUserRetos, &idUsuarioValue, &idReto, &completo)
		if err != nil {
			log.Println("Error al escanear el registro:", err)
			return nil, err
		}

		userReto.SetIDUserRetos(idUserRetos)
		userReto.SetIDUsuario(idUsuarioValue)
		userReto.SetIDReto(idReto)
		userReto.SetCompleto(completo)

		userRetos = append(userRetos, userReto)
	}

	return userRetos, nil
}

// GetByReto - Obtener todos los usuarios que tienen un reto específico
func (mysql *MySQLUserRetosRepository) GetByReto(idReto int32) ([]entities.UserRetos, error) {
	query := `
		SELECT id_user_retos, id_usuario, id_reto, completo
		FROM user_retos
		WHERE id_reto = ?
		ORDER BY id_user_retos ASC
	`
	rows, err := mysql.conn.Query(query, idReto)
	if err != nil {
		log.Println("Error al obtener usuarios con este reto:", err)
		return nil, err
	}
	defer rows.Close()

	var userRetos []entities.UserRetos
	for rows.Next() {
		var userReto entities.UserRetos
		var idUserRetos int32
		var idUsuario int32
		var idRetoValue int32
		var completo bool

		err := rows.Scan(&idUserRetos, &idUsuario, &idRetoValue, &completo)
		if err != nil {
			log.Println("Error al escanear el registro:", err)
			return nil, err
		}

		userReto.SetIDUserRetos(idUserRetos)
		userReto.SetIDUsuario(idUsuario)
		userReto.SetIDReto(idRetoValue)
		userReto.SetCompleto(completo)

		userRetos = append(userRetos, userReto)
	}

	return userRetos, nil
}

// GetByUserAndReto - Verificar si un usuario tiene asignado un reto específico
func (mysql *MySQLUserRetosRepository) GetByUserAndReto(idUsuario, idReto int32) (*entities.UserRetos, error) {
	query := `
		SELECT id_user_retos, id_usuario, id_reto, completo
		FROM user_retos
		WHERE id_usuario = ? AND id_reto = ?
	`
	row := mysql.conn.QueryRow(query, idUsuario, idReto)

	var userReto entities.UserRetos
	var idUserRetos int32
	var idUsuarioValue int32
	var idRetoValue int32
	var completo bool

	err := row.Scan(&idUserRetos, &idUsuarioValue, &idRetoValue, &completo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No encontrado, no es error
		}
		log.Println("Error al buscar el registro:", err)
		return nil, err
	}

	userReto.SetIDUserRetos(idUserRetos)
	userReto.SetIDUsuario(idUsuarioValue)
	userReto.SetIDReto(idRetoValue)
	userReto.SetCompleto(completo)

	return &userReto, nil
}

// GetCompletedByUser - Obtener solo los retos completados por un usuario
func (mysql *MySQLUserRetosRepository) GetCompletedByUser(idUsuario int32) ([]entities.UserRetos, error) {
	query := `
		SELECT id_user_retos, id_usuario, id_reto, completo
		FROM user_retos
		WHERE id_usuario = ? AND completo = 1
		ORDER BY id_user_retos ASC
	`
	rows, err := mysql.conn.Query(query, idUsuario)
	if err != nil {
		log.Println("Error al obtener retos completados:", err)
		return nil, err
	}
	defer rows.Close()

	var userRetos []entities.UserRetos
	for rows.Next() {
		var userReto entities.UserRetos
		var idUserRetos int32
		var idUsuarioValue int32
		var idReto int32
		var completo bool

		err := rows.Scan(&idUserRetos, &idUsuarioValue, &idReto, &completo)
		if err != nil {
			log.Println("Error al escanear el registro:", err)
			return nil, err
		}

		userReto.SetIDUserRetos(idUserRetos)
		userReto.SetIDUsuario(idUsuarioValue)
		userReto.SetIDReto(idReto)
		userReto.SetCompleto(completo)

		userRetos = append(userRetos, userReto)
	}

	return userRetos, nil
}

// GetPendingByUser - Obtener solo los retos pendientes de un usuario
func (mysql *MySQLUserRetosRepository) GetPendingByUser(idUsuario int32) ([]entities.UserRetos, error) {
	query := `
		SELECT id_user_retos, id_usuario, id_reto, completo
		FROM user_retos
		WHERE id_usuario = ? AND completo = 0
		ORDER BY id_user_retos ASC
	`
	rows, err := mysql.conn.Query(query, idUsuario)
	if err != nil {
		log.Println("Error al obtener retos pendientes:", err)
		return nil, err
	}
	defer rows.Close()

	var userRetos []entities.UserRetos
	for rows.Next() {
		var userReto entities.UserRetos
		var idUserRetos int32
		var idUsuarioValue int32
		var idReto int32
		var completo bool

		err := rows.Scan(&idUserRetos, &idUsuarioValue, &idReto, &completo)
		if err != nil {
			log.Println("Error al escanear el registro:", err)
			return nil, err
		}

		userReto.SetIDUserRetos(idUserRetos)
		userReto.SetIDUsuario(idUsuarioValue)
		userReto.SetIDReto(idReto)
		userReto.SetCompleto(completo)

		userRetos = append(userRetos, userReto)
	}

	return userRetos, nil
}

// CompleteReto - Marcar un reto como completado
func (mysql *MySQLUserRetosRepository) CompleteReto(idUsuario, idReto int32) error {
	query := `
		UPDATE user_retos
		SET completo = 1
		WHERE id_usuario = ? AND id_reto = ?
	`
	result, err := mysql.conn.Exec(query, idUsuario, idReto)
	if err != nil {
		log.Println("Error al completar el reto:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el reto para el usuario %d con reto %d", idUsuario, idReto)
	}

	return nil
}

// GetUserRetosConDetalles - Obtener retos del usuario con detalles del reto
func (mysql *MySQLUserRetosRepository) GetUserRetosConDetalles(idUsuario int32) ([]repositories.UserRetoDetalles, error) {
	query := `
		SELECT 
			ur.id_user_retos,
			ur.id_usuario,
			u.username,
			ur.id_reto,
			r.titulo,
			r.descripcion,
			r.puntos_generados,
			ur.completo
		FROM user_retos ur
		INNER JOIN usuarios u ON ur.id_usuario = u.id_usuario
		INNER JOIN retos r ON ur.id_reto = r.id_retos
		WHERE ur.id_usuario = ?
		ORDER BY ur.completo ASC, r.puntos_generados DESC
	`
	rows, err := mysql.conn.Query(query, idUsuario)
	if err != nil {
		log.Println("Error al obtener detalles de retos del usuario:", err)
		return nil, err
	}
	defer rows.Close()

	var detalles []repositories.UserRetoDetalles
	for rows.Next() {
		var detalle repositories.UserRetoDetalles
		err := rows.Scan(
			&detalle.IDUserRetos,
			&detalle.IDUsuario,
			&detalle.Username,
			&detalle.IDReto,
			&detalle.TituloReto,
			&detalle.DescripcionReto,
			&detalle.PuntosGenerados,
			&detalle.Completo,
		)
		if err != nil {
			log.Println("Error al escanear el detalle:", err)
			return nil, err
		}
		detalles = append(detalles, detalle)
	}

	return detalles, nil
}