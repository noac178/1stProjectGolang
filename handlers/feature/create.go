package feature

import (
	"net/http"

	"github.com/noac178/1stProjectGolang/pkg/render"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	render.RenderTemplate(w, "edit", data)
}
