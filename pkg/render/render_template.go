package render

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("web/template/create.html", "web/template/update.html", "web/template/product_list.html",
	"web/template/product_list_cate.html", "web/template/product_list_subcate.html",
	"web/template/product_list_cate1.html", "web/template/product_list_cate2.html",
	"web/template/pdp.html"))

func RenderTemplate[T any](w http.ResponseWriter, tmpl string, p T) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
