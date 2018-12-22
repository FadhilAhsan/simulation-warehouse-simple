package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type RequestItem struct {
	Id              int               `json:"id"`
	CreatedAtStr    string            `json:"created_at" orm:"-"`
	UpdatedAtStr    string            `json:"updated_at" orm:"-"`
	SKUItem         string            `json:"sku_item" orm:"-"`
	ItemName        string            `json:"item_name" orm:"-"`
	RequestNumber   int               `json:"request_number"`
	RequestRecieved int               `json:"request_recieved"`
	PriceItem       int64             `json:"price_item"`
	TotalPrice      int64             `json:"total_price"`
	Invoice         string            `json:"invoice"`
	Note            string            `json:"note" orm:"-"`
	IsComplete      bool              `json:"is_complete"`
	CreatedAt       *time.Time        `json:"-"`
	UpdatedAt       *time.Time        `json:"-"`
	Item            *Item             `json:"-" orm:"rel(fk)"`
	LogRequestItem  []*LogRequestItem `json:"-" orm:"reverse(many)"`
}

type LogRequestItem struct {
	Id           int          `json:"-"`
	RequestItem  *RequestItem `json:"-" orm:"rel(fk)"`
	ItemReceived int          `json:"item_received"`
	CreatedAt    *time.Time   `json:"created_at"`
}

func init() {
	orm.RegisterModel(new(RequestItem))
	orm.RegisterModel(new(LogRequestItem))
}

func (this *RequestItem) TableName() string {
	return "request_item"
}

func (this *LogRequestItem) TableName() string {
	return "log_request_item"
}
