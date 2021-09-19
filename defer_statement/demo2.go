package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	defer Log()
	fmt.Println("Ürün kaydedildi: ", p.productName)
}

func Log() {
	fmt.Println("Loglandı")
}

func Demo2() {
	p := product{productName: "Laptop", unitPrice: 5000}
	p.save()
	p = product{productName: "Mouse", unitPrice: 45}
	fmt.Println("İşlem başarılı")
}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 5000}
	defer p.save()
	p = product{productName: "Mouse", unitPrice: 45}
	fmt.Println("İşlem başarılı")
}
