package services

import (
	"github.com/FadhilAhsan/simulation-warehouse-simple/models"
	"github.com/astaxie/beego/orm"
	"strconv"
)

func GetListItem(limit int, offset int) (*[]models.Item, error) {
	dbORM := orm.NewOrm()
	dbORM.Using("default")
	var err error

	items := []models.Item{}
	_, err = dbORM.QueryTable("items").Limit(limit).Offset(offset).RelatedSel().All(&items)
	if err != nil {
		return nil, err

	}

	for i, _ := range items {
		items[i].Name = items[i].Name + "(" + items[i].Size.Size + "," + items[i].Color.Name + ")"
	}

	return &items, nil

}

func GetRequestListItem(limit int, offset int) (*[]models.RequestItem, error) {
	dbORM := orm.NewOrm()
	dbORM.Using("default")
	var err error
	timeLayoutLog := "2006-01-02"
	timeLayoutRequestItem := "2006-01-02 15:04:05"

	requestItems := []models.RequestItem{}
	_, err = dbORM.QueryTable("request_item").Limit(limit).Offset(offset).RelatedSel().All(&requestItems)
	if err != nil {
		return nil, err

	}

	for i, requestItem := range requestItems {
		err = dbORM.Read(&requestItem)
		dbORM.LoadRelated(&requestItem, "LogRequestItem")

		var note = ""
		for _, log := range requestItem.LogRequestItem {
			note = note + log.CreatedAt.Format(timeLayoutLog) + " terima " + strconv.Itoa(log.ItemReceived) + "; "
		}

		if requestItems[i].IsComplete == false {
			note = note + "Masih menunggu"
		}

		requestItems[i].Note = note
		requestItems[i].SKUItem = requestItems[i].Item.Sku
		requestItems[i].ItemName = requestItems[i].Item.Name + "(" + requestItems[i].Item.Size.Size + "," + requestItems[i].Item.Color.Name + ")"
		requestItems[i].CreatedAtStr = requestItems[i].CreatedAt.Format(timeLayoutRequestItem)
		requestItems[i].UpdatedAtStr = requestItems[i].UpdatedAt.Format(timeLayoutRequestItem)

		if requestItems[i].Invoice == "" {
			requestItems[i].Invoice = "(Hilang)"
		}
	}

	return &requestItems, nil
}

func GetListLogItemOut(limit int, offset int) (*[]models.LogItemOut, error) {
	dbORM := orm.NewOrm()
	dbORM.Using("default")
	var err error
	timeLayout := "2006-01-02 15:04:05"

	logItemsOut := []models.LogItemOut{}
	_, err = dbORM.QueryTable("log_item_out").Limit(limit).Offset(offset).RelatedSel().All(&logItemsOut)
	if err != nil {
		return nil, err

	}

	for i, _ := range logItemsOut {
		if logItemsOut[i].TypeLog == "lost" {
			logItemsOut[i].Note = "Barang Hilang"
		} else if logItemsOut[i].TypeLog == "damaged" {
			logItemsOut[i].Note = "Barang Rusak"
		} else {
			logItemsOut[i].Note = "Pesanan " + logItemsOut[i].IdOrder
		}

		logItemsOut[i].CreatedAtStr = logItemsOut[i].CreatedAt.Format(timeLayout)
	}

	return &logItemsOut, nil

}

func GetReportValueItem() (*models.ReportValueItem, error) {
	queryBuilder, _ := orm.NewQueryBuilder("mysql")
	queryBuilder.Select("items.sku, items.name, size.size as size, color.name as color, items.stock, request_item.price_item, items.stock * request_item.price_item as total").From("items").InnerJoin("size").On("items.size_id = size.id").InnerJoin("color").On("items.color_id = color.id").InnerJoin("request_item").On("items.id = request_item.item_id").GroupBy("items.sku")

	querySql := queryBuilder.String()

	dbORM := orm.NewOrm()
	dbORM.Using("default")
	var err error

	valueItem := []models.ValueItem{}
	totalRows, err := dbORM.Raw(querySql).QueryRows(&valueItem)
	if err != nil {
		return nil, err

	}

	headerValueItem := []models.HeaderValueItem{}
	querySql = "select sum(total) as total_value, sum(stock) as total_stock from (" + querySql + ")"
	_, err = dbORM.Raw(querySql).QueryRows(&headerValueItem)
	if err != nil {
		return nil, err

	}
	headerValueItem[0].TotalSku = strconv.FormatInt(totalRows, 10)

	reportValueItems := models.ReportValueItem{
		Header: headerValueItem[0],
		Items:  valueItem,
	}
	return &reportValueItems, nil

}
