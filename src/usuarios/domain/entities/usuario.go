// usuario.go
package entities

type Usuario struct {
	IDUsuario       int32  `json:"id_usuario" gorm:"column:id_usuario;primaryKey;autoIncrement"`
	Username        string `json:"username" gorm:"column:username"`
	Email           string `json:"email" gorm:"column:email;unique"`
	Password        string `json:"password" gorm:"column:password"`
	IDRol           int32  `json:"id_rol" gorm:"column:id_rol"`
	IDMascotaActiva *int32 `json:"id_mascota_activa,omitempty" gorm:"column:id_mascota_activa"`
}

// Setters
func (u *Usuario) SetIDUsuario(id int32) {
	u.IDUsuario = id
}

func (u *Usuario) SetUsername(username string) {
	u.Username = username
}

func (u *Usuario) SetEmail(email string) {
	u.Email = email
}

func (u *Usuario) SetPassword(password string) {
	u.Password = password
}

func (u *Usuario) SetIDRol(idRol int32) {
	u.IDRol = idRol
}

func (u *Usuario) SetIDMascotaActiva(idMascotaActiva *int32) {
	u.IDMascotaActiva = idMascotaActiva
}

// Getters
func (u *Usuario) GetIDUsuario() int32 {
	return u.IDUsuario
}

func (u *Usuario) GetUsername() string {
	return u.Username
}

func (u *Usuario) GetEmail() string {
	return u.Email
}

func (u *Usuario) GetPassword() string {
	return u.Password
}

func (u *Usuario) GetIDRol() int32 {
	return u.IDRol
}

func (u *Usuario) GetIDMascotaActiva() *int32 {
	return u.IDMascotaActiva
}

// Constructor
func NewUsuario(username, email, password string, idRol int32) *Usuario {
	return &Usuario{
		Username: username,
		Email:    email,
		Password: password,
		IDRol:    idRol,
	}
}