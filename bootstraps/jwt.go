package bootstraps

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/oktapascal/app-barayya/models/web"
	"time"
)

type Jwt interface {
	GetJwtKey() string
	GenerateAccessToken(user web.UserResponses) (string, time.Time, error)
}

type JwtClaims struct {
	IdUser     uint   `json:"id_user"`
	KodeLokasi string `json:"kode_lokasi"`
	Role       string `json:"role"`
	jwt.RegisteredClaims
}

type JwtImpl struct {
	Config Config
}

func NewJwtImpl(config Config) *JwtImpl {
	return &JwtImpl{Config: config}
}

func (config *JwtImpl) GetJwtKey() string {
	return config.Config.Get("JWT_KEY")
}

func (config *JwtImpl) GenerateAccessToken(user web.UserResponses) (string, time.Time, error) {
	expiration := time.Now().Add(8 * time.Hour) // 8 jam

	return config.generateToken(user, expiration, []byte(config.GetJwtKey()))
}

func (config *JwtImpl) generateToken(user web.UserResponses, expiration time.Time, secret []byte) (string, time.Time, error) {
	claims := &JwtClaims{
		IdUser:     user.Id,
		KodeLokasi: user.KodeLokasi,
		Role:       user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expiration},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(secret)

	if err != nil {
		return "", time.Now(), err
	}

	return tokenStr, expiration, nil
}
