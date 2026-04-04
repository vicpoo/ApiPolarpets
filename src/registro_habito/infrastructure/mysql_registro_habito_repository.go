// mysql_registro_habito_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/vicpoo/ApiPolarpets/src/core"
	repositories "github.com/vicpoo/ApiPolarpets/src/registro_habito/domain"
	"github.com/vicpoo/ApiPolarpets/src/registro_habito/domain/entities"
)

type MySQLRegistroHabitoRepository struct {
	conn *sql.DB
}

func NewMySQLRegistroHabitoRepository() repositories.IRegistroHabito {
	conn := core.GetBD()
	return &MySQLRegistroHabitoRepository{conn: conn}
}

// ========== CRUD BÁSICO ==========

// Save - Guardar un nuevo registro de hábito
func (mysql *MySQLRegistroHabitoRepository) Save(registro *entities.RegistroHabito) error {
	query := `
		INSERT INTO registro_habito (id_habito, fecha_realizada, puntos_generados)
		VALUES (?, ?, ?)
	`
	result, err := mysql.conn.Exec(query,
		registro.GetIDHabito(),
		registro.GetFechaRealizada(),
		registro.GetPuntosGenerados(),
	)
	if err != nil {
		log.Println("Error al guardar el registro de hábito:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	registro.SetIDRegistroHabito(int32(id))

	return nil
}

// Update - Actualizar un registro de hábito existente
func (mysql *MySQLRegistroHabitoRepository) Update(registro *entities.RegistroHabito) error {
	query := `
		UPDATE registro_habito
		SET id_habito = ?, fecha_realizada = ?, puntos_generados = ?
		WHERE id_registro_habito = ?
	`
	result, err := mysql.conn.Exec(query,
		registro.GetIDHabito(),
		registro.GetFechaRealizada(),
		registro.GetPuntosGenerados(),
		registro.GetIDRegistroHabito(),
	)
	if err != nil {
		log.Println("Error al actualizar el registro de hábito:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("registro de hábito con ID %d no encontrado", registro.GetIDRegistroHabito())
	}

	return nil
}

// Delete - Eliminar un registro de hábito por ID
func (mysql *MySQLRegistroHabitoRepository) Delete(id int32) error {
	query := "DELETE FROM registro_habito WHERE id_registro_habito = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el registro de hábito:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("registro de hábito con ID %d no encontrado", id)
	}

	return nil
}

// GetById - Obtener un registro de hábito por ID
func (mysql *MySQLRegistroHabitoRepository) GetById(id int32) (*entities.RegistroHabito, error) {
	query := `
		SELECT id_registro_habito, id_habito, fecha_realizada, puntos_generados
		FROM registro_habito
		WHERE id_registro_habito = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var registro entities.RegistroHabito
	var idRegistro int32
	var idHabito int32
	var fechaRealizada time.Time
	var puntosGenerados int32

	err := row.Scan(&idRegistro, &idHabito, &fechaRealizada, &puntosGenerados)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("registro de hábito con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el registro de hábito por ID:", err)
		return nil, err
	}

	registro.SetIDRegistroHabito(idRegistro)
	registro.SetIDHabito(idHabito)
	registro.SetFechaRealizada(fechaRealizada)
	registro.SetPuntosGenerados(puntosGenerados)

	return &registro, nil
}

// GetAll - Obtener todos los registros de hábitos
func (mysql *MySQLRegistroHabitoRepository) GetAll() ([]entities.RegistroHabito, error) {
	query := `
		SELECT id_registro_habito, id_habito, fecha_realizada, puntos_generados
		FROM registro_habito
		ORDER BY fecha_realizada DESC
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los registros de hábitos:", err)
		return nil, err
	}
	defer rows.Close()

	var registros []entities.RegistroHabito
	for rows.Next() {
		var registro entities.RegistroHabito
		var idRegistro int32
		var idHabito int32
		var fechaRealizada time.Time
		var puntosGenerados int32

		err := rows.Scan(&idRegistro, &idHabito, &fechaRealizada, &puntosGenerados)
		if err != nil {
			log.Println("Error al escanear el registro de hábito:", err)
			return nil, err
		}

		registro.SetIDRegistroHabito(idRegistro)
		registro.SetIDHabito(idHabito)
		registro.SetFechaRealizada(fechaRealizada)
		registro.SetPuntosGenerados(puntosGenerados)

		registros = append(registros, registro)
	}

	return registros, nil
}

// ========== MÉTODOS ADICIONALES ==========

// GetByHabito - Obtener todos los registros de un hábito específico
func (mysql *MySQLRegistroHabitoRepository) GetByHabito(idHabito int32) ([]entities.RegistroHabito, error) {
	query := `
		SELECT id_registro_habito, id_habito, fecha_realizada, puntos_generados
		FROM registro_habito
		WHERE id_habito = ?
		ORDER BY fecha_realizada DESC
	`
	rows, err := mysql.conn.Query(query, idHabito)
	if err != nil {
		log.Println("Error al obtener registros por hábito:", err)
		return nil, err
	}
	defer rows.Close()

	var registros []entities.RegistroHabito
	for rows.Next() {
		var registro entities.RegistroHabito
		var idRegistro int32
		var idHabitoValue int32
		var fechaRealizada time.Time
		var puntosGenerados int32

		err := rows.Scan(&idRegistro, &idHabitoValue, &fechaRealizada, &puntosGenerados)
		if err != nil {
			log.Println("Error al escanear el registro:", err)
			return nil, err
		}

		registro.SetIDRegistroHabito(idRegistro)
		registro.SetIDHabito(idHabitoValue)
		registro.SetFechaRealizada(fechaRealizada)
		registro.SetPuntosGenerados(puntosGenerados)

		registros = append(registros, registro)
	}

	return registros, nil
}

// GetByHabitoAndFecha - Verificar si un hábito fue completado en una fecha específica
func (mysql *MySQLRegistroHabitoRepository) GetByHabitoAndFecha(idHabito int32, fecha time.Time) (*entities.RegistroHabito, error) {
	query := `
		SELECT id_registro_habito, id_habito, fecha_realizada, puntos_generados
		FROM registro_habito
		WHERE id_habito = ? AND DATE(fecha_realizada) = DATE(?)
		LIMIT 1
	`
	row := mysql.conn.QueryRow(query, idHabito, fecha)

	var registro entities.RegistroHabito
	var idRegistro int32
	var idHabitoValue int32
	var fechaRealizada time.Time
	var puntosGenerados int32

	err := row.Scan(&idRegistro, &idHabitoValue, &fechaRealizada, &puntosGenerados)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No encontrado, no es error
		}
		log.Println("Error al buscar registro por hábito y fecha:", err)
		return nil, err
	}

	registro.SetIDRegistroHabito(idRegistro)
	registro.SetIDHabito(idHabitoValue)
	registro.SetFechaRealizada(fechaRealizada)
	registro.SetPuntosGenerados(puntosGenerados)

	return &registro, nil
}

