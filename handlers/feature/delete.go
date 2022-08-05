package feature

import "net/http"

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	product_id := r.URL.Path[len("/delete/"):]

	db, err := OpenDb()
	_, err1 := db.Exec(`DELETE FROM product_info WHERE id = ?`, product_id)
	CheckErr(err1)

	http.Redirect(w, r, "/product_list", http.StatusFound)
}
