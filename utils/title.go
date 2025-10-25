package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func Title(title string) {
	lgt := len(title)
	space := "   "
	spc := ""
	dashN := lgt + 36
	dash := ""
	for i := range dashN{
		dash += "="
		if i < (dashN-27)/2{
			spc += "="
		}

	}

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println(dash)
	fmt.Printf("===============%s%s%s===============\n", space, title, space)
	fmt.Printf("%s    %s    %s\n", spc, "J A N J I    J I W A", spc)
	fmt.Printf("%s\n\n",dash)
}