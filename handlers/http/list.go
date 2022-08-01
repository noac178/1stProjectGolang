package http

import "net/http"

func listHandler(w http.ResponseWriter, r *http.Request) {
	db, _ := openDb()

	rows, err := db.Query(`SELECT id, name, price, cate_report, sub_cate_report,
						cate1, cate2, image FROM product_info`)
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

	renderTemplate(w, "product_list", ProductList)
}
