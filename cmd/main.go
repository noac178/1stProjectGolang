package main

import (
	"log"
	"net/http"

	"github.com/noac178/1stProjectGolang/handlers/feature"
	"github.com/noac178/1stProjectGolang/handlers/http_v1"
	"github.com/noac178/1stProjectGolang/handlers/save"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/create", feature.CreateHandler)
	http.HandleFunc("/save_create", save.SaveCreateHandler)
	http.HandleFunc("/product_list", http_v1.ListHandler)
	http.HandleFunc("/update/", feature.UpdateHandler)
	http.HandleFunc("/save_update/", save.SaveUpdateHandler)
	http.HandleFunc("/delete/", feature.DeleteHandler)
	http.HandleFunc("/product_list/cate", http_v1.FilterCateHandler)
	http.HandleFunc("/product_list/c/", http_v1.CateHandler)
	http.HandleFunc("/product_list/subcate", http_v1.FilterSubCateHandler)
	http.HandleFunc("/product_list/sc/", http_v1.SubCateHandler)
	http.HandleFunc("/product_list/cate1", http_v1.FilterCate1Handler)
	http.HandleFunc("/product_list/c1/", http_v1.Cate1Handler)
	http.HandleFunc("/product_list/cate2", http_v1.FilterCate2Handler)
	http.HandleFunc("/product_list/c2/", http_v1.Cate2Handler)
	http.HandleFunc("/p/", http_v1.PdpHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
