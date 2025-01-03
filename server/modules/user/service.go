package users

import (
	"an-overengineered-app/internal/config"
	models "an-overengineered-app/modules/user/models"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTBody struct {
	models.UserPayload
	jwt.RegisteredClaims
}

func CreateJWT(ctx context.Context, data models.User) (string, error) {

	claims := JWTBody{
		UserPayload: models.UserPayload{
			ID:            data.ID,
			Username:      data.Username,
			Email:         data.Email,
			FirstName:     data.FirstName,
			LastName:      data.LastName,
			AccountStatus: data.AccountStatus,
			Avatar:        data.Avatar,
			UserTimezone:  data.UserTimezone,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.AppConfig.JWTExpiry))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.AppConfig.JwtSecret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
