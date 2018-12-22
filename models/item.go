package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Item struct {
	Id        int        `json:"id"`
	Sku       string     `json:"sku"`
	Name      string     `json:"name"`
	Stock     int        `json:"stock"`
	Size      *Size      `json:"-" orm:"rel(fk)"`
	Color     *Color     `json:"-" orm:"rel(fk)"`
	Price     int64      `json:"-"`
	IsActive  int        `json:"-"`
	CreatedAt *time.Time `json:"-" orm:"type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Item))
}

func (this *Item) TableName() string {
	return "items"
}
