package main

import (
	"fmt"
	"os"

	"github.com/siddiq24/golang-weekly/pages"
	"github.com/siddiq24/golang-weekly/utils"
)

var data pages.Db

func main() {
	defer func() {
		if r := recover(); r != nil {
			msg := "\n" + fmt.Sprint(r) + "\n>> Back to HOME...."
			utils.Alert(msg)
			main()
		}
	}()

	for {
		utils.Title("Welcome to My System")
		fmt.Println("   1. Menu\n   2. Cart\n   3. History")
		fmt.Println("\n   x. Exit")
		fmt.Printf("\nChoose a menu: ")
		var input string
		fmt.Scanln(&input)

		switch input {
		case "":
			main()
		case "x":
			utils.Alert("Exitting...")
			os.Exit(0)
		case "1":
			data.MenuPage()
		case "2":
			data.Cart()
		case "3":
			data.History()
		default:
			continue
		}
	}
}
