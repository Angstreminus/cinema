package auth

import (
	"strconv"
	"time"

	"github.com/Angstreminus/cinema/config"
	"github.com/Angstreminus/cinema/internal/apperrors"
	"github.com/Angstreminus/cinema/internal/dto"
	"github.com/Angstreminus/cinema/internal/entity"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type CustomToken struct {
	UUID  uuid.UUID `json:"uuid"`
	Login string    `json:"login"`
	Role  string    `json:"role"`
	jwt.RegisteredClaims
}

func CreateToken(cfg *config.Config, user *entity.User) (string, apperrors.AppError) {
	tokenExpTime, err := strconv.Atoi(cfg.AccExp)
	if err != nil {
		return "", &apperrors.TokenError{
			Message: err.Error(),
		}
	}
	claims := &CustomToken{
		Login: user.Login,
		UUID:  user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Minute * time.Duration(tokenExpTime)),
			},
		},
	}
	templ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := templ.SignedString([]byte(cfg.AccSecr))
	if err != nil {
		return "", err
	}
	return token, err
}

func IsAuthorized(token string, cfg *config.Config) (bool, apperrors.AppError) {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &apperrors.AuthError{
				Message: "INVALID SIGNING METHOD",
			}
		}
		return 0, nil
	})
	if err != nil {
		return false, &apperrors.AuthError{
			Message: err.Error(),
		}
	}
	return true, nil
}

func ExtractFromToken(reqtoken string, cfg *config.Config) (*dto.UserSignature, apperrors.AppError) {
	token, err := jwt.Parse(reqtoken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &apperrors.AuthError{
				Message: "INVALID SIGNING METHOD",
			}
		}
		return []byte(cfg.AccSecr), nil
	})

	if err != nil {
		return nil, &apperrors.AuthError{
			Message: err.Error(),
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, &apperrors.AuthError{
			Message: "INVALID TOKEN",
		}
	}
	var usrSign dto.UserSignature
	usrSign.Login = claims["login"].(string)
	usrSign.Password = claims["password"].(string)
	return &usrSign, nil
}

