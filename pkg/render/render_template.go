package internal

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("create.html", "product_list.html", "update.html", "product_list_cate.html",
	"product_list_subcate.html", "product_list_cate1.html", "product_list_cate2.html", "pdp.html"))

func renderTemplate[T any](w http.ResponseWriter, tmpl string, p T) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
