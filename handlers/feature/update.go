package feature

import "net/http"

func updateHandler(w http.ResponseWriter, r *http.Request) {
	product_id := r.URL.Path[len("/update/"):]

	db, err := openDb()

	var p ProductInfo

	query := `SELECT id, sku, name, price, number, cate_report, sub_cate_report, 
			cate1, cate2, color, size, brand, image FROM product_info WHERE id = ?`
	err = db.QueryRow(query, product_id).Scan(&p.Id, &p.Sku, &p.Name, &p.Price, &p.Number, &p.CateReport, &p.SubCateReport,
		&p.Cate1, &p.Cate2, &p.Color, &p.Size, &p.Brand, &p.Image)
	checkErr(err)

	renderTemplate(w, "update", p)
}
