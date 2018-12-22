package models

import (
	"github.com/astaxie/beego/orm"
)

type Color struct {
	Id        int    `json:"-"`
	Name      string `json:"color"`
	ColorCode string `json:"color_code"`
}

func init() {
	orm.RegisterModel(new(Color))
}

func (this *Color) TableName() string {
	return "color"
}
