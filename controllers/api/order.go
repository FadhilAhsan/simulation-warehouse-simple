package api

import (
	"github.com/FadhilAhsan/simulation-warehouse-simple/models"
	"github.com/FadhilAhsan/simulation-warehouse-simple/services"
	"github.com/FadhilAhsan/simulation-warehouse-simple/utilities"
	"github.com/astaxie/beego"
	"io/ioutil"
	"path"
	"strings"
)

type OrderController struct {
	beego.Controller
}

// @Title ViewListOrderItem
// @Description Get List All Order Item
// @Param limit query int false "limit list item"
// @Param offset query int false "offset list item"
// @Success 200 {object} models.json.JSONResponse
// @Failed 500 something wrong in server
// @router / [get]
func (this *OrderController) ViewListOrderItem() {
	meta := models.JSONResponseMeta{}

	limit, errLimit := this.GetInt("limit")
	if errLimit != nil {
		limit = 0
	}

	offset, _ := this.GetInt("offset")

	data, err := services.GetListOrderItem(limit, offset)
	if err != nil {
		meta = models.JSONResponseMeta{
			Code:    500,
			Status:  "Failed",
			Message: "Failed to get list order item. err : " + err.Error(),
		}
	} else {
		meta = models.JSONResponseMeta{
			Code:    200,
			Status:  "Success",
			Message: "Success to get list order item.",
		}
	}

	result := models.JSONResponse{
		Meta: meta,
		Data: data,
	}

	this.Data["json"] = result
	this.ServeJSON()
}

// @Title ExportReportOrderItem
// @Description for export report sales
// @Param startDate query string false "start date for report sales"
// @Param endDate query string false "end date for report sales"
// @router /report [get]
func (this *OrderController) ExportReportOrderItem() {
	startDate := this.GetString("startDate")
	if startDate == "" {
		startDate = "2018-10-01"
	}

	endDate := this.GetString("endDate")
	if endDate == "" {
		endDate = "2018-12-30"
	}
	data, _ := services.GetReportOrder(startDate, endDate)

	filename := utilities.ExportReportOrderItemToCsv(*data, startDate, endDate)

	basepath := path.Base("/csv-data")
	filepath := path.Base(filename)

	fileBytes, _ := ioutil.ReadFile(strings.Join([]string{basepath, filepath}, "/"))

	this.Ctx.ResponseWriter.Header().Set("Content-Type", "text/csv")
	this.Ctx.ResponseWriter.Header().Set("Content-Disposition", "attachment;filename="+filename)
	this.Ctx.ResponseWriter.Write(fileBytes)

}
