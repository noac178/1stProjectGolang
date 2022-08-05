package http

import "net/http"

func FilterCate1Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cate1 := r.Form.Get("cate1")

	http.Redirect(w, r, "/product_list/c1/"+CreatePagePath(cate1), http.StatusFound)
}

func Cate1Handler(w http.ResponseWriter, r *http.Request) {
	cate1 := r.URL.Path[len("/product_list/c1/"):]

	db, err := OpenDb()
	CheckErr(err)

	rows, err := db.Query(`
		SELECT DISTINCT id, name, price, cate_report, sub_cate_report,
				cate1, cate2, image
		FROM product_info 
		WHERE 1=1
		AND (
		CASE
			WHEN cate1 NOT LIKE '%-%' THEN LOWER(REPLACE(REPLACE(cate1, ' - ', '-'), ' ', '-')) = ?
			ELSE LOWER(REPLACE(REPLACE(cate1, ' - ', '-'), ' ', '-')) = ?
		END)`, cate1, cate1)
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
			cate_report,
			sub_cate_report,
			cate1
		FROM product_info 
		WHERE 1=1
		AND (
		CASE
			WHEN cate1 NOT LIKE '%-%' THEN LOWER(REPLACE(REPLACE(cate1, ' - ', '-'), ' ', '-')) = ?
			ELSE LOWER(REPLACE(REPLACE(cate1, ' - ', '-'), ' ', '-')) = ?
		END)`

	var p ProductInfo
	err = db.QueryRow(query, cate1, cate1).Scan(&p.CateReport, &p.SubCateReport, &p.Cate1)
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
		SubCateEng:    CreatePagePath(p.SubCateReport),
		Cate1Eng:      cate1,
		ProductLists:  ProductList,
		CateReport:    p.CateReport,
		SubCateReport: p.SubCateReport,
		Cate1:         p.Cate1,
	}

	RenderTemplate(w, "product_list_cate1", page)
}
