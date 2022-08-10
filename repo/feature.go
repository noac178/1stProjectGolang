package repo

import (
	"context"
	"net/http"

	"github.com/noac178/1stProjectGolang/entity"
	"github.com/noac178/1stProjectGolang/infra"
)

type FeatureMysqlRepo struct {
	db *infra.MysqlConnPool
}

func (f *FeatureMysqlRepo) Create(ctx context.Context) error {
	productInfo := new(entity.ProductInfo)
	f.db.Conn.Create(&productInfo)

	return c.JSONPretty(http.StatusOK, subject, "  ")
}

func (f *FeatureMysqlRepo) Delete(ctx context.Context, id string) error {
	return f.db.Conn.Delete(&entity.ProductInfo{}, id).Error
}
