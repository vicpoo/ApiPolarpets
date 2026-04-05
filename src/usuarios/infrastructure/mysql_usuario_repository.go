// mysql_usuario_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiPolarpets/src/core"
	repositories "github.com/vicpoo/ApiPolarpets/src/usuarios/domain"
	"github.com/vicpoo/ApiPolarpets/src/usuarios/domain/entities"
)

type MySQLUsuarioRepository struct {
	conn *sql.DB
}

func NewMySQLUsuarioRepository() repositories.IUsuario {
	conn := core.GetBD()
	return &MySQLUsuarioRepository{conn: conn}
}

// Register - Registrar un nuevo usuario
func (mysql *MySQLUsuarioRepository) Register(usuario *entities.Usuario) error {
	query := `
		INSERT INTO usuarios (username, email, password, id_rol, id_mascota_activa)
		VALUES (?, ?, ?, ?, ?)
	`
	result, err := mysql.conn.Exec(query, 
		usuario.GetUsername(), 
		usuario.GetEmail(), 
		usuario.GetPassword(), 
		usuario.GetIDRol(),
		usuario.GetIDMascotaActiva(),
	)
	if err != nil {
		log.Println("Error al registrar el usuario:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	usuario.SetIDUsuario(int32(id))

	return nil
}

// Login - Autenticar usuario por email y password (solo retorna el usuario)
func (mysql *MySQLUsuarioRepository) Login(email, password string) (*entities.Usuario, error) {
	query := `
		SELECT id_usuario, username, email, password, id_rol, id_mascota_activa
		FROM usuarios
		WHERE email = ?
	`
	row := mysql.conn.QueryRow(query, email)

	var usuario entities.Usuario
	var idUsuario int32
	var username string
	var userEmail string
	var userPassword string
	var idRol int32
	var idMascotaActiva sql.NullInt32

	err := row.Scan(&idUsuario, &username, &userEmail, &userPassword, &idRol, &idMascotaActiva)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("credenciales inválidas")
		}
		log.Println("Error al hacer login:", err)
		return nil, err
	}

	usuario.SetIDUsuario(idUsuario)
	usuario.SetUsername(username)
	usuario.SetEmail(userEmail)
	usuario.SetPassword(userPassword)
	usuario.SetIDRol(idRol)
	
	if idMascotaActiva.Valid {
		usuario.SetIDMascotaActiva(&idMascotaActiva.Int32)
	} else {
		usuario.SetIDMascotaActiva(nil)
	}

	// Verificar contraseña
	if !usuario.CheckPassword(password) {
		return nil, fmt.Errorf("credenciales inválidas")
	}

	return &usuario, nil
}

// Save - Guardar usuario (CRUD)
func (mysql *MySQLUsuarioRepository) Save(usuario *entities.Usuario) error {
	query := `
		INSERT INTO usuarios (username, email, password, id_rol, id_mascota_activa)
		VALUES (?, ?, ?, ?, ?)
	`
	result, err := mysql.conn.Exec(query,
		usuario.GetUsername(),
		usuario.GetEmail(),
		usuario.GetPassword(),
		usuario.GetIDRol(),
		usuario.GetIDMascotaActiva(),
	)
	if err != nil {
		log.Println("Error al guardar el usuario:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	usuario.SetIDUsuario(int32(id))

	return nil
}

// Update - Actualizar usuario existente
func (mysql *MySQLUsuarioRepository) Update(usuario *entities.Usuario) error {
	query := `
		UPDATE usuarios
		SET username = ?, email = ?, password = ?, id_rol = ?, id_mascota_activa = ?
		WHERE id_usuario = ?
	`
	result, err := mysql.conn.Exec(query,
		usuario.GetUsername(),
		usuario.GetEmail(),
		usuario.GetPassword(),
		usuario.GetIDRol(),
		usuario.GetIDMascotaActiva(),
		usuario.GetIDUsuario(),
	)
	if err != nil {
		log.Println("Error al actualizar el usuario:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario con ID %d no encontrado", usuario.GetIDUsuario())
	}

	return nil
}

// Delete - Eliminar usuario por ID
func (mysql *MySQLUsuarioRepository) Delete(id int32) error {
	query := "DELETE FROM usuarios WHERE id_usuario = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el usuario:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario con ID %d no encontrado", id)
	}

	return nil
}

