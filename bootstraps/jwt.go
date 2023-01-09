package bootstraps

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/oktapascal/app-barayya/models/web"
	"time"
)

// Jwt is an interface for generating and validating JSON Web Tokens (JWTs).
type Jwt interface {
	// GetJwtKey returns the secret key used to sign JWTs.
	GetJwtKey() string
	// GenerateAccessToken creates a new JWT with the given user's information.
	GenerateAccessToken(user web.UserResponses) (string, time.Time, error)
}

// JwtImpl is a concrete implementation of the Jwt interface.
type JwtImpl struct {
	// Config is a Config interface used to get the secret key for signing JWTs.
	Config Config
}

// NewJwtImpl creates a new JwtImpl object using the given Config object.
func NewJwtImpl(config Config) *JwtImpl {
	// Return a new JwtImpl object initialized with the given Config object.
	return &JwtImpl{Config: config}
}

// GetJwtKey returns the secret key used to sign JWTs.
func (config *JwtImpl) GetJwtKey() string {
	// Use the Config object to get the JWT secret key.
	return config.Config.Get("JWT_KEY")
}

// GenerateAccessToken creates a new JWT with the given user's information.
func (config *JwtImpl) GenerateAccessToken(user web.UserResponses) (string, time.Time, error) {
	// Set the JWT expiration date to 8 hours from now.
	expiration := time.Now().Add(8 * time.Hour)
	// Generate a new JWT with the given user's information and expiration date.
	return config.generateToken(user, expiration, []byte(config.GetJwtKey()))
}

// generateToken creates a new JWT with the given user's information, expiration date, and secret key.
func (config *JwtImpl) generateToken(user web.UserResponses, expiration time.Time, secret []byte) (string, time.Time, error) {
	// Set the claims for the JWT using the given user's information.
	claims := &web.JwtClaims{
		IdUser:     user.Id,
		KodeLokasi: user.KodeLokasi,
		Role:       user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expiration},
		},
	}
	// Create a new JWT with the given claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the JWT with the given secret key and convert it to a string.
	tokenStr, err := token.SignedString(secret)

	if err != nil {
		// If an error occurred, return an empty string and the current time as the expiration date.
		return "", time.Now(), err
	}

	// Return the signed JWT string and the expiration date.
	return tokenStr, expiration, nil
}
