package services

import (
	"github.com/FadhilAhsan/simulation-warehouse-simple/models"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
)

func GetListOrderItem(limit int, offset int) (*[]models.OrderItem, error) {
	dbORM := orm.NewOrm()
	dbORM.Using("default")
	var err error
	timeLayout := "2006-01-02 15:04:05"

	orderItems := []models.OrderItem{}
	_, err = dbORM.QueryTable("order_item").Limit(limit).Offset(offset).RelatedSel().All(&orderItems)
	if err != nil {
		return nil, err

	}

	for i, _ := range orderItems {
		orderItems[i].CreatedAtStr = orderItems[i].CreatedAt.Format(timeLayout)
	}

	return &orderItems, nil

}

func GetReportOrder(startDate string, endDate string) (*models.ReportOrder, error) {
	queryBuilder, _ := orm.NewQueryBuilder("mysql")
	queryBuilder.Select("order_item.id_order, order_item.created_at, items.sku, items.name, size.size as size, color.name as color, order_item.qty, items.price, request_item.price_item, order_item.qty * items.price as total, request_item.price_item as purchase_price, (order_item.qty * items.price) -(request_item.price_item *order_item.qty) as profit").From("order_item").InnerJoin("items").On("order_item.item_id = items.id").InnerJoin("size").On("items.size_id = size.id").InnerJoin("color").On("items.color_id = color.id").InnerJoin("request_item").On("items.id = request_item.item_id").Where("order_item.created_at > '" + startDate + "' and order_item.created_at < '" + endDate + "'").OrderBy("order_item.created_at").Asc()

	querySql := queryBuilder.String()
	log.Println(querySql)

	dbORM := orm.NewOrm()
	dbORM.Using("default")
	var err error

	reportOrderItem := []models.ReportOrderItem{}
	totalRows, err := dbORM.Raw(querySql).QueryRows(&reportOrderItem)
	if err != nil {
		log.Println(err)
		return nil, err

	}

	headerReportOrderItem := []models.HeaderReportOrderItem{}
	querySql = "select sum(qty) as total_item, sum(total) as total_income, sum(profit) as total_profit from (" + querySql + ")"
	_, err = dbORM.Raw(querySql).QueryRows(&headerReportOrderItem)
	if err != nil {
		return nil, err

	}
	headerReportOrderItem[0].TotalSeles = strconv.FormatInt(totalRows, 10)

	reportOrder := models.ReportOrder{
		Header: headerReportOrderItem[0],
		Items:  reportOrderItem,
	}

	log.Println(reportOrderItem)
	return &reportOrder, nil

}
