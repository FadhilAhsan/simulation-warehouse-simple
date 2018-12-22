package main

import (
	_ "github.com/FadhilAhsan/simulation-warehouse-simple/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "file:database/warehouse.db")
}

func main() {

	beego.Run()
}
