package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type OrderItem struct {
	Id           int        `json:"-"`
	CreatedAt    *time.Time `json:"-" orm:"type(datetime)"`
	IdOrder      string     `json:"id_order"`
	Sku          string     `json:"sku"`
	NameItem     string     `json:"name_item"`
	Qty          int        `json:"qty"`
	Price        int64      `json:"price"`
	TotalPrice   int64      `json:"total_price"`
	CreatedAtStr string     `json:"created_at" orm:"-"`
}

func init() {
	orm.RegisterModel(new(OrderItem))
}

func (this *OrderItem) TableName() string {
	return "order_item"
}
