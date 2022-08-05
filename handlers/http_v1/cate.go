package http

import "net/http"

func FilterCateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cate_report := r.Form.Get("cate_report")

	http.Redirect(w, r, "/product_list/c/"+CreatePagePath(cate_report), http.StatusFound)
}

func CateHandler(w http.ResponseWriter, r *http.Request) {
	cate_report := r.URL.Path[len("/product_list/c/"):]

	db, err := OpenDb()
	rows, err := db.Query(`
		SELECT DISTINCT id, name, price, cate_report, sub_cate_report,
				cate1, cate2, image
		FROM product_info 
		WHERE 1=1
		AND (
		CASE
			WHEN cate_report NOT LIKE '%-%' THEN LOWER(REPLACE(REPLACE(cate_report, ' - ', '-'), ' ', '-')) = ?
			ELSE LOWER(REPLACE(REPLACE(cate_report, ' - ', '-'), ' ', '-')) = ?
		END)`, cate_report, cate_report)
	CheckErr(err)
	defer rows.Close()

	var ProductList []ProductInfo

	for rows.Next() {
		var p ProductInfo

		err := rows.Scan(&p.Id, &p.Name, &p.Price, &p.CateReport, &p.SubCateReport,
			&p.Cate1, &p.Cate2, &p.Image)
		CheckErr(err)

		ProductList = append(ProductList, p)
	}
	err = rows.Err()
	CheckErr(err)

	query := `
		SELECT
			DISTINCT
			cate_report
		FROM product_info 
		WHERE 1=1
		AND (
		CASE
			WHEN cate_report NOT LIKE '%-%' THEN LOWER(REPLACE(REPLACE(cate_report, ' - ', '-'), ' ', '-')) = ?
			ELSE LOWER(REPLACE(REPLACE(cate_report, ' - ', '-'), ' ', '-')) = ?
		END)`

	var p ProductInfo
	err = db.QueryRow(query, cate_report, cate_report).Scan(&p.CateReport)
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
		CateEng:      cate_report,
		ProductLists: ProductList,
		CateReport:   p.CateReport,
	}

	RenderTemplate(w, "product_list_cate", page)
}
