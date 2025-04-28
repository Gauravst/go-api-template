package jwtToken

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyJwtAndGetData(jwtToken string, key string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token uses the correct signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Convert key (string) to []byte for HS256
		return []byte(key), nil
	})

	// Extract claims regardless of token validity
	var claims jwt.MapClaims
	if token != nil {
		if data, ok := token.Claims.(jwt.MapClaims); ok {
			claims = data
		}
	}

	// Check if there's an error, but still return claims if possible
	if err != nil {
		// If token expired, return claims with specific error
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, fmt.Errorf("invalid token signature")
		}
		if claims != nil {
			return claims, fmt.Errorf("token has expired")
		}
		return nil, err
	}

	return claims, nil
}

func CreateNewToken(data interface{}, key string) (string, error) {
	claims, ok := data.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims type")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func SetAccessToken(w http.ResponseWriter, token string, secure bool) {
	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})
}
