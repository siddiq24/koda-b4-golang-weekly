package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/siddiq24/golang-weekly/pages"
	"github.com/siddiq24/golang-weekly/utils"
)

var data pages.Db 

func init(){
	utils.Title("W E L L C O M E    T O")

	err := godotenv.Load()
    if err != nil {
        fmt.Println("Tidak dapat memuat file .env, lanjutkan dengan env sistem.")
    }


	wak, er := strconv.Atoi(os.Getenv("TIME"))
	if er != nil{
		utils.Alert("\nGagal membaca ENV")
		wak = 15
		utils.Alert(fmt.Sprintf("\n\nTemporary file akan diperbaharui setiap %d detik", wak))
		return
	}
	utils.Alert("\nBerhasil membaca env")
	utils.Alert(fmt.Sprintf("\n\nTemporary file akan diperbaharui setiap %d detik", wak))
}


func main() {
	defer func() {
		if r := recover(); r != nil {
			msg := "\n" + fmt.Sprint(r) + "\n>> Back to HOME...."
			utils.Alert(msg)
			main()
		}
	}()

	for {
		utils.Title("W E L L C O M E    T O")
		fmt.Println("   1. Menu\n   2. Cart\n   3. History\n   4. Settings")
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
		case "4":
			data.Option()
		default:
			continue
		}
	}
}
