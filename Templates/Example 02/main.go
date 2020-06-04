package main

import (
	"log"
	"os"
	"text/template"
)

type region struct {
	Region     string
	CaliHotels []caliHotel
}

type caliHotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []region{
		region{
			Region: "Southern",
			CaliHotels: []caliHotel{
				caliHotel{
					Name:    "Ibis",
					Address: `Qualquer endereço funcionã, n222003`,
					City:    `A`,
					Zip:     `210392183912`,
				},
				caliHotel{
					Name:    "Premium",
					Address: `Qualquer endereço funciona certo, n222145003`,
					City:    `B`,
					Zip:     `210392183123124912`,
				},
			},
		},
		region{
			Region: "Central",
			CaliHotels: []caliHotel{
				caliHotel{
					Name:    "IbisCentral",
					Address: `Qualquer endereço funcionã, n22200`,
					City:    `F`,
					Zip:     `2103921312383912`,
				},
				caliHotel{
					Name:    "PremiumCentral",
					Address: `Qualquer endereço funciona certo, n222`,
					City:    `F`,
					Zip:     `21039215161783912`,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
