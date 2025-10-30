package models

import (
	"fmt"
)

type Product struct {
	Id    int
	Name  string
	Price uint
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
	fmt.Printf("%d%s| %s%s|  Rp %s\n", i, spaceI, p.Name, spaceN, ToRP(p.Price))
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
		fmt.Printf("%d%s| %s%s|   %d%s  |  Rp %s\n", i, spaceI, c.Name, spaceN, c.Qty, spaceQ, ToRP(uint(c.Total)))
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

func ToRP(nomi uint) string {
	nomiStr := fmt.Sprintf("%d", nomi)
	leng := len(nomiStr)
	result := ""
	for i := 1; i <= leng; i++ {
		if (leng-i+1)%3 == 0 && i != 1 {
			result += "."
		}
		result += string(nomiStr[i-1])
	}
	return result
}
