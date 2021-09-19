package functions

import "fmt"

func Topla(sayi1 int, sayi2 int) int {
	var toplam int = sayi1 + sayi2
	return toplam
}

func DortIslem(sayi1 int, sayi2 int) (int, int, int, float32) {
	toplam := sayi1 + sayi2
	cikarim := sayi1 - sayi2
	carpim := sayi1 * sayi2
	bolum := float32(sayi1 / sayi2)
	return toplam, cikarim, carpim, bolum
}

func SelamVer() {
	fmt.Println("Merhaba")
}
