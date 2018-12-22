package routers

import (
	"github.com/FadhilAhsan/simulation-warehouse-simple/controllers"
	"github.com/FadhilAhsan/simulation-warehouse-simple/controllers/api"
	"github.com/astaxie/beego"
)

func init() {
	namespace := beego.NewNamespace("/api",

		beego.NSNamespace("/item",
			beego.NSInclude(&api.ItemController{}),
		),

		beego.NSNamespace("/order",
			beego.NSInclude(&api.OrderController{}),
		),
	)
	beego.Router("/", &controllers.MainController{})
	beego.AddNamespace(namespace)
}
