package feature

import "net/http"

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	product_id := r.URL.Path[len("/delete/"):]

	db, err := openDb()
	_, err1 := db.Exec(`DELETE FROM product_info WHERE id = ?`, product_id)
	checkErr(err1)

	http.Redirect(w, r, "/product_list", http.StatusFound)
}
