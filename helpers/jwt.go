// helpers/jwt.go
package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
	"strings"
	"strconv"
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

// ExtractUserIDFromToken extracts the user ID from the JWT token in the Authorization header
func ExtractUserIDFromToken(r *http.Request) (uint64, error) {
	tokenString := ExtractTokenFromHeader(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, jwt.ErrSignatureInvalid
	}

	userID, err := parseUserIDFromClaims(claims)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func ExtractTokenFromHeader(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func parseUserIDFromClaims(claims jwt.MapClaims) (uint64, error) {
	userID, ok := claims["userID"].(string)
	if !ok {
		return 0, jwt.ErrSignatureInvalid
	}

	parsedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return 0, jwt.ErrSignatureInvalid
	}

	return parsedUserID, nil
}