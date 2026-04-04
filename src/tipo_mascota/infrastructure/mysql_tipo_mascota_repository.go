// mysql_tipo_mascota_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiPolarpets/src/core"
	repositories "github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain"
	"github.com/vicpoo/ApiPolarpets/src/tipo_mascota/domain/entities"
)

type MySQLTipoMascotaRepository struct {
	conn *sql.DB
}

func NewMySQLTipoMascotaRepository() repositories.ITipoMascota {
	conn := core.GetBD()
	return &MySQLTipoMascotaRepository{conn: conn}
}

// Save - Guardar un nuevo tipo de mascota
func (mysql *MySQLTipoMascotaRepository) Save(tipoMascota *entities.TipoMascota) error {
	query := `
		INSERT INTO tipo_mascota (nombre, descripcion)
		VALUES (?, ?)
	`
	result, err := mysql.conn.Exec(query, tipoMascota.GetNombre(), tipoMascota.GetDescripcion())
	if err != nil {
		log.Println("Error al guardar el tipo de mascota:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	tipoMascota.SetIDTipoMascota(int32(id))

	return nil
}

// Update - Actualizar un tipo de mascota existente
func (mysql *MySQLTipoMascotaRepository) Update(tipoMascota *entities.TipoMascota) error {
	query := `
		UPDATE tipo_mascota
		SET nombre = ?, descripcion = ?
		WHERE id_tipo_mascota = ?
	`
	result, err := mysql.conn.Exec(query, tipoMascota.GetNombre(), tipoMascota.GetDescripcion(), tipoMascota.GetIDTipoMascota())
	if err != nil {
		log.Println("Error al actualizar el tipo de mascota:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("tipo de mascota con ID %d no encontrado", tipoMascota.GetIDTipoMascota())
	}

	return nil
}

// Delete - Eliminar un tipo de mascota por ID
func (mysql *MySQLTipoMascotaRepository) Delete(id int32) error {
	query := "DELETE FROM tipo_mascota WHERE id_tipo_mascota = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el tipo de mascota:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("tipo de mascota con ID %d no encontrado", id)
	}

	return nil
}

// GetById - Obtener un tipo de mascota por ID
func (mysql *MySQLTipoMascotaRepository) GetById(id int32) (*entities.TipoMascota, error) {
	query := `
		SELECT id_tipo_mascota, nombre, descripcion
		FROM tipo_mascota
		WHERE id_tipo_mascota = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var tipoMascota entities.TipoMascota
	var idTipoMascota int32
	var nombre string
	var descripcion string

	err := row.Scan(&idTipoMascota, &nombre, &descripcion)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tipo de mascota con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el tipo de mascota por ID:", err)
		return nil, err
	}

	tipoMascota.SetIDTipoMascota(idTipoMascota)
	tipoMascota.SetNombre(nombre)
	tipoMascota.SetDescripcion(descripcion)

	return &tipoMascota, nil
}

// GetAll - Obtener todos los tipos de mascota
func (mysql *MySQLTipoMascotaRepository) GetAll() ([]entities.TipoMascota, error) {
	query := `
		SELECT id_tipo_mascota, nombre, descripcion
		FROM tipo_mascota
		ORDER BY nombre ASC
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los tipos de mascota:", err)
		return nil, err
	}
	defer rows.Close()

	var tiposMascota []entities.TipoMascota
	for rows.Next() {
		var tipoMascota entities.TipoMascota
		var idTipoMascota int32
		var nombre string
		var descripcion string

		err := rows.Scan(&idTipoMascota, &nombre, &descripcion)
		if err != nil {
			log.Println("Error al escanear el tipo de mascota:", err)
			return nil, err
		}

		tipoMascota.SetIDTipoMascota(idTipoMascota)
		tipoMascota.SetNombre(nombre)
		tipoMascota.SetDescripcion(descripcion)

		tiposMascota = append(tiposMascota, tipoMascota)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return tiposMascota, nil
}

// GetByNombre - Obtener tipo de mascota por nombre
func (mysql *MySQLTipoMascotaRepository) GetByNombre(nombre string) (*entities.TipoMascota, error) {
	query := `
		SELECT id_tipo_mascota, nombre, descripcion
		FROM tipo_mascota
		WHERE nombre = ?
	`
	row := mysql.conn.QueryRow(query, nombre)

	var tipoMascota entities.TipoMascota
	var idTipoMascota int32
	var nombreValue string
	var descripcion string

	err := row.Scan(&idTipoMascota, &nombreValue, &descripcion)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tipo de mascota con nombre '%s' no encontrado", nombre)
		}
		log.Println("Error al buscar el tipo de mascota por nombre:", err)
		return nil, err
	}

	tipoMascota.SetIDTipoMascota(idTipoMascota)
	tipoMascota.SetNombre(nombreValue)
	tipoMascota.SetDescripcion(descripcion)

	return &tipoMascota, nil
}