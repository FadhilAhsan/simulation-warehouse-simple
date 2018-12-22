package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/FadhilAhsan/simulation-warehouse-simple/controllers/api:ItemController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/simulation-warehouse-simple/controllers/api:ItemController"],
		beego.ControllerComments{
			Method: "ViewListItem",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/simulation-warehouse-simple/controllers/api:ItemController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/simulation-warehouse-simple/controllers/api:ItemController"],
		beego.ControllerComments{
			Method: "ViewRequestListItem",
			Router: `/request`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/simulation-warehouse-simple/controllers/api:ItemController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/simulation-warehouse-simple/controllers/api:ItemController"],
		beego.ControllerComments{
			Method: "ViewLogItemOut",
			Router: `/log-item-out`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/simulation-warehouse-simple/controllers/api:ItemController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/simulation-warehouse-simple/controllers/api:ItemController"],
		beego.ControllerComments{
			Method: "ExportReportValueItem",
			Router: `/report`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/simulation-warehouse-simple/controllers/api:OrderController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/simulation-warehouse-simple/controllers/api:OrderController"],
		beego.ControllerComments{
			Method: "ViewListOrderItem",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/FadhilAhsan/simulation-warehouse-simple/controllers/api:OrderController"] = append(beego.GlobalControllerRouter["github.com/FadhilAhsan/simulation-warehouse-simple/controllers/api:OrderController"],
		beego.ControllerComments{
			Method: "ExportReportOrderItem",
			Router: `/report`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