// GetByFechaRange - Obtener registros en un rango de fechas
func (mysql *MySQLRegistroHabitoRepository) GetByFechaRange(fechaInicio, fechaFin time.Time) ([]entities.RegistroHabito, error) {
	query := `
		SELECT id_registro_habito, id_habito, fecha_realizada, puntos_generados
		FROM registro_habito
		WHERE DATE(fecha_realizada) BETWEEN DATE(?) AND DATE(?)
		ORDER BY fecha_realizada DESC
	`
	rows, err := mysql.conn.Query(query, fechaInicio, fechaFin)
	if err != nil {
		log.Println("Error al obtener registros por rango de fechas:", err)
		return nil, err
	}
	defer rows.Close()

	var registros []entities.RegistroHabito
	for rows.Next() {
		var registro entities.RegistroHabito
		var idRegistro int32
		var idHabito int32
		var fechaRealizada time.Time
		var puntosGenerados int32

		err := rows.Scan(&idRegistro, &idHabito, &fechaRealizada, &puntosGenerados)
		if err != nil {
			log.Println("Error al escanear el registro:", err)
			return nil, err
		}

		registro.SetIDRegistroHabito(idRegistro)
		registro.SetIDHabito(idHabito)
		registro.SetFechaRealizada(fechaRealizada)
		registro.SetPuntosGenerados(puntosGenerados)

		registros = append(registros, registro)
	}

	return registros, nil
}

