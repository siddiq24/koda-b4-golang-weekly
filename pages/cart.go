package pages

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/TigorLazuardi/tanggal"
	"github.com/siddiq24/golang-weekly/models"
	"github.com/siddiq24/golang-weekly/utils"
)

func (db *Db) Cart() {
	utils.Title(" C A R T S ")

	if len(db.Carts) == 0 {
		utils.Alert("Your Carts as empty....")
		return
	} else {
		fmt.Println("__________________________________________________________")
		fmt.Println("NO  |  NAME                            |  Qty  |  PRICE")
		fmt.Println("----------------------------------------------------------")

		for i := range len(db.Carts) + 5 {
			if i < len(db.Carts) {
				db.Carts[i].PrintProduct(i + 1)
				continue
			}
			fmt.Println("    |                                  |       |")
		}
		fmt.Println("----------------------------------------------------------")
		fmt.Printf("       TOTAL                                   | Rp.%d\n", db.Total)
		fmt.Println("----------------------------------------------------------")
	}

	var n string
	utils.Alert("\n>> Checkout right now? (y/n)")
	fmt.Scanln(&n)
	if n == "y" || n == "" {
		db.Payment()
		return
	}
	if n == "n" {
		panic("")
	}
	db.Cart()
}

func (db *Db) Payment() {
	utils.Title("P A Y M E N T")

	paymed := []string{"MANDIRI", "BRI    ", "BCA    ", "BSI    ", "DANA   ", "OVO    ", "GOPAY  ", "CASH   "}
	fmt.Println("Choose Payment Method")
	fmt.Println("   1. MANDIRI     5. DANA")
	fmt.Println("   2. BRI         6. OVO")
	fmt.Println("   3. BCA         7. GOPAY")
	fmt.Println("   4. BSI         8. CASH")

	utils.Alert("\n>> Insert Payment Method: ")

	var n string
	fmt.Scanln(&n)
	in, r := strconv.Atoi(n)

	if in <= 0 || in > 8 || r != nil {
		db.Payment()
	}

	format := []tanggal.Format{
		tanggal.Hari,
		tanggal.NamaBulan,
		tanggal.Tahun,
	}
	formatx := []tanggal.Format{
		tanggal.PukulDenganDetik,
		tanggal.ZonaWaktu,
	}
	tgl, err := tanggal.Papar(time.Now(), "", tanggal.WIB)
	if err != nil {
		log.Fatal(err)
	}

	for {
		utils.Title("P A Y M E N T")

		var p string
		fmt.Printf("\nTotal transaction : Rp.%d\n", db.Total)
		fmt.Print("\nBayar : Rp.")
		fmt.Scanln(&p)
		pay, er := strconv.Atoi(p)
		if er != nil {
			continue
		}

		if db.Total > uint(pay) {
			utils.Alert("Your Money tidak cukup..")
			continue
		}
		db.Bills = append(db.Bills, models.Bill{
			Carts:   db.Carts,
			Invoice: utils.NewOrderCode(),
			Date:    tgl.Format(" ", format),
			Time:    tgl.Format(" ", formatx),
			Payment: paymed[in-1],
			Total:   db.Total,
			Bayar:   uint(pay),
			Change:  uint(pay) - db.Total,
		})

		utils.Title("P A Y M E N T")

		if len(db.Bills) > 0 {
			utils.ShowStruk(db.Bills, len(db.Bills)-1)
		}
		time.Sleep(time.Second)
		utils.Alert("\n   >> Payment Successfull..")
		fmt.Println("\n Enter to exit")
		fmt.Scanln()
		db.Carts = []models.Cart{}
		db.Total = 0
		panic("Riwayat tersimpan di history....")
	}
}
