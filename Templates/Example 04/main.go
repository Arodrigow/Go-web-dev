package main

import (
	"log"
	"os"
	"text/template"
)

type restaurants struct {
	Name string
	Menu menu
}

type menu struct {
	Breakfast []breakfast
	Lunch     []lunch
	Dinner    []dinner
}
type breakfast struct {
	Item  string
	Price float64
}

type lunch struct {
	Item  string
	Price float64
}

type dinner struct {
	Item  string
	Price float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	restaurants := []restaurants{
		restaurants{
			Name: "Rest01",
			Menu: menu{
				Breakfast: []breakfast{
					breakfast{
						Item:  `Bread`,
						Price: 0.30,
					},
					breakfast{
						Item:  `Juice`,
						Price: 10.00,
					},
				},
				Lunch: []lunch{
					lunch{
						Item:  `Rice`,
						Price: 20.00,
					},
					lunch{
						Item:  `Beans`,
						Price: 400.00,
					},
				},
				Dinner: []dinner{
					dinner{
						Item:  `Soup`,
						Price: 3.00,
					},
					dinner{
						Item:  `shake`,
						Price: 23.00,
					},
				},
			},
		},
		restaurants{
			Name: "Rest02",
			Menu: menu{
				Breakfast: []breakfast{
					breakfast{
						Item:  `Bread`,
						Price: 0.30,
					},
					breakfast{
						Item:  `Juice`,
						Price: 10.00,
					},
				},
				Lunch: []lunch{
					lunch{
						Item:  `Rice`,
						Price: 20.00,
					},
					lunch{
						Item:  `Beans`,
						Price: 400.00,
					},
				},
				Dinner: []dinner{
					dinner{
						Item:  `Soup`,
						Price: 3.00,
					},
					dinner{
						Item:  `shake`,
						Price: 23.00,
					},
				},
			},
		},
		restaurants{
			Name: "Rest03",
			Menu: menu{
				Breakfast: []breakfast{
					breakfast{
						Item:  `Bread`,
						Price: 0.30,
					},
					breakfast{
						Item:  `Juice`,
						Price: 10.00,
					},
				},
				Lunch: []lunch{
					lunch{
						Item:  `Rice`,
						Price: 20.00,
					},
					lunch{
						Item:  `Beans`,
						Price: 400.00,
					},
				},
				Dinner: []dinner{
					dinner{
						Item:  `Soup`,
						Price: 3.00,
					},
					dinner{
						Item:  `shake`,
						Price: 23.00,
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
