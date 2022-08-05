package http

import "net/http"

func FilterSubCateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sub_cate_report := r.Form.Get("sub_cate_report")

	http.Redirect(w, r, "/product_list/sc/"+CreatePagePath(sub_cate_report), http.StatusFound)
}

func SubCateHandler(w http.ResponseWriter, r *http.Request) {
	sub_cate_report := r.URL.Path[len("/product_list/sc/"):]

	db, err := OpenDb()
	CheckErr(err)

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
	CheckErr(err)
	defer rows.Close()

	var ProductList []ProductInfo

	for rows.Next() {
		var p ProductInfo

		err := rows.Scan(&p.Id, &p.Name, &p.Price, &p.CateReport, &p.SubCateReport,
			&p.Cate1, &p.Cate2, &p.Image)
		errorx.CheckErr(err)

		ProductList = append(ProductList, p)
	}
	err = rows.Err()
	CheckErr(err)

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

	var p ProductInfo
	err = db.QueryRow(query, sub_cate_report, sub_cate_report).Scan(&p.CateReport, &p.SubCateReport)
	CheckErr(err)

	type Page struct {
		CateEng       string
		CateReport    string
		SubCateEng    string
		SubCateReport string
		Cate1Eng      string
		Cate1         string
		Cate2Eng      string
		Cate2         string
		ProductLists  []ProductInfo
	}

	page := &Page{
		CateEng:       CreatePagePath(p.CateReport),
		SubCateEng:    sub_cate_report,
		ProductLists:  ProductList,
		CateReport:    p.CateReport,
		SubCateReport: p.SubCateReport,
	}

	RenderTemplate(w, "product_list_subcate", page)
}
