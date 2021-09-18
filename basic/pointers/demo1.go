package pointers

import "fmt"

func Demo1(sayi *int) { //yıldız ram adresi anlamına gelir
	*sayi = *sayi + 1
	fmt.Println("demo:", sayi)
	fmt.Println("demo:", *sayi)
}
