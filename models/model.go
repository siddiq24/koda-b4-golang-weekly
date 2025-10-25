package models

import (
	"fmt"
)

type Product struct {
	Name  string
	Price uint
}

var Products []Product = []Product{
	{Name: "Americano", Price: 20000},
	{Name: "Kopi Susu", Price: 24000},
	{Name: "Latte", Price: 23000},
	{Name: "Caramel Latte", Price: 28000},
	{Name: "Caramel Machiato", Price: 30000},
	{Name: "Mocha", Price: 29000},
	{Name: "Earl Gray Tea", Price: 20000},
	{Name: "Lemon Tea", Price: 26000},
	{Name: "Lychee Tea", Price: 26000},
	{Name: "Earl Gray Milk Tea", Price: 28000},
	{Name: "Hazelnut Choco", Price: 28000},
	{Name: "Royale Milo", Price: 30000},
	{Name: "Milo Macchiato", Price: 30000},
	{Name: "Americano 1L", Price: 75000},
	{Name: "Latte 1L", Price: 90000},
	{Name: "Kopi Susu 1L", Price: 100000},
	{Name: "Classic Chocolate 1L", Price: 100000},
	{Name: "Caramel Latte 1L", Price: 100000},
	{Name: "Vanilla Latte 1L", Price: 100000},
	{Name: "Hazelnut Latte 1L", Price: 90000},
	{Name: "Hazelnut Choco 1L", Price: 100000},
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
