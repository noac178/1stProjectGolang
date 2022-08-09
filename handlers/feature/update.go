package feature

import (
	"net/http"

	"github.com/noac178/1stProjectGolang/entity"
	"github.com/noac178/1stProjectGolang/pkg/errorx"
	"github.com/noac178/1stProjectGolang/pkg/render"
	"github.com/noac178/1stProjectGolang/repo"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	product_id := r.URL.Path[len("/update/"):]

	db, err := repo.OpenDb()
	errorx.CheckErr(err)

	var p entity.ProductInfo

	query := `SELECT id, sku, name, price, number, cate_report, sub_cate_report, 
			cate1, cate2, color, size, brand, image FROM product_info WHERE id = ?`
	err1 := db.QueryRow(query, product_id).Scan(&p.Id, &p.Sku, &p.Name, &p.Price, &p.Number, &p.CateReport, &p.SubCateReport,
		&p.Cate1, &p.Cate2, &p.Color, &p.Size, &p.Brand, &p.Image)
	errorx.CheckErr(err1)

	render.RenderTemplate(w, "update", p)
}
