package util

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"jarvan/src/pkg/setting"
	"time"
)

var (
	privateKey         = []byte(setting.AppSetting.JwtSecret)
	issuer             = setting.AppSetting.JwtIssuer
	expireTime         = 3 * time.Hour
	TokenInvalid error = errors.New("Couldn't handle this token \r")
)

type Claims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// generate a token
func GenerateToken(email, password string) (string, error) {
	now := time.Now()
	expireTime := now.Add(expireTime)

	claims := Claims{
		email,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(privateKey)

	return token, err
}

// parse token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func RefreshToken(refreshToken string, expireTime time.Duration) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(refreshToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now

		claims.StandardClaims.ExpiresAt = time.Now().Add(expireTime).Unix()

		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err := tokenClaims.SignedString(privateKey)
		return token, err
	}

	return "", TokenInvalid
}
