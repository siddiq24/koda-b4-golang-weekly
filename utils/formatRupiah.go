package utils

import "fmt"

func ToRP(nomi uint) string {
	nomiStr := fmt.Sprintf("%d", nomi)
	leng := len(nomiStr)
	result := ""
	for i := 1; i <= leng; i++ {
		if (leng-i+1)%3 == 0 && i != 1 {
			result += "."
		}
		result += string(nomiStr[i-1])
	}
	return result
}
