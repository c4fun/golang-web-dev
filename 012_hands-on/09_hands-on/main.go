package main

import (
	"log"
	"os"
	"text/template"
	"encoding/csv"
	"time"
	"strconv"
)

type Record struct {
	Date time.Time
	Open float64
}

var tpl *template.Template

func ParseRecords(filePath string) []Record {
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()
	
	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	// https://tour.golang.org/moretypes/13
	records := make([]Record, 0, len(rows))

	for i, row := range rows {
		if i == 0 {
			continue
		}
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}
	return records
}

func main() {
	// 1. parse the csv file
	var filePath = "table.csv"
	records := ParseRecords(filePath)

	// 2. parse template
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// 3. 把从csv读取到的数据表达在template里面
	err = tpl.Execute(os.Stdout, records)
	if err != nil {
		log.Fatalln(err)
	}
}
