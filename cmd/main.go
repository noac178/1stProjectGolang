package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/save_create", saveCreateHandler)
	http.HandleFunc("/product_list", listHandler)
	http.HandleFunc("/update/", updateHandler)
	http.HandleFunc("/save_update/", saveUpdateHandler)
	http.HandleFunc("/delete/", deleteHandler)
	http.HandleFunc("/product_list/cate", filterCateHandler)
	http.HandleFunc("/product_list/c/", cateHandler)
	http.HandleFunc("/product_list/subcate", filterSubCateHandler)
	http.HandleFunc("/product_list/sc/", subCateHandler)
	http.HandleFunc("/product_list/cate1", filterCate1Handler)
	http.HandleFunc("/product_list/c1/", cate1Handler)
	http.HandleFunc("/product_list/cate2", filterCate2Handler)
	http.HandleFunc("/product_list/c2/", cate2Handler)
	http.HandleFunc("/p/", pdpHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
