// helpers/jwt.go
package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretKey = []byte("tentusajasecretkey") // Change this to a secure secret key

// GenerateJWTToken generates a JWT token for the given email, username, and userID
func GenerateJWTToken(email, username, userID string) (string, error) {
	// Set up the token claims
	claims := jwt.MapClaims{
		"email":    email,
		"username": username,
		"userID":   userID,
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

// DecodeJWTToken decodes a JWT token and returns the claims
func DecodeJWTToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}