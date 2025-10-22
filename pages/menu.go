package pages

import (
	"fmt"

	"github.com/siddiq24/golang-weekly/models"
	"github.com/siddiq24/golang-weekly/utils"
)

func MenuPage() {
	listMenu()
	utils.Alert("\n>> You wanna buy? (y/n): ")

	var choice string
	fmt.Scanln(&choice)

	if choice == "y" {
		addToCart([]int{}, 0)
	} else {
		utils.Alert("Exitting...")
	}
}

func listMenu() {
	utils.Title("    M E N U    P R O D U C T    ")

	fmt.Println("__________________________________________________")
	fmt.Println("NO  |  NAME                            |  PRICE")
	fmt.Println("--------------------------------------------------")

	for i, product := range models.Products {
		product.PrintProduct(i + 1)
	}
}

func addToCart(ids []int, total int) {
	listMenu()

	fmt.Println("\n\nMy Cart")
	fmt.Println("__________________________________________________")
	fmt.Println("NO  |  NAME                            |  PRICE")
	fmt.Println("--------------------------------------------------")

	if len(ids) > 0 {
		for i, id := range ids {
			if id < 0 || id >= len(models.Products) {
				continue // hindari panic jika index invalid
			}
			product := models.Products[id]
			product.PrintProduct(i + 1)
			total += product.Price
		}
		fmt.Println("__________________________________________________")
		fmt.Printf("  TOTAL                                | Rp.%d\n", total)
		fmt.Println("--------------------------------------------------")
	}

	fmt.Println("\npress 0 to exit")
	utils.Alert("Add To Cart (product number): ")

	var n int
	if _, err := fmt.Scan(&n); err != nil {
		utils.Alert("Invalid input. Please enter a number.")
		addToCart(ids, total)
		return
	}

	if n == 0 {
		utils.Alert("Exitting...")
		return
	}

	if n < 1 || n > len(models.Products) {
		utils.Alert("Product number not found!")
		addToCart(ids, total)
		return
	}

	ids = append(ids, n-1)
	addToCart(ids, total)
}
