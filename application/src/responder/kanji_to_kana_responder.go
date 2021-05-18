package responder

import (
	"github.com/labstack/echo"
	"net/http"
)

func KanjiToKanaResponder (context echo.Context, word string) error {
	return context.JSON(http.StatusOK, word)
}