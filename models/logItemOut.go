package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type LogItemOut struct {
	Id           int        `json:"-"`
	CreatedAt    *time.Time `json:"-" orm:"type(datetime)"`
	TypeLog      string     `json:"-"`
	IdOrder      string     `json:"-"`
	Sku          string     `json:"sku"`
	NameItem     string     `json:"name_item"`
	Qty          int        `json:"qty"`
	Price        int64      `json:"price"`
	TotalPrice   int64      `json:"total_price"`
	CreatedAtStr string     `json:"created_at" orm:"-"`
	Note         string     `json:"note" orm:"-"`
}

func init() {
	orm.RegisterModel(new(LogItemOut))
}

func (this *LogItemOut) TableName() string {
	return "log_item_out"
}
