// mysql_skins_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiPolarpets/src/core"
	repositories "github.com/vicpoo/ApiPolarpets/src/skins/domain"
	"github.com/vicpoo/ApiPolarpets/src/skins/domain/entities"
)

type MySQLSkinsRepository struct {
	conn *sql.DB
}

func NewMySQLSkinsRepository() repositories.ISkins {
	conn := core.GetBD()
	return &MySQLSkinsRepository{conn: conn}
}

// Save - Guardar una nueva skin
func (mysql *MySQLSkinsRepository) Save(skin *entities.Skins) error {
	query := `
		INSERT INTO skins (id_tipo_mascota, nombre, imagen_url)
		VALUES (?, ?, ?)
	`
	result, err := mysql.conn.Exec(query, 
		skin.GetIDTipoMascota(), 
		skin.GetNombre(), 
		skin.GetImagenURL())
	if err != nil {
		log.Println("Error al guardar la skin:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	skin.SetIDSkins(int32(id))

	return nil
}

// Update - Actualizar una skin existente
func (mysql *MySQLSkinsRepository) Update(skin *entities.Skins) error {
	query := `
		UPDATE skins
		SET id_tipo_mascota = ?, nombre = ?, imagen_url = ?
		WHERE id_skins = ?
	`
	result, err := mysql.conn.Exec(query, 
		skin.GetIDTipoMascota(), 
		skin.GetNombre(), 
		skin.GetImagenURL(), 
		skin.GetIDSkins())
	if err != nil {
		log.Println("Error al actualizar la skin:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("skin con ID %d no encontrada", skin.GetIDSkins())
	}

	return nil
}

// Delete - Eliminar una skin por ID
func (mysql *MySQLSkinsRepository) Delete(id int32) error {
	query := "DELETE FROM skins WHERE id_skins = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar la skin:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("skin con ID %d no encontrada", id)
	}

	return nil
}

// GetById - Obtener una skin por ID
func (mysql *MySQLSkinsRepository) GetById(id int32) (*entities.Skins, error) {
	query := `
		SELECT id_skins, id_tipo_mascota, nombre, imagen_url
		FROM skins
		WHERE id_skins = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var skin entities.Skins
	var idSkins int32
	var idTipoMascota int32
	var nombre string
	var imagenURL string

	err := row.Scan(&idSkins, &idTipoMascota, &nombre, &imagenURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("skin con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la skin por ID:", err)
		return nil, err
	}

	skin.SetIDSkins(idSkins)
	skin.SetIDTipoMascota(idTipoMascota)
	skin.SetNombre(nombre)
	skin.SetImagenURL(imagenURL)

	return &skin, nil
}

// GetAll - Obtener todas las skins
func (mysql *MySQLSkinsRepository) GetAll() ([]entities.Skins, error) {
	query := `
		SELECT id_skins, id_tipo_mascota, nombre, imagen_url
		FROM skins
		ORDER BY nombre ASC
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todas las skins:", err)
		return nil, err
	}
	defer rows.Close()

	var skins []entities.Skins
	for rows.Next() {
		var skin entities.Skins
		var idSkins int32
		var idTipoMascota int32
		var nombre string
		var imagenURL string

		err := rows.Scan(&idSkins, &idTipoMascota, &nombre, &imagenURL)
		if err != nil {
			log.Println("Error al escanear la skin:", err)
			return nil, err
		}

		skin.SetIDSkins(idSkins)
		skin.SetIDTipoMascota(idTipoMascota)
		skin.SetNombre(nombre)
		skin.SetImagenURL(imagenURL)

		skins = append(skins, skin)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return skins, nil
}

// GetByTipoMascota - Obtener todas las skins de un tipo de mascota
func (mysql *MySQLSkinsRepository) GetByTipoMascota(idTipoMascota int32) ([]entities.Skins, error) {
	query := `
		SELECT id_skins, id_tipo_mascota, nombre, imagen_url
		FROM skins
		WHERE id_tipo_mascota = ?
		ORDER BY nombre ASC
	`
	rows, err := mysql.conn.Query(query, idTipoMascota)
	if err != nil {
		log.Println("Error al obtener skins por tipo de mascota:", err)
		return nil, err
	}
	defer rows.Close()

	var skins []entities.Skins
	for rows.Next() {
		var skin entities.Skins
		var idSkins int32
		var idTipoMascotaValue int32
		var nombre string
		var imagenURL string

		err := rows.Scan(&idSkins, &idTipoMascotaValue, &nombre, &imagenURL)
		if err != nil {
			log.Println("Error al escanear la skin:", err)
			return nil, err
		}

		skin.SetIDSkins(idSkins)
		skin.SetIDTipoMascota(idTipoMascotaValue)
		skin.SetNombre(nombre)
		skin.SetImagenURL(imagenURL)

		skins = append(skins, skin)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return skins, nil
}

// GetByNombre - Obtener una skin por nombre
func (mysql *MySQLSkinsRepository) GetByNombre(nombre string) (*entities.Skins, error) {
	query := `
		SELECT id_skins, id_tipo_mascota, nombre, imagen_url
		FROM skins
		WHERE nombre = ?
	`
	row := mysql.conn.QueryRow(query, nombre)

	var skin entities.Skins
	var idSkins int32
	var idTipoMascota int32
	var nombreValue string
	var imagenURL string

	err := row.Scan(&idSkins, &idTipoMascota, &nombreValue, &imagenURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("skin con nombre '%s' no encontrada", nombre)
		}
		log.Println("Error al buscar la skin por nombre:", err)
		return nil, err
	}

	skin.SetIDSkins(idSkins)
	skin.SetIDTipoMascota(idTipoMascota)
	skin.SetNombre(nombreValue)
	skin.SetImagenURL(imagenURL)

	return &skin, nil
}

// GetByTipoMascotaAndNombre - Obtener una skin por tipo de mascota y nombre
func (mysql *MySQLSkinsRepository) GetByTipoMascotaAndNombre(idTipoMascota int32, nombre string) (*entities.Skins, error) {
	query := `
		SELECT id_skins, id_tipo_mascota, nombre, imagen_url
		FROM skins
		WHERE id_tipo_mascota = ? AND nombre = ?
	`
	row := mysql.conn.QueryRow(query, idTipoMascota, nombre)

	var skin entities.Skins
	var idSkins int32
	var idTipoMascotaValue int32
	var nombreValue string
	var imagenURL string

	err := row.Scan(&idSkins, &idTipoMascotaValue, &nombreValue, &imagenURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("skin con tipo mascota %d y nombre '%s' no encontrada", idTipoMascota, nombre)
		}
		log.Println("Error al buscar la skin por tipo y nombre:", err)
		return nil, err
	}

	skin.SetIDSkins(idSkins)
	skin.SetIDTipoMascota(idTipoMascotaValue)
	skin.SetNombre(nombreValue)
	skin.SetImagenURL(imagenURL)

	return &skin, nil
}