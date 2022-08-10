package entity

import (
	"time"
)

type ProductInfo struct {
	ID        uint `gorm:"primaryKey"`
	Sku       string
	Name      string
	Price     int64
	Number    int64
	Cate1     string
	Cate2     string
	Cate3     string
	Cate4     string
	Color     string
	Size      string
	Brand     string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
