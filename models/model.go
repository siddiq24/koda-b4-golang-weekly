package models

import (
	"fmt"
)

type Product struct {
	Name  string
	Price uint
}

var Products []Product = []Product{
	{Name: "Big Mac", Price: 47273},
	{Name: "Cheeseburger", Price: 25000},
	{Name: "McChicken", Price: 36364},
	{Name: "McSpicy", Price: 45455},
	{Name: "Fish Fillet Burger", Price: 36364},
	{Name: "Ayam Krispy (1 pcs)", Price: 20000},
	{Name: "Paket PaNas 1", Price: 38182},
	{Name: "Paket PaNas 2", Price: 54545},
	{Name: "PaMer 5", Price: 127273},
	{Name: "McSpaghetti", Price: 27273},
	{Name: "Korean Soy Garlic Wings (6 pcs)", Price: 54545},
	{Name: "Chicken Muffin", Price: 20000},
	{Name: "Hotcakes 3 pcs", Price: 27273},
	{Name: "French Fries (Medium)", Price: 20000},
	{Name: "McNuggets 6 pcs", Price: 27273},
	{Name: "Spicy Chicken Bites", Price: 18182},
	{Name: "Hashbrown", Price: 13636},
	{Name: "Coca-Cola (Medium)", Price: 13636},
	{Name: "Iced Coffee Float", Price: 20000},
	{Name: "Sundae Coklat", Price: 13636},
	{Name: "McFlurry Oreo", Price: 18182},
}

func (p Product) PrintProduct(i int) {
	spaceI := "   "
	if i > 9 {
		spaceI = "  "
	}
	spaceN := ""
	for range 33 - len(p.Name) {
		spaceN += " "
	}
	fmt.Printf("%d%s| %s%s|  Rp.%d\n", i, spaceI, p.Name, spaceN, p.Price)
}

func (c Cart) PrintProduct(i int) {
	spaceI := "   "
	if c.Id > 9 {
		spaceI = "  "
	}
	spaceN := ""
	for range 33 - len(c.Name) {
		spaceN += " "
	}
	spaceQ := " "
	if c.Qty > 9 {
		spaceQ = ""
	}
	if i < 0 {
		fmt.Printf("    |                                  |       |\n")
	} else {
		fmt.Printf("%d%s| %s%s|   %d%s  |  Rp.%d\n", i, spaceI, c.Name, spaceN, c.Qty, spaceQ, c.Total)
	}
}

type Cart struct {
	Id    int
	Name  string
	Price uint
	Qty   int
	Total int
}

type Bill struct {
	Carts   []Cart
	Invoice string
	Date    string
	Time    string
	Payment string
	Total   uint
	Bayar   uint
	Change  uint
}
