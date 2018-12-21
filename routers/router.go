package routers

import (
	"github.com/FadhilAhsan/simulation-warehouse-simple/simulation-warehouse-simple/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
