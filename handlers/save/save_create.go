package save

import (
	"net/http"

	"github.com/noac178/1stProjectGolang/repo"
)

func SaveCreateHandler(w http.ResponseWriter, r *http.Request) {
	sku := r.FormValue("sku")
	name := r.FormValue("name")
	price := r.FormValue("price")
	number := r.FormValue("number")
	cate_report := r.FormValue("cate_report")
	sub_cate_report := r.FormValue("sub_cate_report")
	cate1 := r.FormValue("cate1")
	cate2 := r.FormValue("cate2")
	color := r.FormValue("color")
	size := r.FormValue("size")
	brand := r.FormValue("brand")
	image := r.FormValue("image")

	db, _ := repo.OpenDb()
	db.Exec(`INSERT INTO product_info (sku, name, price, number, cate_report, sub_cate_report, cate1, cate2, color, size, brand, image) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		sku, name, price, number, cate_report, sub_cate_report, cate1, cate2, color, size, brand, image)

	http.Redirect(w, r, "/product_list", http.StatusFound)
}
