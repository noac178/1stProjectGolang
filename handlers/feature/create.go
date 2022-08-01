package feature

import (
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "edit", nil)
}
