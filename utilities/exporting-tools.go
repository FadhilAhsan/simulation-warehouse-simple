package utilities

import (
	"encoding/csv"
	"github.com/FadhilAhsan/simulation-warehouse-simple/models"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

func ExportReportValueItemToCsv(data models.ReportValueItem) string {
	filename := "report-value-item.csv"

	if stat, err := os.Stat("/csv-data"); err == nil && stat.IsDir() {
		os.Mkdir("csv-data", os.ModePerm)
	}

	basepath := path.Base("/csv-data")
	filepath := path.Base(filename)

	file, err := os.Create(strings.Join([]string{basepath, filepath}, "/"))
	checkError("Cannot create file", err)
	defer file.Close()

	// Initialize csv writer
	writer := csv.NewWriter(file)

	// Write header into file
	writer.Write([]string{"LAPORAN NILAI BARANG"})
	writer.Write([]string{""})

	layoutTime := "02 Jan 06"
	record := []string{
		"Tanggal Cetak", time.Now().Format(layoutTime),
	}
	writer.Write(record)

	record = []string{
		"Jumlah SKU ", data.Header.TotalSku,
	}
	writer.Write(record)

	record = []string{
		"Jumlah Total Barang ", data.Header.TotalStock,
	}
	writer.Write(record)

	record = []string{
		"Total Nilai", "Rp " + data.Header.TotalValue,
	}
	writer.Write(record)
	writer.Write([]string{""})

	record = []string{
		"SKU", "Nama Item", "Jumlah", "Harga Beli", "Total",
	}
	writer.Write(record)

	for _, item := range data.Items {
		record = []string{
			item.Sku,
			item.Name + " (" + item.Size + "," + item.Color + ")",
			item.Stock,
			"Rp " + item.PriceItem,
			"Rp " + item.Total,
		}
		err = writer.Write(record)
		checkError("Cannot write to file", err)
	}
	defer writer.Flush()
	return filename
}

func ExportReportOrderItemToCsv(data models.ReportOrder, startDate string, endDate string) string {
	filename := "report-sales.csv"

	if stat, err := os.Stat("/csv-data"); err == nil && stat.IsDir() {
		os.Mkdir("csv-data", os.ModePerm)
	}

	basepath := path.Base("/csv-data")
	filepath := path.Base(filename)

	file, err := os.Create(strings.Join([]string{basepath, filepath}, "/"))
	checkError("Cannot create file", err)
	defer file.Close()

	// Initialize csv writer
	writer := csv.NewWriter(file)

	// Write header into file
	writer.Write([]string{"LAPORAN Penjualan"})
	writer.Write([]string{""})

	layoutTime := "02 Jan 06"
	record := []string{
		"Tanggal Cetak", time.Now().Format(layoutTime),
	}
	writer.Write(record)

	record = []string{
		"Tanggal", startDate + " - " + endDate,
	}
	writer.Write(record)

	record = []string{
		"Total Omzet ", "Rp " + data.Header.TotalIncome,
	}
	writer.Write(record)

	record = []string{
		"Total Laba Kotor ", "Rp " + data.Header.TotalProfit,
	}
	writer.Write(record)

	record = []string{
		"Total Penjualan ", data.Header.TotalSeles,
	}
	writer.Write(record)

	record = []string{
		"Total Barang ", data.Header.TotalItem,
	}
	writer.Write(record)
	writer.Write([]string{""})

	record = []string{
		"ID Pesanan", "Waktu", "SKU", "Nama Barang", "Jumlah", "Harga Jual", "Total", "Harga Beli", "Laba",
	}
	writer.Write(record)

	for _, item := range data.Items {
		record = []string{
			item.IdOrder,
			item.CreatedAt,
			item.Sku,
			item.Name + " (" + item.Size + "," + item.Color + ")",
			item.Qty,
			"Rp " + item.Price,
			"Rp " + item.Total,
			"Rp " + item.PurchasePrice,
			"Rp " + item.Profit,
		}
		err = writer.Write(record)
		checkError("Cannot write to file", err)
	}
	defer writer.Flush()
	return filename
}

func checkError(message string, err error) {
	if err != nil {
		log.Println(message, err)
	}
}