// GetById - Obtener usuario por ID
func (mysql *MySQLUsuarioRepository) GetById(id int32) (*entities.Usuario, error) {
	query := `
		SELECT id_usuario, username, email, password, id_rol, id_mascota_activa
		FROM usuarios
		WHERE id_usuario = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var usuario entities.Usuario
	var idUsuario int32
	var username string
	var email string
	var password string
	var idRol int32
	var idMascotaActiva sql.NullInt32

	err := row.Scan(&idUsuario, &username, &email, &password, &idRol, &idMascotaActiva)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el usuario por ID:", err)
		return nil, err
	}

	usuario.SetIDUsuario(idUsuario)
	usuario.SetUsername(username)
	usuario.SetEmail(email)
	usuario.SetPassword("") // Limpiar contraseña
	usuario.SetIDRol(idRol)
	
	if idMascotaActiva.Valid {
		usuario.SetIDMascotaActiva(&idMascotaActiva.Int32)
	} else {
		usuario.SetIDMascotaActiva(nil)
	}

	return &usuario, nil
}

// GetAll - Obtener todos los usuarios
func (mysql *MySQLUsuarioRepository) GetAll() ([]entities.Usuario, error) {
	query := `
		SELECT id_usuario, username, email, password, id_rol, id_mascota_activa
		FROM usuarios
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los usuarios:", err)
		return nil, err
	}
	defer rows.Close()

	var usuarios []entities.Usuario
	for rows.Next() {
		var usuario entities.Usuario
		var idUsuario int32
		var username string
		var email string
		var password string
		var idRol int32
		var idMascotaActiva sql.NullInt32

		err := rows.Scan(&idUsuario, &username, &email, &password, &idRol, &idMascotaActiva)
		if err != nil {
			log.Println("Error al escanear el usuario:", err)
			return nil, err
		}

		usuario.SetIDUsuario(idUsuario)
		usuario.SetUsername(username)
		usuario.SetEmail(email)
		usuario.SetPassword("") // Limpiar contraseña
		usuario.SetIDRol(idRol)
		
		if idMascotaActiva.Valid {
			usuario.SetIDMascotaActiva(&idMascotaActiva.Int32)
		} else {
			usuario.SetIDMascotaActiva(nil)
		}

		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return usuarios, nil
}

// GetByEmail - Obtener usuario por email
func (mysql *MySQLUsuarioRepository) GetByEmail(email string) (*entities.Usuario, error) {
	query := `
		SELECT id_usuario, username, email, password, id_rol, id_mascota_activa
		FROM usuarios
		WHERE email = ?
	`
	row := mysql.conn.QueryRow(query, email)

	var usuario entities.Usuario
	var idUsuario int32
	var username string
	var userEmail string
	var password string
	var idRol int32
	var idMascotaActiva sql.NullInt32

	err := row.Scan(&idUsuario, &username, &userEmail, &password, &idRol, &idMascotaActiva)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario con email %s no encontrado", email)
		}
		log.Println("Error al buscar el usuario por email:", err)
		return nil, err
	}

	usuario.SetIDUsuario(idUsuario)
	usuario.SetUsername(username)
	usuario.SetEmail(userEmail)
	usuario.SetPassword("") // Limpiar contraseña
	usuario.SetIDRol(idRol)
	
	if idMascotaActiva.Valid {
		usuario.SetIDMascotaActiva(&idMascotaActiva.Int32)
	} else {
		usuario.SetIDMascotaActiva(nil)
	}

	return &usuario, nil
}

// UpdateMascotaActiva - Actualizar la mascota activa del usuario
func (mysql *MySQLUsuarioRepository) UpdateMascotaActiva(idUsuario int32, idMascotaActiva *int32) error {
	query := `
		UPDATE usuarios
		SET id_mascota_activa = ?
		WHERE id_usuario = ?
	`
	result, err := mysql.conn.Exec(query, idMascotaActiva, idUsuario)
	if err != nil {
		log.Println("Error al actualizar la mascota activa:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario con ID %d no encontrado", idUsuario)
	}

	return nil
}