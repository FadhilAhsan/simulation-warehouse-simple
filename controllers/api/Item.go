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

type ItemController struct {
	beego.Controller
}

// @Title ViewListItem
// @Description Get List Item
// @Param limit query int false "limit list item"
// @Param offset query int false "offset list item"
// @Success 200 {object} models.json.JSONResponse
// @Failed 500 something wrong in server
// @router / [get]
func (this *ItemController) ViewListItem() {
	meta := models.JSONResponseMeta{}

	limit, errLimit := this.GetInt("limit")
	if errLimit != nil {
		limit = 0
	}

	offset, _ := this.GetInt("offset")

	data, err := services.GetListItem(limit, offset)
	if err != nil {
		meta = models.JSONResponseMeta{
			Code:    500,
			Status:  "Failed",
			Message: "Failed to get list item. err : " + err.Error(),
		}
	} else {
		meta = models.JSONResponseMeta{
			Code:    200,
			Status:  "Success",
			Message: "Success to get list item",
		}
	}

	result := models.JSONResponse{
		Meta: meta,
		Data: data,
	}

	this.Data["json"] = result
	this.ServeJSON()
}

// @Title ViewRequestListItem
// @Description Get List request Item
// @Param limit query int false "limit list item"
// @Param offset query int false "offset list item"
// @Success 200 {object} models.json.JSONResponse
// @Failed 500 something wrong in server
// @router /request [get]
func (this *ItemController) ViewRequestListItem() {
	meta := models.JSONResponseMeta{}

	limit, errLimit := this.GetInt("limit")
	if errLimit != nil {
		limit = 0
	}

	offset, _ := this.GetInt("offset")

	data, err := services.GetRequestListItem(limit, offset)
	if err != nil {
		meta = models.JSONResponseMeta{
			Code:    500,
			Status:  "Failed",
			Message: "Failed to get Request list item. err : " + err.Error(),
		}
	} else {
		meta = models.JSONResponseMeta{
			Code:    200,
			Status:  "Success",
			Message: "Success to get Request list item",
		}
	}

	result := models.JSONResponse{
		Meta: meta,
		Data: data,
	}

	this.Data["json"] = result
	this.ServeJSON()
}

// @Title ViewLogItemOut
// @Description Get Log Item out on warehouse
// @Param limit query int false "limit list item"
// @Param offset query int false "offset list item"
// @Success 200 {object} models.json.JSONResponse
// @Failed 500 something wrong in server
// @router /log-item-out [get]
func (this *ItemController) ViewLogItemOut() {
	meta := models.JSONResponseMeta{}

	limit, errLimit := this.GetInt("limit")
	if errLimit != nil {
		limit = 0
	}

	offset, _ := this.GetInt("offset")

	data, err := services.GetListLogItemOut(limit, offset)
	if err != nil {
		meta = models.JSONResponseMeta{
			Code:    500,
			Status:  "Failed",
			Message: "Failed to get Request list item. err : " + err.Error(),
		}
	} else {
		meta = models.JSONResponseMeta{
			Code:    200,
			Status:  "Success",
			Message: "Success to get Request list item",
		}
	}

	result := models.JSONResponse{
		Meta: meta,
		Data: data,
	}

	this.Data["json"] = result
	this.ServeJSON()
}

// @Title ExportReportValueItem
// @Description for export report value item
// @router /report [get]
func (this *ItemController) ExportReportValueItem() {

	data, _ := services.GetReportValueItem()

	filename := utilities.ExportReportValueItemToCsv(*data)

	basepath := path.Base("/csv-data")
	filepath := path.Base(filename)

	fileBytes, _ := ioutil.ReadFile(strings.Join([]string{basepath, filepath}, "/"))

	this.Ctx.ResponseWriter.Header().Set("Content-Type", "text/csv")
	this.Ctx.ResponseWriter.Header().Set("Content-Disposition", "attachment;filename="+filename)
	this.Ctx.ResponseWriter.Write(fileBytes)

}
