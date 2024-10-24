package pkg

import (
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
)

var tokenAuth = jwtauth.New("HS256", []byte("secretkey"), nil)

func GenerateJWT(userID uuid.UUID, role string) (string, error) {
	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 8).Unix(),
	})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
