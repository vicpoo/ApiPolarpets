package core

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func init() {
	if len(jwtSecret) == 0 {
		jwtSecret = []byte("tu-secreto-super-seguro-cambiar-en-produccion")
	}
}

type Claims struct {
	IDUsuario int32  `json:"id_usuario"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	IDRol     int32  `json:"id_rol"`
	jwt.RegisteredClaims
}

// GenerarToken crea un nuevo JWT para el usuario
func GenerarToken(idUsuario int32, email, username string, idRol int32) (string, error) {
	claims := Claims{
		IDUsuario: idUsuario,
		Email:     email,
		Username:  username,
		IDRol:     idRol,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidarToken verifica y decodifica el token JWT
func ValidarToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}