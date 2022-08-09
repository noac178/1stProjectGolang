package http_v1

import (
	"net/http"

	"github.com/noac178/1stProjectGolang/entity"
	"github.com/noac178/1stProjectGolang/pkg/errorx"
	"github.com/noac178/1stProjectGolang/pkg/pagepath"
	"github.com/noac178/1stProjectGolang/pkg/render"
	"github.com/noac178/1stProjectGolang/repo"
)

func FilterCate2Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cate2 := r.Form.Get("cate2")

	http.Redirect(w, r, "/product_list/c2/"+pagepath.CreatePagePath(cate2), http.StatusFound)
}

func Cate2Handler(w http.ResponseWriter, r *http.Request) {
	cate2 := r.URL.Path[len("/product_list/c2/"):]

	db, err := repo.OpenDb()
	errorx.CheckErr(err)

	rows, err := db.Query(`
		SELECT DISTINCT id, name, price, cate_report, sub_cate_report,
				cate1, cate2, image
		FROM product_info 
		WHERE 1=1
		AND LOWER(REPLACE(REPLACE(cate2, ' - ', '-'), ' ', '-')) = ?`, cate2)
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

	query := `
		SELECT
			DISTINCT
			cate_report,
			sub_cate_report,
			cate1,
			cate2
		FROM product_info 
		WHERE 1=1
		AND (
		CASE
			WHEN cate2 NOT LIKE '%-%' THEN LOWER(REPLACE(REPLACE(cate2, ' - ', '-'), ' ', '-')) = ?
			ELSE LOWER(REPLACE(REPLACE(cate2, ' - ', '-'), ' ', '-')) = ?
		END)`

	var p entity.ProductInfo
	err = db.QueryRow(query, cate2, cate2).Scan(&p.CateReport, &p.SubCateReport, &p.Cate1, &p.Cate2)
	errorx.CheckErr(err)

	type Page struct {
		CateEng       string
		CateReport    string
		SubCateEng    string
		SubCateReport string
		Cate1Eng      string
		Cate1         string
		Cate2Eng      string
		Cate2         string
		ProductLists  []entity.ProductInfo
	}

	page := &Page{
		CateEng:       pagepath.CreatePagePath(p.CateReport),
		SubCateEng:    pagepath.CreatePagePath(p.SubCateReport),
		Cate1Eng:      pagepath.CreatePagePath(p.Cate1),
		Cate2Eng:      cate2,
		ProductLists:  ProductList,
		CateReport:    p.CateReport,
		SubCateReport: p.SubCateReport,
		Cate1:         p.Cate1,
		Cate2:         p.Cate2,
	}

	render.RenderTemplate(w, "product_list_cate2", page)
}
