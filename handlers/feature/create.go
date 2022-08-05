package feature

import (
	"net/http"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "edit", nil)
}
