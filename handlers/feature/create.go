package feature

import (
	"net/http"

	"github.com/noac178/1stProjectGolang/pkg/render"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "edit", nil)
}
