package pages

import (
	"fmt"
	"github.com/siddiq24/golang-weekly/utils"
)

func (db *Db) Option() {
	for {
		utils.Title("C A C H E   M A N A G E M E N T")

		fmt.Println("     1. Cache Location")
		fmt.Println("     2. Clear Cache")
		fmt.Println("\n     x. Exit")

		var n string

		fmt.Print("\nChoice one:")
		fmt.Scan(&n)

		if n == "1" {
			utils.Title("CACHE MANAGEMENT")
			utils.Alert(fmt.Sprintf("\n>>> Cache tersimpan di:  %s", CacheFile))
			fmt.Scanln()
		}
		if n == "2" {
			db.ClearCache()
		}
		if n == "x" {
			panic("")
		}
	}
}
