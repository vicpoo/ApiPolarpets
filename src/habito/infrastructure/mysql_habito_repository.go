// mysql_habito_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiPolarpets/src/core"
	repositories "github.com/vicpoo/ApiPolarpets/src/habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/habito/domain/entities"
)

type MySQLHabitoRepository struct {
	conn *sql.DB
}

func NewMySQLHabitoRepository() repositories.IHabito {
	conn := core.GetBD()
	return &MySQLHabitoRepository{conn: conn}
}

// Save - Guardar un nuevo hábito
func (mysql *MySQLHabitoRepository) Save(habito *entities.Habito) error {
	query := `
		INSERT INTO habito (id_user, titulo, descripcion, puntos)
		VALUES (?, ?, ?, ?)
	`
	result, err := mysql.conn.Exec(query,
		habito.GetIDUser(),
		habito.GetTitulo(),
		habito.GetDescripcion(),
		habito.GetPuntos(),
	)
	if err != nil {
		log.Println("Error al guardar el hábito:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	habito.SetIDHabito(int32(id))

	return nil
}

// Update - Actualizar un hábito existente
func (mysql *MySQLHabitoRepository) Update(habito *entities.Habito) error {
	query := `
		UPDATE habito
		SET id_user = ?, titulo = ?, descripcion = ?, puntos = ?
		WHERE id_habito = ?
	`
	result, err := mysql.conn.Exec(query,
		habito.GetIDUser(),
		habito.GetTitulo(),
		habito.GetDescripcion(),
		habito.GetPuntos(),
		habito.GetIDHabito(),
	)
	if err != nil {
		log.Println("Error al actualizar el hábito:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("hábito con ID %d no encontrado", habito.GetIDHabito())
	}

	return nil
}

// Delete - Eliminar un hábito por ID
func (mysql *MySQLHabitoRepository) Delete(id int32) error {
	query := "DELETE FROM habito WHERE id_habito = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el hábito:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("hábito con ID %d no encontrado", id)
	}

	return nil
}

// GetById - Obtener un hábito por ID
func (mysql *MySQLHabitoRepository) GetById(id int32) (*entities.Habito, error) {
	query := `
		SELECT id_habito, id_user, titulo, descripcion, puntos
		FROM habito
		WHERE id_habito = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var habito entities.Habito
	var idHabito int32
	var idUser int32
	var titulo string
	var descripcion string
	var puntos int32

	err := row.Scan(&idHabito, &idUser, &titulo, &descripcion, &puntos)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("hábito con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el hábito por ID:", err)
		return nil, err
	}

	habito.SetIDHabito(idHabito)
	habito.SetIDUser(idUser)
	habito.SetTitulo(titulo)
	habito.SetDescripcion(descripcion)
	habito.SetPuntos(puntos)

	return &habito, nil
}

// GetAll - Obtener todos los hábitos
func (mysql *MySQLHabitoRepository) GetAll() ([]entities.Habito, error) {
	query := `
		SELECT id_habito, id_user, titulo, descripcion, puntos
		FROM habito
		ORDER BY id_habito ASC
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los hábitos:", err)
		return nil, err
	}
	defer rows.Close()

	var habitos []entities.Habito
	for rows.Next() {
		var habito entities.Habito
		var idHabito int32
		var idUser int32
		var titulo string
		var descripcion string
		var puntos int32

		err := rows.Scan(&idHabito, &idUser, &titulo, &descripcion, &puntos)
		if err != nil {
			log.Println("Error al escanear el hábito:", err)
			return nil, err
		}

		habito.SetIDHabito(idHabito)
		habito.SetIDUser(idUser)
		habito.SetTitulo(titulo)
		habito.SetDescripcion(descripcion)
		habito.SetPuntos(puntos)

		habitos = append(habitos, habito)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return habitos, nil
}

// GetByUser - Obtener todos los hábitos de un usuario
func (mysql *MySQLHabitoRepository) GetByUser(idUser int32) ([]entities.Habito, error) {
	query := `
		SELECT id_habito, id_user, titulo, descripcion, puntos
		FROM habito
		WHERE id_user = ?
		ORDER BY id_habito ASC
	`
	rows, err := mysql.conn.Query(query, idUser)
	if err != nil {
		log.Println("Error al obtener hábitos por usuario:", err)
		return nil, err
	}
	defer rows.Close()

	var habitos []entities.Habito
	for rows.Next() {
		var habito entities.Habito
		var idHabito int32
		var idUserValue int32
		var titulo string
		var descripcion string
		var puntos int32

		err := rows.Scan(&idHabito, &idUserValue, &titulo, &descripcion, &puntos)
		if err != nil {
			log.Println("Error al escanear el hábito:", err)
			return nil, err
		}

		habito.SetIDHabito(idHabito)
		habito.SetIDUser(idUserValue)
		habito.SetTitulo(titulo)
		habito.SetDescripcion(descripcion)
		habito.SetPuntos(puntos)

		habitos = append(habitos, habito)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return habitos, nil
}

// GetByTitulo - Obtener un hábito por título
func (mysql *MySQLHabitoRepository) GetByTitulo(titulo string) (*entities.Habito, error) {
	query := `
		SELECT id_habito, id_user, titulo, descripcion, puntos
		FROM habito
		WHERE titulo = ?
	`
	row := mysql.conn.QueryRow(query, titulo)

	var habito entities.Habito
	var idHabito int32
	var idUser int32
	var tituloValue string
	var descripcion string
	var puntos int32

	err := row.Scan(&idHabito, &idUser, &tituloValue, &descripcion, &puntos)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("hábito con título '%s' no encontrado", titulo)
		}
		log.Println("Error al buscar el hábito por título:", err)
		return nil, err
	}

	habito.SetIDHabito(idHabito)
	habito.SetIDUser(idUser)
	habito.SetTitulo(tituloValue)
	habito.SetDescripcion(descripcion)
	habito.SetPuntos(puntos)

	return &habito, nil
}

// GetByUserAndTitulo - Obtener un hábito por usuario y título
func (mysql *MySQLHabitoRepository) GetByUserAndTitulo(idUser int32, titulo string) (*entities.Habito, error) {
	query := `
		SELECT id_habito, id_user, titulo, descripcion, puntos
		FROM habito
		WHERE id_user = ? AND titulo = ?
	`
	row := mysql.conn.QueryRow(query, idUser, titulo)

	var habito entities.Habito
	var idHabito int32
	var idUserValue int32
	var tituloValue string
	var descripcion string
	var puntos int32

	err := row.Scan(&idHabito, &idUserValue, &tituloValue, &descripcion, &puntos)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("hábito con usuario %d y título '%s' no encontrado", idUser, titulo)
		}
		log.Println("Error al buscar el hábito por usuario y título:", err)
		return nil, err
	}

	habito.SetIDHabito(idHabito)
	habito.SetIDUser(idUserValue)
	habito.SetTitulo(tituloValue)
	habito.SetDescripcion(descripcion)
	habito.SetPuntos(puntos)

	return &habito, nil
}

// GetTotalPuntosByUser - Obtener la suma total de puntos de un usuario
func (mysql *MySQLHabitoRepository) GetTotalPuntosByUser(idUser int32) (int32, error) {
	query := `
		SELECT COALESCE(SUM(puntos), 0)
		FROM habito
		WHERE id_user = ?
	`
	var totalPuntos int32
	err := mysql.conn.QueryRow(query, idUser).Scan(&totalPuntos)
	if err != nil {
		log.Println("Error al obtener total de puntos:", err)
		return 0, err
	}

	return totalPuntos, nil
}