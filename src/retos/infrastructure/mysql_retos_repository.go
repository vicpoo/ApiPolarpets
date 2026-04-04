// mysql_retos_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiPolarpets/src/core"
	repositories "github.com/vicpoo/ApiPolarpets/src/retos/domain"
	"github.com/vicpoo/ApiPolarpets/src/retos/domain/entities"
)

type MySQLRetosRepository struct {
	conn *sql.DB
}

func NewMySQLRetosRepository() repositories.IRetos {
	conn := core.GetBD()
	return &MySQLRetosRepository{conn: conn}
}

// ========== CRUD BÁSICO ==========

// Save - Guardar un nuevo reto
func (mysql *MySQLRetosRepository) Save(reto *entities.Retos) error {
	query := `
		INSERT INTO retos (titulo, descripcion, puntos_generados)
		VALUES (?, ?, ?)
	`
	result, err := mysql.conn.Exec(query,
		reto.GetTitulo(),
		reto.GetDescripcion(),
		reto.GetPuntosGenerados(),
	)
	if err != nil {
		log.Println("Error al guardar el reto:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	reto.SetIDRetos(int32(id))

	return nil
}

// Update - Actualizar un reto existente
func (mysql *MySQLRetosRepository) Update(reto *entities.Retos) error {
	query := `
		UPDATE retos
		SET titulo = ?, descripcion = ?, puntos_generados = ?
		WHERE id_retos = ?
	`
	result, err := mysql.conn.Exec(query,
		reto.GetTitulo(),
		reto.GetDescripcion(),
		reto.GetPuntosGenerados(),
		reto.GetIDRetos(),
	)
	if err != nil {
		log.Println("Error al actualizar el reto:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("reto con ID %d no encontrado", reto.GetIDRetos())
	}

	return nil
}

// Delete - Eliminar un reto por ID
func (mysql *MySQLRetosRepository) Delete(id int32) error {
	query := "DELETE FROM retos WHERE id_retos = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el reto:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("reto con ID %d no encontrado", id)
	}

	return nil
}

// GetById - Obtener un reto por ID
func (mysql *MySQLRetosRepository) GetById(id int32) (*entities.Retos, error) {
	query := `
		SELECT id_retos, titulo, descripcion, puntos_generados
		FROM retos
		WHERE id_retos = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var reto entities.Retos
	var idRetos int32
	var titulo string
	var descripcion string
	var puntosGenerados int32

	err := row.Scan(&idRetos, &titulo, &descripcion, &puntosGenerados)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("reto con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el reto por ID:", err)
		return nil, err
	}

	reto.SetIDRetos(idRetos)
	reto.SetTitulo(titulo)
	reto.SetDescripcion(descripcion)
	reto.SetPuntosGenerados(puntosGenerados)

	return &reto, nil
}

// GetAll - Obtener todos los retos
func (mysql *MySQLRetosRepository) GetAll() ([]entities.Retos, error) {
	query := `
		SELECT id_retos, titulo, descripcion, puntos_generados
		FROM retos
		ORDER BY id_retos ASC
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los retos:", err)
		return nil, err
	}
	defer rows.Close()

	var retos []entities.Retos
	for rows.Next() {
		var reto entities.Retos
		var idRetos int32
		var titulo string
		var descripcion string
		var puntosGenerados int32

		err := rows.Scan(&idRetos, &titulo, &descripcion, &puntosGenerados)
		if err != nil {
			log.Println("Error al escanear el reto:", err)
			return nil, err
		}

		reto.SetIDRetos(idRetos)
		reto.SetTitulo(titulo)
		reto.SetDescripcion(descripcion)
		reto.SetPuntosGenerados(puntosGenerados)

		retos = append(retos, reto)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return retos, nil
}

// ========== MÉTODOS ADICIONALES ==========

// GetByTitulo - Obtener un reto por título
func (mysql *MySQLRetosRepository) GetByTitulo(titulo string) (*entities.Retos, error) {
	query := `
		SELECT id_retos, titulo, descripcion, puntos_generados
		FROM retos
		WHERE titulo = ?
	`
	row := mysql.conn.QueryRow(query, titulo)

	var reto entities.Retos
	var idRetos int32
	var tituloValue string
	var descripcion string
	var puntosGenerados int32

	err := row.Scan(&idRetos, &tituloValue, &descripcion, &puntosGenerados)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("reto con título '%s' no encontrado", titulo)
		}
		log.Println("Error al buscar el reto por título:", err)
		return nil, err
	}

	reto.SetIDRetos(idRetos)
	reto.SetTitulo(tituloValue)
	reto.SetDescripcion(descripcion)
	reto.SetPuntosGenerados(puntosGenerados)

	return &reto, nil
}

// GetByPuntosRange - Obtener retos por rango de puntos
func (mysql *MySQLRetosRepository) GetByPuntosRange(minPuntos, maxPuntos int32) ([]entities.Retos, error) {
	query := `
		SELECT id_retos, titulo, descripcion, puntos_generados
		FROM retos
		WHERE puntos_generados BETWEEN ? AND ?
		ORDER BY puntos_generados ASC
	`
	rows, err := mysql.conn.Query(query, minPuntos, maxPuntos)
	if err != nil {
		log.Println("Error al obtener retos por rango de puntos:", err)
		return nil, err
	}
	defer rows.Close()

	var retos []entities.Retos
	for rows.Next() {
		var reto entities.Retos
		var idRetos int32
		var titulo string
		var descripcion string
		var puntosGenerados int32

		err := rows.Scan(&idRetos, &titulo, &descripcion, &puntosGenerados)
		if err != nil {
			log.Println("Error al escanear el reto:", err)
			return nil, err
		}

		reto.SetIDRetos(idRetos)
		reto.SetTitulo(titulo)
		reto.SetDescripcion(descripcion)
		reto.SetPuntosGenerados(puntosGenerados)

		retos = append(retos, reto)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return retos, nil
}