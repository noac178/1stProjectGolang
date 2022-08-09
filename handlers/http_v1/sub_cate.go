package http_v1

import (
	"net/http"

	"github.com/noac178/1stProjectGolang/entity"
	"github.com/noac178/1stProjectGolang/pkg/errorx"
	"github.com/noac178/1stProjectGolang/pkg/pagepath"
	"github.com/noac178/1stProjectGolang/pkg/render"
	"github.com/noac178/1stProjectGolang/repo"
)

func FilterSubCateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sub_cate_report := r.Form.Get("sub_cate_report")

	http.Redirect(w, r, "/product_list/sc/"+pagepath.CreatePagePath(sub_cate_report), http.StatusFound)
}

func SubCateHandler(w http.ResponseWriter, r *http.Request) {
	sub_cate_report := r.URL.Path[len("/product_list/sc/"):]

	db, err := repo.OpenDb()
	errorx.CheckErr(err)

	rows, err := db.Query(`
		SELECT DISTINCT id, name, price, cate_report, sub_cate_report,
				cate1, cate2, image
		FROM product_info
		WHERE 1=1
		AND (
		CASE
			WHEN sub_cate_report NOT LIKE '%-%' THEN LOWER(REPLACE(REPLACE(sub_cate_report, ' - ', '-'), ' ', '-')) = ?
			ELSE LOWER(REPLACE(REPLACE(sub_cate_report, ' - ', '-'), ' ', '-')) = ?
		END)`, sub_cate_report, sub_cate_report)
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
			sub_cate_report
		FROM product_info
		WHERE 1=1
		AND (
		CASE
			WHEN sub_cate_report NOT LIKE '%-%' THEN LOWER(REPLACE(REPLACE(sub_cate_report, ' - ', '-'), ' ', '-')) = ?
			ELSE LOWER(REPLACE(REPLACE(sub_cate_report, ' - ', '-'), ' ', '-')) = ?
		END)`

	var p entity.ProductInfo
	err = db.QueryRow(query, sub_cate_report, sub_cate_report).Scan(&p.CateReport, &p.SubCateReport)
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
		SubCateEng:    sub_cate_report,
		ProductLists:  ProductList,
		CateReport:    p.CateReport,
		SubCateReport: p.SubCateReport,
	}

	render.RenderTemplate(w, "product_list_subcate", page)
}
