package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

type record struct {
	Date time.Time
	Open float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	src, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	r := csv.NewReader(src)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	records := make([]record, 0, len(rows))

	for i, row := range rows {
		if i == 0 {
			continue
		}
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, record{
			Date: date,
			Open: open,
		})
	}

	err = tpl.Execute(os.Stdout, records)
	if err != nil {
		log.Fatalln(err)
	}

	//To see in a web page, just dump main on a html file
	//go run main.go > index.html
}
