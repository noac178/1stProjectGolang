package http_v1

import (
	"net/http"

	"github.com/noac178/1stProjectGolang/entity"
	"github.com/noac178/1stProjectGolang/pkg/errorx"
	"github.com/noac178/1stProjectGolang/pkg/pagepath"
	"github.com/noac178/1stProjectGolang/pkg/render"
	"github.com/noac178/1stProjectGolang/repo"
)

func PdpHandler(w http.ResponseWriter, r *http.Request) {
	product_id := r.URL.Path[len("/p/"):]

	db, err := repo.OpenDb()
	errorx.CheckErr(err)

	var p entity.ProductInfo

	query := `SELECT id, sku, name, price, number, cate_report, sub_cate_report, 
			cate1, cate2, color, size, brand, image FROM product_info WHERE id = ?`
	err = db.QueryRow(query, product_id).Scan(&p.Id, &p.Sku, &p.Name, &p.Price, &p.Number, &p.CateReport, &p.SubCateReport,
		&p.Cate1, &p.Cate2, &p.Color, &p.Size, &p.Brand, &p.Image)
	errorx.CheckErr(err)

	type Page struct {
		CateEng       string
		CateReport    string
		SubCateEng    string
		SubCateReport string
		Cate1Eng      string
		Cate1         string
		Cate2Eng      string
		Cate2         string
		ProductInfos  entity.ProductInfo
	}

	page := &Page{
		CateEng:      pagepath.CreatePagePath(p.CateReport),
		SubCateEng:   pagepath.CreatePagePath(p.SubCateReport),
		Cate1Eng:     pagepath.CreatePagePath(p.Cate1),
		Cate2Eng:     pagepath.CreatePagePath(p.Cate2),
		ProductInfos: p,
	}

	render.RenderTemplate(w, "pdp", page)
}
