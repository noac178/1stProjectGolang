package save

import "net/http"

func saveUpdateHandler(w http.ResponseWriter, r *http.Request) {
	product_id := r.URL.Path[len("/save_update/"):]

	sku := r.FormValue("sku")
	name := r.FormValue("name")
	price := r.FormValue("price")
	number := r.FormValue("number")
	cate_report := r.FormValue("cate_report")
	sub_cate_report := r.FormValue("sub_cate_report")
	cate1 := r.FormValue("cate1")
	cate2 := r.FormValue("cate2")
	color := r.FormValue("color")
	size := r.FormValue("size")
	brand := r.FormValue("brand")
	image := r.FormValue("image")

	db, err := openDb()
	updateInfo, err := db.Prepare(`UPDATE product_info 
								SET sku = ?, 
									name = ?,
									price = ?, 
									number = ?,
									cate_report = ?, 
									sub_cate_report = ?,
									cate1 = ?, 
									cate2 = ?, 
									color = ?,
									size = ?,
									brand = ?,
									image = ?
								WHERE id = ?`)
	checkErr(err)

	_, err1 := updateInfo.Exec(sku, name, price, number, cate_report, sub_cate_report,
		cate1, cate2, color, size, brand, image, product_id)
	checkErr(err1)
	updateInfo.Close()

	http.Redirect(w, r, "/product_list", http.StatusFound)
}
