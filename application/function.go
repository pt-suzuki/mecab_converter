package function

import (
	"github.com/pt-suzuki/mecab_converter/application/infrastructure/server"
	"net/http"
)

func MecabConverterFunction (w http.ResponseWriter, r *http.Request) {
	echo := server.CreateEcho()
	echo.ServeHTTP(w, r)
}
