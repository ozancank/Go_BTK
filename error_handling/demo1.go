package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("error_handling/demo1.txt")
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamad─▒. ", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamad─▒.")
		}
	} else {
		fmt.Println(f.Name())
	}
}
