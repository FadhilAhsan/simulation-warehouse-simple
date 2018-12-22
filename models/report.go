package models

import ()

type ReportValueItem struct {
	Header HeaderValueItem `json:"header"`
	Items  []ValueItem     `json:"items"`
}

type HeaderValueItem struct {
	TotalValue string `json:"total_value"`
	TotalStock string `json:"total_stock"`
	TotalSku   string `json:"total_sku"`
}

type ValueItem struct {
	Sku       string `json:"sku"`
	Name      string `json:"name"`
	Size      string `json:"size"`
	Color     string `json:"color"`
	Stock     string `json:"stock"`
	PriceItem string `json:"price_item"`
	Total     string `json:"total"`
}

type ReportOrder struct {
	Header HeaderReportOrderItem `json:"header"`
	Items  []ReportOrderItem     `json:"items"`
}

type HeaderReportOrderItem struct {
	TotalProfit string `json:"total_profit"`
	TotalIncome string `json:"total_income"`
	TotalSeles  string `json:"total_seles"`
	TotalItem   string `json:"total_item"`
}

type ReportOrderItem struct {
	IdOrder       string `json:"id_order"`
	CreatedAt     string `json:"created_at"`
	Sku           string `json:"sku"`
	Name          string `json:"name"`
	Size          string `json:"size"`
	Color         string `json:"color"`
	Qty           string `json:"qty"`
	Price         string `json:"price"`
	Total         string `json:"total"`
	PurchasePrice string `json:"purchase_price"`
	Profit        string `jsong:"profit"`
}
