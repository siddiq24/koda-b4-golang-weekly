package pages

import (
	"fmt"
	"strconv"

	"github.com/siddiq24/golang-weekly/models"
	"github.com/siddiq24/golang-weekly/utils"
)

type Db struct {
	Products []models.Product
	Carts    []models.Cart
	Bills    []models.Bill
	Total    uint
}

type InterfaceProduct interface {
	PrintProduct(i int)
}

func ShowProduct(ip InterfaceProduct, i int) {
	ip.PrintProduct(i)
}

func (db *Db) MenuPage() {
	db.Products = models.Products
	db.listMenu()
	fmt.Println("Total :", db.Total)
	fmt.Print("\n>> You wanna buy? ")
	utils.Alert("(y/n): ")
	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "":
		db.addToCart(0)
	case "y":
		db.addToCart(0)
		return
	case "n":
		panic("Its OK")
	default:
		utils.Alert("\nEnter the command correctly")
		db.MenuPage()
	}
}

func (db *Db) listMenu() {
	utils.Title("    M E N U    P R O D U C T    ")

	fmt.Println("┌─────┬──────────────────────────────────┬─────────")
	fmt.Println("  NO  |  NAME                            |  PRICE")
	fmt.Println("---------------------------------------------------")

	for i, product := range db.Products {
		fmt.Print("  ")
		ShowProduct(product, i+1)
	}
}

func (db *Db) addToCart(errr int) {
	if errr == 3 {
		panic("")
	}
	delete := false
	for {
		q := 0

		for _, cart := range db.Carts {
			q += cart.Qty
		}

		db.showChart(q)

		fmt.Println("\n\n==========================================================")
		if delete {
			fmt.Println("   add  : add item to cart")
		} else {
			fmt.Println("   del  : delete item from cart")
		}
		fmt.Println("   x    : back home")

		var n string

		if delete {
			utils.Alert("\nDelete from Cart: ")
			if _, err := fmt.Scanln(&n); err != nil {
				continue
			}

			if n == "add" {
				delete = false
				continue
			}

			if n == "x" {
				db.Exit()
				continue
			}

			in, _ := strconv.Atoi(n)

			if in < 0 || in > len(db.Carts) {
				continue
			}

			db.Carts = append(db.Carts[:in-1], db.Carts[in:]...)
			continue
		}

		fmt.Print("\nAdd To Cart (product ")
		utils.Alert("number): ")
		if _, err := fmt.Scanln(&n); err != nil {
			continue
		}
		if n == "del" {
			if len(db.Carts) > 0 {
				delete = true
				continue
			}
			fmt.Println("Cannot Delete Cart")
			utils.Alert("Your Chart is empty...")
			continue
		}

		if n == "x" {
			db.Exit()
			continue
		}

		num := db.StoI(n, errr)

		if num < 1 || num > len(db.Products) {
			utils.Alert("Product number not found!")
			continue
		}

		product := db.Products[num-1]
		found := false
		for i := range db.Carts {
			if db.Carts[i].Name == product.Name {
				db.Carts[i].Qty++
				db.Carts[i].Total+=int(product.Price)
				found = true
				break
			}
		}
		if !found {
			db.Carts = append(db.Carts, models.Cart{
				Id:    len(db.Carts),
				Name:  product.Name,
				Price: product.Price,
				Qty:   1,
				Total: int(product.Price),
			})
		}
	}
}

func (db *Db) StoI(str string, r int) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		utils.Alert("\nInvalid input")
		db.addToCart(r + 1)
	}
	return num
}

func (db *Db) showChart(q int) {
	db.listMenu()
	fmt.Println("\n\nMy Cart: ", q, "Items")
	fmt.Println("__________________________________________________________")
	fmt.Println("NO  |  NAME                            |  Qty  |  PRICE")
	fmt.Println("----------------------------------------------------------")

	var total uint
	for i, cart := range db.Carts {
		ShowProduct(cart, i+1)
		total += uint(cart.Total)
		q += cart.Qty
	}
	db.Total =total

	if len(db.Carts) > 0 {
		fmt.Println("----------------------------------------------------------")
		fmt.Printf("     TOTAL                                     | Rp.%d\n", db.Total)
		fmt.Println("----------------------------------------------------------")
	}
}

func (db Db) Exit() {
	var nn string
	if len(db.Carts) > 0 {
		panic("\nProduct has been saved...")
	}
	fmt.Println("Anda Belum memesan apapun.")
	fmt.Print("\nAre you sure you want to leave...(y/n)? ")
	fmt.Scanln(&nn)
	if nn == "n" {
		utils.Alert("\nContinue to buying....")
	} else {
		panic("Its OK..")
	}
}
