package responder

import (
	"github.com/labstack/echo"
	"net/http"
)

func AuthFirebaseCustomTokenCreateResponder (context echo.Context, token string) error {
	return context.JSON(http.StatusOK, token)
}