package main

import (
	"log"
	"net/http"

	"github.com/noac178/1stProjectGolang/handlers/feature"
	"github.com/noac178/1stProjectGolang/handlers/http_v1"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/create", feature.CreateHandler)
	http.HandleFunc("/save_create", feature.SaveCreateHandler)
	http.HandleFunc("/product_list", ListHandler)
	http.HandleFunc("/update/", UpdateHandler)
	http.HandleFunc("/save_update/", SaveUpdateHandler)
	http.HandleFunc("/delete/", feature.DeleteHandler)
	http.HandleFunc("/product_list/cate", FilterCateHandler)
	http.HandleFunc("/product_list/c/", CateHandler)
	http.HandleFunc("/product_list/subcate", FilterSubCateHandler)
	http.HandleFunc("/product_list/sc/", SubCateHandler)
	http.HandleFunc("/product_list/cate1", FilterCate1Handler)
	http.HandleFunc("/product_list/c1/", http_v1.Cate1Handler)
	http.HandleFunc("/product_list/cate2", FilterCate2Handler)
	http.HandleFunc("/product_list/c2/", Cate2Handler)
	http.HandleFunc("/p/", PdpHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
