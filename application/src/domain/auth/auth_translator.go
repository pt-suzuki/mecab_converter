package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"strings"
)

type Translator interface {
	TranslateRequestToAuthId(context echo.Context) string
}

type translator struct {
}

func TranslatorImpl() Translator {
	return &translator{}
}

func (t *translator) TranslateRequestToAuthId(context echo.Context) string {
	auth := context.Request().Header.Get("Authorization")
	splitAuth := strings.Split(auth, " ")
	if len(splitAuth) != 2 {
		return ""
	}
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(splitAuth[1], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	sid, ok := claims["sub"].(string)

	if ok {
		return sid
	}

	sid, ok = claims["user_id"].(string)
	if ok {
		return sid
	}

	return "0"
}

