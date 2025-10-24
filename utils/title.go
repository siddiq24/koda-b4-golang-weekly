package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func Title(title string) {
	lgt := len(title)
	space := "   "
	dashN := lgt + 6 + 30
	dash := ""
	for range dashN{
		dash += "="
	}

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println(dash)
	fmt.Printf("===============%s%s%s===============\n", space, title, space)
	fmt.Printf("%s\n\n",dash)
}