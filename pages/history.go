package pages

import (
	"fmt"

	"github.com/siddiq24/golang-weekly/utils"
)

func (db *Db) History() {
	utils.Title(" H I S T O R Y ")

	for i, bill := range db.Bills {
		fmt.Println("==================================================")
		fmt.Println("#", i+1)
		fmt.Printf("INVOICE     :  %s\n", bill.Invoice)
		fmt.Println("DATE        : ", bill.Date)
		fmt.Println("PAYMENT     : ", bill.Payment)
		fmt.Println("")
	}

	fmt.Scanln()
}
