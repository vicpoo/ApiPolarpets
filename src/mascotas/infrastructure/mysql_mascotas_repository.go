// mysql_mascotas_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiPolarpets/src/core"
	repositories "github.com/vicpoo/ApiPolarpets/src/mascotas/domain"
	"github.com/vicpoo/ApiPolarpets/src/mascotas/domain/entities"
)

type MySQLMascotaRepository struct {
	conn *sql.DB
}

func NewMySQLMascotaRepository() repositories.IMascota {
	conn := core.GetBD()
	return &MySQLMascotaRepository{conn: conn}
}

// Save - Guardar una nueva mascota
func (mysql *MySQLMascotaRepository) Save(mascota *entities.Mascota) error {
	query := `
		INSERT INTO mascotas (id_user, id_niveles, id_skin, id_tipo_mascota, experiencia_actual)
		VALUES (?, ?, ?, ?, ?)
	`
	result, err := mysql.conn.Exec(query,
		mascota.GetIDUser(),
		mascota.GetIDNiveles(),
		mascota.GetIDSkin(),
		mascota.GetIDTipoMascota(),
		mascota.GetExperienciaActual(),
	)
	if err != nil {
		log.Println("Error al guardar la mascota:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	mascota.SetIDMascota(int32(id))

	return nil
}

// Update - Actualizar una mascota existente
func (mysql *MySQLMascotaRepository) Update(mascota *entities.Mascota) error {
	query := `
		UPDATE mascotas
		SET id_user = ?, id_niveles = ?, id_skin = ?, id_tipo_mascota = ?, experiencia_actual = ?
		WHERE id_mascota = ?
	`
	result, err := mysql.conn.Exec(query,
		mascota.GetIDUser(),
		mascota.GetIDNiveles(),
		mascota.GetIDSkin(),
		mascota.GetIDTipoMascota(),
		mascota.GetExperienciaActual(),
		mascota.GetIDMascota(),
	)
	if err != nil {
		log.Println("Error al actualizar la mascota:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("mascota con ID %d no encontrada", mascota.GetIDMascota())
	}

	return nil
}

// Delete - Eliminar una mascota por ID
func (mysql *MySQLMascotaRepository) Delete(id int32) error {
	query := "DELETE FROM mascotas WHERE id_mascota = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar la mascota:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("mascota con ID %d no encontrada", id)
	}

	return nil
}

// GetById - Obtener una mascota por ID
func (mysql *MySQLMascotaRepository) GetById(id int32) (*entities.Mascota, error) {
	query := `
		SELECT id_mascota, id_user, id_niveles, id_skin, id_tipo_mascota, experiencia_actual
		FROM mascotas
		WHERE id_mascota = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var mascota entities.Mascota
	var idMascota int32
	var idUser int32
	var idNiveles int32
	var idSkin int32
	var idTipoMascota int32
	var experienciaActual int32

	err := row.Scan(&idMascota, &idUser, &idNiveles, &idSkin, &idTipoMascota, &experienciaActual)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("mascota con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la mascota por ID:", err)
		return nil, err
	}

	mascota.SetIDMascota(idMascota)
	mascota.SetIDUser(idUser)
	mascota.SetIDNiveles(idNiveles)
	mascota.SetIDSkin(idSkin)
	mascota.SetIDTipoMascota(idTipoMascota)
	mascota.SetExperienciaActual(experienciaActual)

	return &mascota, nil
}

// GetAll - Obtener todas las mascotas
func (mysql *MySQLMascotaRepository) GetAll() ([]entities.Mascota, error) {
	query := `
		SELECT id_mascota, id_user, id_niveles, id_skin, id_tipo_mascota, experiencia_actual
		FROM mascotas
		ORDER BY id_mascota ASC
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todas las mascotas:", err)
		return nil, err
	}
	defer rows.Close()

	var mascotas []entities.Mascota
	for rows.Next() {
		var mascota entities.Mascota
		var idMascota int32
		var idUser int32
		var idNiveles int32
		var idSkin int32
		var idTipoMascota int32
		var experienciaActual int32

		err := rows.Scan(&idMascota, &idUser, &idNiveles, &idSkin, &idTipoMascota, &experienciaActual)
		if err != nil {
			log.Println("Error al escanear la mascota:", err)
			return nil, err
		}

		mascota.SetIDMascota(idMascota)
		mascota.SetIDUser(idUser)
		mascota.SetIDNiveles(idNiveles)
		mascota.SetIDSkin(idSkin)
		mascota.SetIDTipoMascota(idTipoMascota)
		mascota.SetExperienciaActual(experienciaActual)

		mascotas = append(mascotas, mascota)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return mascotas, nil
}

// GetByUser - Obtener todas las mascotas de un usuario
func (mysql *MySQLMascotaRepository) GetByUser(idUser int32) ([]entities.Mascota, error) {
	query := `
		SELECT id_mascota, id_user, id_niveles, id_skin, id_tipo_mascota, experiencia_actual
		FROM mascotas
		WHERE id_user = ?
		ORDER BY id_mascota ASC
	`
	rows, err := mysql.conn.Query(query, idUser)
	if err != nil {
		log.Println("Error al obtener mascotas por usuario:", err)
		return nil, err
	}
	defer rows.Close()

	var mascotas []entities.Mascota
	for rows.Next() {
		var mascota entities.Mascota
		var idMascota int32
		var idUserValue int32
		var idNiveles int32
		var idSkin int32
		var idTipoMascota int32
		var experienciaActual int32

		err := rows.Scan(&idMascota, &idUserValue, &idNiveles, &idSkin, &idTipoMascota, &experienciaActual)
		if err != nil {
			log.Println("Error al escanear la mascota:", err)
			return nil, err
		}

		mascota.SetIDMascota(idMascota)
		mascota.SetIDUser(idUserValue)
		mascota.SetIDNiveles(idNiveles)
		mascota.SetIDSkin(idSkin)
		mascota.SetIDTipoMascota(idTipoMascota)
		mascota.SetExperienciaActual(experienciaActual)

		mascotas = append(mascotas, mascota)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return mascotas, nil
}

// GetByTipoMascota - Obtener mascotas por tipo
func (mysql *MySQLMascotaRepository) GetByTipoMascota(idTipoMascota int32) ([]entities.Mascota, error) {
	query := `
		SELECT id_mascota, id_user, id_niveles, id_skin, id_tipo_mascota, experiencia_actual
		FROM mascotas
		WHERE id_tipo_mascota = ?
		ORDER BY id_mascota ASC
	`
	rows, err := mysql.conn.Query(query, idTipoMascota)
	if err != nil {
		log.Println("Error al obtener mascotas por tipo:", err)
		return nil, err
	}
	defer rows.Close()

	var mascotas []entities.Mascota
	for rows.Next() {
		var mascota entities.Mascota
		var idMascota int32
		var idUser int32
		var idNiveles int32
		var idSkin int32
		var idTipoMascotaValue int32
		var experienciaActual int32

		err := rows.Scan(&idMascota, &idUser, &idNiveles, &idSkin, &idTipoMascotaValue, &experienciaActual)
		if err != nil {
			log.Println("Error al escanear la mascota:", err)
			return nil, err
		}

		mascota.SetIDMascota(idMascota)
		mascota.SetIDUser(idUser)
		mascota.SetIDNiveles(idNiveles)
		mascota.SetIDSkin(idSkin)
		mascota.SetIDTipoMascota(idTipoMascotaValue)
		mascota.SetExperienciaActual(experienciaActual)

		mascotas = append(mascotas, mascota)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return mascotas, nil
}

// GetBySkin - Obtener mascotas por skin
func (mysql *MySQLMascotaRepository) GetBySkin(idSkin int32) ([]entities.Mascota, error) {
	query := `
		SELECT id_mascota, id_user, id_niveles, id_skin, id_tipo_mascota, experiencia_actual
		FROM mascotas
		WHERE id_skin = ?
		ORDER BY id_mascota ASC
	`
	rows, err := mysql.conn.Query(query, idSkin)
	if err != nil {
		log.Println("Error al obtener mascotas por skin:", err)
		return nil, err
	}
	defer rows.Close()

	var mascotas []entities.Mascota
	for rows.Next() {
		var mascota entities.Mascota
		var idMascota int32
		var idUser int32
		var idNiveles int32
		var idSkinValue int32
		var idTipoMascota int32
		var experienciaActual int32

		err := rows.Scan(&idMascota, &idUser, &idNiveles, &idSkinValue, &idTipoMascota, &experienciaActual)
		if err != nil {
			log.Println("Error al escanear la mascota:", err)
			return nil, err
		}

		mascota.SetIDMascota(idMascota)
		mascota.SetIDUser(idUser)
		mascota.SetIDNiveles(idNiveles)
		mascota.SetIDSkin(idSkinValue)
		mascota.SetIDTipoMascota(idTipoMascota)
		mascota.SetExperienciaActual(experienciaActual)

		mascotas = append(mascotas, mascota)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return mascotas, nil
}

// GetByNivel - Obtener mascotas por nivel
func (mysql *MySQLMascotaRepository) GetByNivel(idNiveles int32) ([]entities.Mascota, error) {
	query := `
		SELECT id_mascota, id_user, id_niveles, id_skin, id_tipo_mascota, experiencia_actual
		FROM mascotas
		WHERE id_niveles = ?
		ORDER BY id_mascota ASC
	`
	rows, err := mysql.conn.Query(query, idNiveles)
	if err != nil {
		log.Println("Error al obtener mascotas por nivel:", err)
		return nil, err
	}
	defer rows.Close()

	var mascotas []entities.Mascota
	for rows.Next() {
		var mascota entities.Mascota
		var idMascota int32
		var idUser int32
		var idNivelesValue int32
		var idSkin int32
		var idTipoMascota int32
		var experienciaActual int32

		err := rows.Scan(&idMascota, &idUser, &idNivelesValue, &idSkin, &idTipoMascota, &experienciaActual)
		if err != nil {
			log.Println("Error al escanear la mascota:", err)
			return nil, err
		}

		mascota.SetIDMascota(idMascota)
		mascota.SetIDUser(idUser)
		mascota.SetIDNiveles(idNivelesValue)
		mascota.SetIDSkin(idSkin)
		mascota.SetIDTipoMascota(idTipoMascota)
		mascota.SetExperienciaActual(experienciaActual)

		mascotas = append(mascotas, mascota)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return mascotas, nil
}

// UpdateExperiencia - Actualizar solo la experiencia de una mascota
func (mysql *MySQLMascotaRepository) UpdateExperiencia(idMascota int32, nuevaExperiencia int32) error {
	query := `
		UPDATE mascotas
		SET experiencia_actual = ?
		WHERE id_mascota = ?
	`
	result, err := mysql.conn.Exec(query, nuevaExperiencia, idMascota)
	if err != nil {
		log.Println("Error al actualizar la experiencia:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("mascota con ID %d no encontrada", idMascota)
	}

	return nil
}

// GetMascotaCompleta - Obtener mascota con todos los datos relacionados
func (mysql *MySQLMascotaRepository) GetMascotaCompleta(idMascota int32) (*repositories.MascotaDetalles, error) {
	query := `
		SELECT 
			m.id_mascota,
			m.id_user,
			u.username,
			u.email,
			m.id_niveles,
			n.nivel,
			n.exp_requerida,
			m.id_skin,
			s.nombre as nombre_skin,
			s.imagen_url,
			m.id_tipo_mascota,
			tm.nombre as nombre_tipo_mascota,
			tm.descripcion as descripcion_tipo,
			m.experiencia_actual
		FROM mascotas m
		INNER JOIN usuarios u ON m.id_user = u.id_usuario
		INNER JOIN niveles n ON m.id_niveles = n.id_niveles
		INNER JOIN skins s ON m.id_skin = s.id_skins
		INNER JOIN tipo_mascota tm ON m.id_tipo_mascota = tm.id_tipo_mascota
		WHERE m.id_mascota = ?
	`
	row := mysql.conn.QueryRow(query, idMascota)

	var detalles repositories.MascotaDetalles
	err := row.Scan(
		&detalles.IDMascota,
		&detalles.IDUser,
		&detalles.Username,
		&detalles.Email,
		&detalles.IDNiveles,
		&detalles.Nivel,
		&detalles.ExpRequerida,
		&detalles.IDSkin,
		&detalles.NombreSkin,
		&detalles.ImagenURL,
		&detalles.IDTipoMascota,
		&detalles.NombreTipoMascota,
		&detalles.DescripcionTipo,
		&detalles.ExperienciaActual,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("mascota con ID %d no encontrada", idMascota)
		}
		log.Println("Error al obtener mascota completa:", err)
		return nil, err
	}

	return &detalles, nil
}