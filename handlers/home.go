package handlers

import (
	"net/http"

	"github.com/dxps/dreampic/views/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}
