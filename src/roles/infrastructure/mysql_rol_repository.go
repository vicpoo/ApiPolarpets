// mysql_rol_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiPolarpets/src/core"
	repositories "github.com/vicpoo/ApiPolarpets/src/roles/domain"
	"github.com/vicpoo/ApiPolarpets/src/roles/domain/entities"
)

type MySQLRolRepository struct {
	conn *sql.DB
}

func NewMySQLRolRepository() repositories.IRol {
	conn := core.GetBD()
	return &MySQLRolRepository{conn: conn}
}

func (mysql *MySQLRolRepository) Save(rol *entities.Rol) error {
	query := `
		INSERT INTO roles (nombre)
		VALUES (?)
	`
	result, err := mysql.conn.Exec(query, rol.GetNombre())
	if err != nil {
		log.Println("Error al guardar el rol:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	rol.SetID(int32(id))

	return nil
}

func (mysql *MySQLRolRepository) Update(rol *entities.Rol) error {
	query := `
		UPDATE roles
		SET nombre = ?
		WHERE id_rol = ?
	`
	result, err := mysql.conn.Exec(query, rol.GetNombre(), rol.GetID())
	if err != nil {
		log.Println("Error al actualizar el rol:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("rol con ID %d no encontrado", rol.GetID())
	}

	return nil
}

func (mysql *MySQLRolRepository) Delete(id int32) error {
	query := "DELETE FROM roles WHERE id_rol = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el rol:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("rol con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MySQLRolRepository) GetById(id int32) (*entities.Rol, error) {
	query := `
		SELECT id_rol, nombre
		FROM roles
		WHERE id_rol = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var rol entities.Rol
	var nombre string
	var idRol int32
	err := row.Scan(&idRol, &nombre)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("rol con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el rol por ID:", err)
		return nil, err
	}

	rol.SetID(idRol)
	rol.SetNombre(nombre)

	return &rol, nil
}

func (mysql *MySQLRolRepository) GetAll() ([]entities.Rol, error) {
	query := `
		SELECT id_rol, nombre
		FROM roles
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los roles:", err)
		return nil, err
	}
	defer rows.Close()

	var roles []entities.Rol
	for rows.Next() {
		var rol entities.Rol
		var idRol int32
		var nombre string
		err := rows.Scan(&idRol, &nombre)
		if err != nil {
			log.Println("Error al escanear el rol:", err)
			return nil, err
		}
		rol.SetID(idRol)
		rol.SetNombre(nombre)
		roles = append(roles, rol)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return roles, nil
}