package defer_statement

import "fmt"

func A() {
	fmt.Println("a fonksiyonu çalıştı")
}

func C() {
	fmt.Println("c fonksiyonu çalıştı")
}

func B() {
	defer A() // b bittikten sonra çalışır
	defer C()
	fmt.Println("b fonksiyonu çalıştı")
}
