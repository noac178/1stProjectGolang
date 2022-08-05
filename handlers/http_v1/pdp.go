package http

import "net/http"

func PdpHandler(w http.ResponseWriter, r *http.Request) {
	product_id := r.URL.Path[len("/p/"):]

	db, err := OpenDb()
	CheckErr(err)

	var p ProductInfo

	query := `SELECT id, sku, name, price, number, cate_report, sub_cate_report, 
			cate1, cate2, color, size, brand, image FROM product_info WHERE id = ?`
	err = db.QueryRow(query, product_id).Scan(&p.Id, &p.Sku, &p.Name, &p.Price, &p.Number, &p.CateReport, &p.SubCateReport,
		&p.Cate1, &p.Cate2, &p.Color, &p.Size, &p.Brand, &p.Image)
	CheckErr(err)

	type Page struct {
		CateEng       string
		CateReport    string
		SubCateEng    string
		SubCateReport string
		Cate1Eng      string
		Cate1         string
		Cate2Eng      string
		Cate2         string
		ProductInfos  ProductInfo
	}

	page := &Page{
		CateEng:      CreatePagePath(p.CateReport),
		SubCateEng:   CreatePagePath(p.SubCateReport),
		Cate1Eng:     CreatePagePath(p.Cate1),
		Cate2Eng:     CreatePagePath(p.Cate2),
		ProductInfos: p,
	}

	RenderTemplate(w, "pdp", page)
}
