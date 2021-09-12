package variables

import "fmt"

func Demo1() {
	var text string = "Hello World"
	fmt.Println(text)
	fmt.Println(text)
	fmt.Println(text)
	fmt.Println(text)

	var sayi int = 100
	var kdv float64 = 0.18
	fmt.Println(kdv)
	fmt.Println(float64(sayi) + float64(sayi)*kdv)

	kdv2 := 25
	fmt.Println(kdv2)
	fmt.Printf("Veri Tipleri : %T - %T\n", kdv, kdv2)

	durum := false
	fmt.Println(durum)
	fmt.Println(!durum)
}
