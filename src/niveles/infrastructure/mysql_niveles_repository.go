// mysql_niveles_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiPolarpets/src/core"
	repositories "github.com/vicpoo/ApiPolarpets/src/niveles/domain"
	"github.com/vicpoo/ApiPolarpets/src/niveles/domain/entities"
)

type MySQLNivelesRepository struct {
	conn *sql.DB
}

func NewMySQLNivelesRepository() repositories.INiveles {
	conn := core.GetBD()
	return &MySQLNivelesRepository{conn: conn}
}

// Save - Guardar un nuevo nivel
func (mysql *MySQLNivelesRepository) Save(nivel *entities.Niveles) error {
	query := `
		INSERT INTO niveles (nivel, exp_requerida)
		VALUES (?, ?)
	`
	result, err := mysql.conn.Exec(query, nivel.GetNivel(), nivel.GetExpRequerida())
	if err != nil {
		log.Println("Error al guardar el nivel:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	nivel.SetIDNiveles(int32(id))

	return nil
}

// Update - Actualizar un nivel existente
func (mysql *MySQLNivelesRepository) Update(nivel *entities.Niveles) error {
	query := `
		UPDATE niveles
		SET nivel = ?, exp_requerida = ?
		WHERE id_niveles = ?
	`
	result, err := mysql.conn.Exec(query, nivel.GetNivel(), nivel.GetExpRequerida(), nivel.GetIDNiveles())
	if err != nil {
		log.Println("Error al actualizar el nivel:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nivel con ID %d no encontrado", nivel.GetIDNiveles())
	}

	return nil
}

// Delete - Eliminar un nivel por ID
func (mysql *MySQLNivelesRepository) Delete(id int32) error {
	query := "DELETE FROM niveles WHERE id_niveles = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el nivel:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nivel con ID %d no encontrado", id)
	}

	return nil
}

// GetById - Obtener un nivel por ID
func (mysql *MySQLNivelesRepository) GetById(id int32) (*entities.Niveles, error) {
	query := `
		SELECT id_niveles, nivel, exp_requerida
		FROM niveles
		WHERE id_niveles = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var nivel entities.Niveles
	var idNiveles int32
	var nivelValue int32
	var expRequerida int32

	err := row.Scan(&idNiveles, &nivelValue, &expRequerida)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("nivel con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el nivel por ID:", err)
		return nil, err
	}

	nivel.SetIDNiveles(idNiveles)
	nivel.SetNivel(nivelValue)
	nivel.SetExpRequerida(expRequerida)

	return &nivel, nil
}

// GetAll - Obtener todos los niveles
func (mysql *MySQLNivelesRepository) GetAll() ([]entities.Niveles, error) {
	query := `
		SELECT id_niveles, nivel, exp_requerida
		FROM niveles
		ORDER BY nivel ASC
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los niveles:", err)
		return nil, err
	}
	defer rows.Close()

	var niveles []entities.Niveles
	for rows.Next() {
		var nivel entities.Niveles
		var idNiveles int32
		var nivelValue int32
		var expRequerida int32

		err := rows.Scan(&idNiveles, &nivelValue, &expRequerida)
		if err != nil {
			log.Println("Error al escanear el nivel:", err)
			return nil, err
		}

		nivel.SetIDNiveles(idNiveles)
		nivel.SetNivel(nivelValue)
		nivel.SetExpRequerida(expRequerida)

		niveles = append(niveles, nivel)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return niveles, nil
}

// GetByNivel - Obtener nivel por número de nivel
func (mysql *MySQLNivelesRepository) GetByNivel(nivel int32) (*entities.Niveles, error) {
	query := `
		SELECT id_niveles, nivel, exp_requerida
		FROM niveles
		WHERE nivel = ?
	`
	row := mysql.conn.QueryRow(query, nivel)

	var nivelEntity entities.Niveles
	var idNiveles int32
	var nivelValue int32
	var expRequerida int32

	err := row.Scan(&idNiveles, &nivelValue, &expRequerida)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("nivel %d no encontrado", nivel)
		}
		log.Println("Error al buscar el nivel por número:", err)
		return nil, err
	}

	nivelEntity.SetIDNiveles(idNiveles)
	nivelEntity.SetNivel(nivelValue)
	nivelEntity.SetExpRequerida(expRequerida)

	return &nivelEntity, nil
}

// GetByExpRequerida - Obtener nivel por experiencia requerida
func (mysql *MySQLNivelesRepository) GetByExpRequerida(expRequerida int32) (*entities.Niveles, error) {
	query := `
		SELECT id_niveles, nivel, exp_requerida
		FROM niveles
		WHERE exp_requerida = ?
	`
	row := mysql.conn.QueryRow(query, expRequerida)

	var nivelEntity entities.Niveles
	var idNiveles int32
	var nivelValue int32
	var expValue int32

	err := row.Scan(&idNiveles, &nivelValue, &expValue)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("nivel con experiencia requerida %d no encontrado", expRequerida)
		}
		log.Println("Error al buscar el nivel por experiencia requerida:", err)
		return nil, err
	}

	nivelEntity.SetIDNiveles(idNiveles)
	nivelEntity.SetNivel(nivelValue)
	nivelEntity.SetExpRequerida(expValue)

	return &nivelEntity, nil
}

// GetNextLevel - Obtener el siguiente nivel según la experiencia actual
func (mysql *MySQLNivelesRepository) GetNextLevel(currentExp int32) (*entities.Niveles, error) {
	query := `
		SELECT id_niveles, nivel, exp_requerida
		FROM niveles
		WHERE exp_requerida > ?
		ORDER BY exp_requerida ASC
		LIMIT 1
	`
	row := mysql.conn.QueryRow(query, currentExp)

	var nivelEntity entities.Niveles
	var idNiveles int32
	var nivelValue int32
	var expRequerida int32

	err := row.Scan(&idNiveles, &nivelValue, &expRequerida)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no hay siguiente nivel para la experiencia %d", currentExp)
		}
		log.Println("Error al obtener el siguiente nivel:", err)
		return nil, err
	}

	nivelEntity.SetIDNiveles(idNiveles)
	nivelEntity.SetNivel(nivelValue)
	nivelEntity.SetExpRequerida(expRequerida)

	return &nivelEntity, nil
}