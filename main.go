package main

import (
	"fmt"
	"os"

	"github.com/siddiq24/golang-weekly/pages"
	"github.com/siddiq24/golang-weekly/utils"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			msg := "\n" + fmt.Sprint(r) + "\nClose Program. . ."
			utils.Alert(msg)
		}
	}()

	for {
		utils.Title("Welcome to My System")
		fmt.Println("   1. Menu\n   2. Cart\n   3. History")
		fmt.Println("\n   0. Exit")
		fmt.Printf("\nChoose a menu: ")
		var input string
		fmt.Scanln(&input)

		switch input {
		case "0":
			utils.Alert("Exitting...")
			os.Exit(0)
		case "1":
			pages.MenuPage()
		case "2":
		case "3":
		default:
			utils.Alert("\n>>> 404 Not Found. . . ")
		}
	}
}