// GetByUser - Obtener todos los registros de un usuario
func (mysql *MySQLRegistroHabitoRepository) GetByUser(idUser int32) ([]entities.RegistroHabito, error) {
	query := `
		SELECT rh.id_registro_habito, rh.id_habito, rh.fecha_realizada, rh.puntos_generados
		FROM registro_habito rh
		INNER JOIN habito h ON rh.id_habito = h.id_habito
		WHERE h.id_user = ?
		ORDER BY rh.fecha_realizada DESC
	`
	rows, err := mysql.conn.Query(query, idUser)
	if err != nil {
		log.Println("Error al obtener registros por usuario:", err)
		return nil, err
	}
	defer rows.Close()

	var registros []entities.RegistroHabito
	for rows.Next() {
		var registro entities.RegistroHabito
		var idRegistro int32
		var idHabito int32
		var fechaRealizada time.Time
		var puntosGenerados int32

		err := rows.Scan(&idRegistro, &idHabito, &fechaRealizada, &puntosGenerados)
		if err != nil {
			log.Println("Error al escanear el registro:", err)
			return nil, err
		}

		registro.SetIDRegistroHabito(idRegistro)
		registro.SetIDHabito(idHabito)
		registro.SetFechaRealizada(fechaRealizada)
		registro.SetPuntosGenerados(puntosGenerados)

		registros = append(registros, registro)
	}

	return registros, nil
}

// GetTotalPuntosByHabito - Suma total de puntos generados por un hábito
func (mysql *MySQLRegistroHabitoRepository) GetTotalPuntosByHabito(idHabito int32) (int32, error) {
	query := `
		SELECT COALESCE(SUM(puntos_generados), 0)
		FROM registro_habito
		WHERE id_habito = ?
	`
	var totalPuntos int32
	err := mysql.conn.QueryRow(query, idHabito).Scan(&totalPuntos)
	if err != nil {
		log.Println("Error al obtener total de puntos por hábito:", err)
		return 0, err
	}
	return totalPuntos, nil
}

// GetTotalPuntosByUser - Suma total de puntos generados por un usuario
func (mysql *MySQLRegistroHabitoRepository) GetTotalPuntosByUser(idUser int32) (int32, error) {
	query := `
		SELECT COALESCE(SUM(rh.puntos_generados), 0)
		FROM registro_habito rh
		INNER JOIN habito h ON rh.id_habito = h.id_habito
		WHERE h.id_user = ?
	`
	var totalPuntos int32
	err := mysql.conn.QueryRow(query, idUser).Scan(&totalPuntos)
	if err != nil {
		log.Println("Error al obtener total de puntos por usuario:", err)
		return 0, err
	}
	return totalPuntos, nil
}

