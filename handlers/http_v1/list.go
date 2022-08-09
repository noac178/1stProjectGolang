package http_v1

import (
	"net/http"

	"github.com/noac178/1stProjectGolang/entity"
	"github.com/noac178/1stProjectGolang/pkg/errorx"
	"github.com/noac178/1stProjectGolang/pkg/render"
	"github.com/noac178/1stProjectGolang/repo"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	db, _ := repo.OpenDb()

	rows, err := db.Query(`SELECT id, name, price, cate_report, sub_cate_report,
						cate1, cate2, image FROM product_info`)
	errorx.CheckErr(err)
	defer rows.Close()

	var ProductList []entity.ProductInfo
	for rows.Next() {
		var p entity.ProductInfo

		err := rows.Scan(&p.Id, &p.Name, &p.Price, &p.CateReport, &p.SubCateReport,
			&p.Cate1, &p.Cate2, &p.Image)
		errorx.CheckErr(err)
		ProductList = append(ProductList, p)
	}
	err = rows.Err()
	errorx.CheckErr(err)

	render.RenderTemplate(w, "product_list", ProductList)
}
