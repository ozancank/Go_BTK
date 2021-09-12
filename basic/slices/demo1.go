package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)

	fmt.Println(isimler)

	isimler[0] = "a"
	isimler[1] = "b"
	isimler[2] = "c"

	isimler = append(isimler, "d")
	fmt.Println(isimler)
}