// GetRegistroCompleto - Obtener registro con detalles del hábito y usuario
func (mysql *MySQLRegistroHabitoRepository) GetRegistroCompleto(idRegistro int32) (*repositories.RegistroHabitoDetalles, error) {
	query := `
		SELECT 
			rh.id_registro_habito,
			rh.id_habito,
			h.titulo,
			h.descripcion,
			h.id_user,
			u.username,
			rh.fecha_realizada,
			rh.puntos_generados
		FROM registro_habito rh
		INNER JOIN habito h ON rh.id_habito = h.id_habito
		INNER JOIN usuarios u ON h.id_user = u.id_usuario
		WHERE rh.id_registro_habito = ?
	`
	row := mysql.conn.QueryRow(query, idRegistro)

	var detalles repositories.RegistroHabitoDetalles
	err := row.Scan(
		&detalles.IDRegistroHabito,
		&detalles.IDHabito,
		&detalles.TituloHabito,
		&detalles.DescripcionHabito,
		&detalles.IDUser,
		&detalles.Username,
		&detalles.FechaRealizada,
		&detalles.PuntosGenerados,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("registro con ID %d no encontrado", idRegistro)
		}
		log.Println("Error al obtener registro completo:", err)
		return nil, err
	}

	return &detalles, nil
}

// ========== MÉTODOS PARA EL FRONTEND ==========

// GetHabitosConEstadoByFecha - Obtiene hábitos de un usuario con estado completado/no completado
func (mysql *MySQLRegistroHabitoRepository) GetHabitosConEstadoByFecha(idUser int32, fecha time.Time) ([]entities.HabitoConEstado, error) {
	query := `
		SELECT 
			h.id_habito,
			h.titulo,
			h.descripcion,
			h.puntos,
			CASE 
				WHEN rh.id_registro_habito IS NOT NULL THEN 1
				ELSE 0
			END AS completado,
			rh.fecha_realizada,
			rh.puntos_generados
		FROM habito h
		LEFT JOIN registro_habito rh 
			ON h.id_habito = rh.id_habito 
			AND DATE(rh.fecha_realizada) = DATE(?)
		WHERE h.id_user = ?
		ORDER BY completado ASC, h.titulo ASC
	`
	
	rows, err := mysql.conn.Query(query, fecha, idUser)
	if err != nil {
		log.Println("Error al obtener hábitos con estado:", err)
		return nil, err
	}
	defer rows.Close()

	var habitos []entities.HabitoConEstado
	for rows.Next() {
		var habito entities.HabitoConEstado
		var fechaRealizada sql.NullTime
		var puntosGenerados sql.NullInt32
		var completado int

		err := rows.Scan(
			&habito.IDHabito,
			&habito.Titulo,
			&habito.Descripcion,
			&habito.Puntos,
			&completado,
			&fechaRealizada,
			&puntosGenerados,
		)
		if err != nil {
			log.Println("Error al escanear hábito:", err)
			return nil, err
		}

		habito.Completado = completado == 1
		if fechaRealizada.Valid {
			habito.FechaRealizada = &fechaRealizada.Time
		}
		if puntosGenerados.Valid {
			habito.PuntosGenerados = &puntosGenerados.Int32
		}

		habitos = append(habitos, habito)
	}

	return habitos, nil
}

// ExistsRegistroHoy - Verifica si un hábito ya fue completado hoy
func (mysql *MySQLRegistroHabitoRepository) ExistsRegistroHoy(idHabito int32) (bool, error) {
	query := `
		SELECT COUNT(*) > 0
		FROM registro_habito
		WHERE id_habito = ? AND DATE(fecha_realizada) = CURDATE()
	`
	var exists bool
	err := mysql.conn.QueryRow(query, idHabito).Scan(&exists)
	if err != nil {
		log.Println("Error al verificar registro hoy:", err)
		return false, err
	}
	return exists, nil
}

// CompletarHabito - Registra que un usuario completó un hábito
func (mysql *MySQLRegistroHabitoRepository) CompletarHabito(idHabito int32, idUser int32, puntos int32) error {
	query := `
		INSERT INTO registro_habito (id_habito, fecha_realizada, puntos_generados)
		VALUES (?, NOW(), ?)
	`
	_, err := mysql.conn.Exec(query, idHabito, puntos)
	if err != nil {
		log.Println("Error al completar hábito:", err)
		return err
	}
	return nil
}