package http

import "net/http"

func filterCate2Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cate2 := r.Form.Get("cate2")

	http.Redirect(w, r, "/product_list/c2/"+createPagePath(cate2), http.StatusFound)
}

func cate2Handler(w http.ResponseWriter, r *http.Request) {
	cate2 := r.URL.Path[len("/product_list/c2/"):]

	db, err := openDb()
	checkErr(err)

	rows, err := db.Query(`
		SELECT DISTINCT id, name, price, cate_report, sub_cate_report,
				cate1, cate2, image
		FROM product_info 
		WHERE 1=1
		AND (
		CASE
			WHEN cate2 NOT LIKE '%-%' THEN LOWER(REPLACE(REPLACE(cate2, ' - ', '-'), ' ', '-')) = ?
			ELSE LOWER(REPLACE(REPLACE(cate2, ' - ', '-'), ' ', '-')) = ?
		END)`, cate2, cate2)
	checkErr(err)
	defer rows.Close()

	var ProductList []ProductInfo

	for rows.Next() {
		var p ProductInfo

		err := rows.Scan(&p.Id, &p.Name, &p.Price, &p.CateReport, &p.SubCateReport,
			&p.Cate1, &p.Cate2, &p.Image)
		checkErr(err)

		ProductList = append(ProductList, p)
	}
	err = rows.Err()
	checkErr(err)

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

	var p ProductInfo
	err = db.QueryRow(query, cate2, cate2).Scan(&p.CateReport, &p.SubCateReport, &p.Cate1, &p.Cate2)
	checkErr(err)

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
		CateEng:       createPagePath(p.CateReport),
		SubCateEng:    createPagePath(p.SubCateReport),
		Cate1Eng:      createPagePath(p.Cate1),
		Cate2Eng:      cate2,
		ProductLists:  ProductList,
		CateReport:    p.CateReport,
		SubCateReport: p.SubCateReport,
		Cate1:         p.Cate1,
		Cate2:         p.Cate2,
	}

	if err := templates.ExecuteTemplate(w, "product_list_cate2.html", page); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
