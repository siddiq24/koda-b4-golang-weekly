package utils

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/siddiq24/golang-weekly/models"
)

func ShowStruk(bills []models.Bill, id int) {
	var wg sync.WaitGroup
	c := ""
	d := ""
	s := "     "
	w := 40
	space := ""
	dash := ""
	totI := 0
	totP := uint(0)
	stot := ""
	spay := ""
	scha := ""
	for i := range w {
		if i < 27-len(strconv.Itoa(int(bills[id].Total))) {
			stot += " "
		}
		if i < 22-len(strconv.Itoa(int(bills[id].Bayar))) {
			spay += " "
		}
		if i < 26-len(strconv.Itoa(int(bills[id].Change))) {
			scha += " "
		}
		if i < w/2-2 {
			space += " "
		}
		if i < w {
			dash += "-"
		}
		c += string(rune(30))
		d += string(rune(31))
	}
	Print(fmt.Sprintf("\n\n%s %s", s, c))
	// ============================================================
	Print(fmt.Sprintf("%s|%sBILL%s|", s, space, space))
	Print(fmt.Sprintf("%s|%s|", s, dash))
	Print(fmt.Sprintf("%s| Order Number: %s           |", s, bills[id].Invoice))
	Print(fmt.Sprintf("%s| Date        :%s           |", s, bills[id].Date))
	Print(fmt.Sprintf("%s| Time        :%s              |", s, bills[id].Time))
	Print(fmt.Sprintf("%s|%s|", s, dash))

	iniChan := make(chan string, len(bills[id].Carts)*3)

	for _, cart := range bills[id].Carts {
		wg.Add(1)
		go withGor(cart, &totI, &totP, w, s, space, &wg, iniChan)
	}

	go func() {
		wg.Wait()
		close(iniChan)
	}()

	for ch := range iniChan {
		Print(ch)
	}

	Print(fmt.Sprintf("%s|%s%s    |", s, space, space))
	Print(fmt.Sprintf("%s|%s|", s, dash))
	Print(fmt.Sprintf("%s|  Total Item %d                          |", s, totI))
	Print(fmt.Sprintf("%s|%s|", s, dash))
	Print(fmt.Sprintf("%s|  Total%sRp %s  |", s, stot, ToRP(totP)))
	Print(fmt.Sprintf("%s|  Payment:                              |", s))
	Print(fmt.Sprintf("%s|     %s%sRp.%s  |", s, bills[id].Payment, spay, ToRP(bills[id].Bayar)))
	Print(fmt.Sprintf("%s|  Change%sRp.%s  |", s, scha, ToRP(bills[id].Change)))
	Print(fmt.Sprintf("%s|%s%s    |", s, space, space))
	Print(fmt.Sprintf("%s|%s%s    |", s, space, space))
	// ============================================================
	Print(fmt.Sprintf("%s %s\n", s, d))
}

func Print(s string) {
	fmt.Println(s)
	time.Sleep(100 * time.Millisecond)
}

func withGor(cart models.Cart, totI *int, totP *uint, w int, s, space string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	total := uint(cart.Qty) * cart.Price
	*totI += cart.Qty
	*totP += total
	sn := w - 2 - len(cart.Name)
	spaceName := ""
	for range sn {
		spaceName += " "
	}

	ch <- fmt.Sprintf("%s|%s%s    |", s, space, space)
	ch <- fmt.Sprintf("%s|  %s%s|", s, cart.Name, spaceName)
	ch <- fmt.Sprintf("%s|     %d X %s              Rp %s  |", s, cart.Qty, ToRP(cart.Price), ToRP(cart.Price))
}
