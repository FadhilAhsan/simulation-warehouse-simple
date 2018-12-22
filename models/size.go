package models

import (
	"github.com/astaxie/beego/orm"
)

type Size struct {
	Id       int    `json:"-"`
	Size     string `json:"size"`
	SizeCode string `json:"size_code"`
}

func init() {
	orm.RegisterModel(new(Size))
}

func (this *Size) TableName() string {
	return "size"
}
