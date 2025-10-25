package models

import (
	"fmt"
)

type Product struct {
	Id int
	Name  string
	Price uint
}

var Products []Product = []Product{
	{Id: 0, Name: "Americano", Price: 20000},
	{Id: 1, Name: "Kopi Susu", Price: 24000},
	{Id: 2, Name: "Latte", Price: 23000},
	{Id: 3, Name: "Caramel Latte", Price: 28000},
	{Id: 4, Name: "Caramel Machiato", Price: 30000},
	{Id: 5, Name: "Mocha", Price: 29000},
	{Id: 6, Name: "Earl Gray Tea", Price: 20000},
	{Id: 7, Name: "Lemon Tea", Price: 26000},
	{Id: 8, Name: "Lychee Tea", Price: 26000},
	{Id: 9, Name: "Earl Gray Milk Tea", Price: 28000},
	{Id: 10, Name: "Hazelnut Choco", Price: 28000},
	{Id: 11, Name: "Royale Milo", Price: 30000},
	{Id: 12, Name: "Milo Macchiato", Price: 30000},
	{Id: 13, Name: "Americano 1L", Price: 75000},
	{Id: 14, Name: "Latte 1L", Price: 90000},
	{Id: 15, Name: "Kopi Susu 1L", Price: 100000},
	{Id: 16, Name: "Classic Chocolate 1L", Price: 100000},
	{Id: 17, Name: "Caramel Latte 1L", Price: 100000},
	{Id: 18, Name: "Vanilla Latte 1L", Price: 100000},
	{Id: 19, Name: "Hazelnut Latte 1L", Price: 90000},
	{Id: 20, Name: "Hazelnut Choco 1L", Price: 100000},
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
	fmt.Printf("%d%s| %s%s|  Rp.%s\n", i, spaceI, p.Name, spaceN, ToRP(p.Price))
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

func ToRP(nomi uint) string{
	nomiStr := fmt.Sprintf("%d", nomi)
	leng := len(nomiStr)
	result := ""
	for i := 1; i <= leng; i++ {
		if i%3 == 0 {
			result += "."
		}
		result += string(nomiStr[i-1])
	}
	return result
}