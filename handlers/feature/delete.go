package feature

import (
	"Minipj/1stProjectGolang/pkg/errorx"
	"net/http"

	"github.com/noac178/1stProjectGolang/pkg/errorx"
	"github.com/noac178/1stProjectGolang/repo"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	product_id := r.URL.Path[len("/delete/"):]

	db, err := repo.OpenDb()
	errorx.CheckErr(err)
	_, err1 := db.Exec(`DELETE FROM product_info WHERE id = ?`, product_id)
	errorx.CheckErr(err1)

	http.Redirect(w, r, "/product_list", http.StatusFound)
}
