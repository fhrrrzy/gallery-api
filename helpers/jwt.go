// helpers/jwt.go
package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretKey = []byte("tentusajasecretkey") // Change this to a secure secret key

// GenerateJWTToken generates a JWT token for the given email and username
func GenerateJWTToken(email, username string) (string, error) {
	// Set up the token claims
	claims := jwt.MapClaims{
		"email":    email,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (e.g., 24 hours)
	}

	// Create the token with the claims and sign it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
